package metrics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricValue struct {
	Average   *float64 `json:"average,omitempty"`
	Count     *float64 `json:"count,omitempty"`
	Maximum   *float64 `json:"maximum,omitempty"`
	Minimum   *float64 `json:"minimum,omitempty"`
	TimeStamp string   `json:"timeStamp"`
	Total     *float64 `json:"total,omitempty"`
}

func (o *MetricValue) GetTimeStampAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.TimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *MetricValue) SetTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeStamp = formatted
}
