package pools

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolUsageMetrics struct {
	EndTime        string  `json:"endTime"`
	PoolId         string  `json:"poolId"`
	StartTime      string  `json:"startTime"`
	TotalCoreHours float64 `json:"totalCoreHours"`
	VMSize         string  `json:"vmSize"`
}

func (o *PoolUsageMetrics) GetEndTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PoolUsageMetrics) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = formatted
}

func (o *PoolUsageMetrics) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PoolUsageMetrics) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
