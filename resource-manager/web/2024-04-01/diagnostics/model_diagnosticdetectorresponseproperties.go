package diagnostics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticDetectorResponseProperties struct {
	AbnormalTimePeriods *[]DetectorAbnormalTimePeriod `json:"abnormalTimePeriods,omitempty"`
	Data                *[][]NameValuePair            `json:"data,omitempty"`
	DetectorDefinition  *DetectorDefinition           `json:"detectorDefinition,omitempty"`
	EndTime             *string                       `json:"endTime,omitempty"`
	IssueDetected       *bool                         `json:"issueDetected,omitempty"`
	Metrics             *[]DiagnosticMetricSet        `json:"metrics,omitempty"`
	ResponseMetaData    *ResponseMetaData             `json:"responseMetaData,omitempty"`
	StartTime           *string                       `json:"startTime,omitempty"`
}

func (o *DiagnosticDetectorResponseProperties) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DiagnosticDetectorResponseProperties) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *DiagnosticDetectorResponseProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DiagnosticDetectorResponseProperties) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
