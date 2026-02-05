package pools

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolStatistics struct {
	LastUpdateTime string              `json:"lastUpdateTime"`
	ResourceStats  *ResourceStatistics `json:"resourceStats,omitempty"`
	StartTime      string              `json:"startTime"`
	Url            string              `json:"url"`
	UsageStats     *UsageStatistics    `json:"usageStats,omitempty"`
}

func (o *PoolStatistics) GetLastUpdateTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastUpdateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PoolStatistics) SetLastUpdateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdateTime = formatted
}

func (o *PoolStatistics) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PoolStatistics) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
