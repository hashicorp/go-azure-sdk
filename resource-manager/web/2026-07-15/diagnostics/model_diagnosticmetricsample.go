package diagnostics

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticMetricSample struct {
	IsAggregated *bool    `json:"isAggregated,omitempty"`
	Maximum      *float64 `json:"maximum,omitempty"`
	Minimum      *float64 `json:"minimum,omitempty"`
	RoleInstance *string  `json:"roleInstance,omitempty"`
	Timestamp    *string  `json:"timestamp,omitempty"`
	Total        *float64 `json:"total,omitempty"`
}

func (o *DiagnosticMetricSample) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *DiagnosticMetricSample) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
