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
	"time"

	"github.com/apache/incubator-devlake/core/context"
	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/core/plugin"
)

var _ plugin.MigrationScript = (*addPrMetadataToPredictions)(nil)

type addPrMetadataToPredictions struct{}

// Up adds PR display metadata columns to _tool_aireview_failure_predictions.
//
// These denormalised columns (pr_title, pr_url, pr_author, repo_name, pr_created_at,
// additions, deletions) allow Grafana drill-down table panels to show human-readable
// PR details without joining back to the pull_requests and repos tables at query time.
//
// prediction_outcome now also accepts the value "NO_CI", written for AI-reviewed PRs
// that had no matching CI pipeline data in any enabled source. ci_failure_source for
// those records is set to "none".
func (script *addPrMetadataToPredictions) Up(basicRes context.BasicRes) errors.Error {
	db := basicRes.GetDal()
	if err := db.AutoMigrate(&failurePredictionAddPrMetadata20260402{}); err != nil {
		return errors.Default.Wrap(err, "failed to add PR metadata columns to _tool_aireview_failure_predictions")
	}
	return nil
}

func (script *addPrMetadataToPredictions) Version() uint64 {
	return 20260402000002
}

func (script *addPrMetadataToPredictions) Name() string {
	return "aireview add pr_title/url/author/repo_name/size columns and NO_CI outcome support to failure predictions"
}

type failurePredictionAddPrMetadata20260402 struct {
	RepoName    string    `gorm:"type:varchar(255)"`
	PrTitle     string    `gorm:"type:varchar(500)"`
	PrUrl       string    `gorm:"type:varchar(1024)"`
	PrAuthor    string    `gorm:"type:varchar(255)"`
	PrCreatedAt time.Time `gorm:"column:pr_created_at"`
	Additions   int
	Deletions   int
}

func (failurePredictionAddPrMetadata20260402) TableName() string {
	return "_tool_aireview_failure_predictions"
}
