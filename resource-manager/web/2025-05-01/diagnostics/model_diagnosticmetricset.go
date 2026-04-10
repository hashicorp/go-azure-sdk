package diagnostics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticMetricSet struct {
	EndTime   *string                   `json:"endTime,omitempty"`
	Name      *string                   `json:"name,omitempty"`
	StartTime *string                   `json:"startTime,omitempty"`
	TimeGrain *string                   `json:"timeGrain,omitempty"`
	Unit      *string                   `json:"unit,omitempty"`
	Values    *[]DiagnosticMetricSample `json:"values,omitempty"`
}

func (o *DiagnosticMetricSet) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DiagnosticMetricSet) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *DiagnosticMetricSet) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DiagnosticMetricSet) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
