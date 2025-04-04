package databaserecommendedactions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionMetricInfo struct {
	MetricName *string  `json:"metricName,omitempty"`
	StartTime  *string  `json:"startTime,omitempty"`
	TimeGrain  *string  `json:"timeGrain,omitempty"`
	Unit       *string  `json:"unit,omitempty"`
	Value      *float64 `json:"value,omitempty"`
}

func (o *RecommendedActionMetricInfo) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RecommendedActionMetricInfo) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
