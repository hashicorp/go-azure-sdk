package metricalertsstatus

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricAlertStatusProperties struct {
	Dimensions *map[string]string `json:"dimensions,omitempty"`
	Status     *string            `json:"status,omitempty"`
	Timestamp  *string            `json:"timestamp,omitempty"`
}

func (o *MetricAlertStatusProperties) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *MetricAlertStatusProperties) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
