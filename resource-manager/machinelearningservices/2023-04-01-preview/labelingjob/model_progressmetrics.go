package labelingjob

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProgressMetrics struct {
	CompletedDatapointCount            *int64  `json:"completedDatapointCount,omitempty"`
	IncrementalDataLastRefreshDateTime *string `json:"incrementalDataLastRefreshDateTime,omitempty"`
	SkippedDatapointCount              *int64  `json:"skippedDatapointCount,omitempty"`
	TotalDatapointCount                *int64  `json:"totalDatapointCount,omitempty"`
}

func (o *ProgressMetrics) GetIncrementalDataLastRefreshDateTimeAsTime() (*time.Time, error) {
	if o.IncrementalDataLastRefreshDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.IncrementalDataLastRefreshDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ProgressMetrics) SetIncrementalDataLastRefreshDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.IncrementalDataLastRefreshDateTime = &formatted
}
