package entities

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityGetInsightsParameters struct {
	AddDefaultExtendedTimeRange *bool     `json:"addDefaultExtendedTimeRange,omitempty"`
	EndTime                     string    `json:"endTime"`
	InsightQueryIds             *[]string `json:"insightQueryIds,omitempty"`
	StartTime                   string    `json:"startTime"`
}

func (o *EntityGetInsightsParameters) GetEndTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *EntityGetInsightsParameters) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = formatted
}

func (o *EntityGetInsightsParameters) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *EntityGetInsightsParameters) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
