package pools

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceStatistics struct {
	AvgCPUPercentage float64 `json:"avgCPUPercentage"`
	AvgDiskGiB       float64 `json:"avgDiskGiB"`
	AvgMemoryGiB     float64 `json:"avgMemoryGiB"`
	DiskReadGiB      float64 `json:"diskReadGiB"`
	DiskReadIOps     int64   `json:"diskReadIOps"`
	DiskWriteGiB     float64 `json:"diskWriteGiB"`
	DiskWriteIOps    int64   `json:"diskWriteIOps"`
	LastUpdateTime   string  `json:"lastUpdateTime"`
	NetworkReadGiB   float64 `json:"networkReadGiB"`
	NetworkWriteGiB  float64 `json:"networkWriteGiB"`
	PeakDiskGiB      float64 `json:"peakDiskGiB"`
	PeakMemoryGiB    float64 `json:"peakMemoryGiB"`
	StartTime        string  `json:"startTime"`
}

func (o *ResourceStatistics) GetLastUpdateTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.LastUpdateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ResourceStatistics) SetLastUpdateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdateTime = formatted
}

func (o *ResourceStatistics) GetStartTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ResourceStatistics) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = formatted
}
