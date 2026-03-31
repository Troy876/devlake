/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package migrationscripts

import (
	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

var _ plugin.MigrationScript = (*addCiBackfillConfig)(nil)

type addCiBackfillConfig struct{}

func (script *addCiBackfillConfig) Up(basicRes context.BasicRes) errors.Error {
	db := basicRes.GetDal()
	if err := db.AutoMigrate(&scopeConfigAddCiBackfill20260401{}); err != nil {
		return errors.Default.Wrap(err, "failed to add ci_backfill fields to _tool_aireview_scope_configs")
	}
	return nil
}

func (script *addCiBackfillConfig) Version() uint64 {
	return 20260401000001
}

func (script *addCiBackfillConfig) Name() string {
	return "aireview add ci_backfill_enabled and ci_backfill_days to scope config"
}

type scopeConfigAddCiBackfill20260401 struct {
	CiBackfillEnabled bool `gorm:"type:boolean;default:false"`
	CiBackfillDays    int  `gorm:"default:180"`
}

func (scopeConfigAddCiBackfill20260401) TableName() string {
	return "_tool_aireview_scope_configs"
}
