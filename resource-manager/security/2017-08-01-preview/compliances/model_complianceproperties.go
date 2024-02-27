package compliances

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceProperties struct {
	AssessmentResult           *[]ComplianceSegment `json:"assessmentResult,omitempty"`
	AssessmentTimestampUtcDate *string              `json:"assessmentTimestampUtcDate,omitempty"`
	ResourceCount              *int64               `json:"resourceCount,omitempty"`
}

func (o *ComplianceProperties) GetAssessmentTimestampUtcDateAsTime() (*time.Time, error) {
	if o.AssessmentTimestampUtcDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AssessmentTimestampUtcDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ComplianceProperties) SetAssessmentTimestampUtcDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AssessmentTimestampUtcDate = &formatted
}
