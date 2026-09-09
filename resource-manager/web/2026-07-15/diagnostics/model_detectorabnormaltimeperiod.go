package diagnostics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectorAbnormalTimePeriod struct {
	EndTime   *string            `json:"endTime,omitempty"`
	Message   *string            `json:"message,omitempty"`
	MetaData  *[][]NameValuePair `json:"metaData,omitempty"`
	Priority  *float64           `json:"priority,omitempty"`
	Solutions *[]Solution        `json:"solutions,omitempty"`
	Source    *string            `json:"source,omitempty"`
	StartTime *string            `json:"startTime,omitempty"`
	Type      *IssueType         `json:"type,omitempty"`
}

func (o *DetectorAbnormalTimePeriod) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DetectorAbnormalTimePeriod) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *DetectorAbnormalTimePeriod) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DetectorAbnormalTimePeriod) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
