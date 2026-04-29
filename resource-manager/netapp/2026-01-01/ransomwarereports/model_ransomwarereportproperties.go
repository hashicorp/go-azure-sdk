package ransomwarereports

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RansomwareReportProperties struct {
	ClearedCount      *int64                    `json:"clearedCount,omitempty"`
	EventTime         *string                   `json:"eventTime,omitempty"`
	ProvisioningState *string                   `json:"provisioningState,omitempty"`
	ReportedCount     *int64                    `json:"reportedCount,omitempty"`
	Severity          *RansomwareReportSeverity `json:"severity,omitempty"`
	State             *RansomwareReportState    `json:"state,omitempty"`
	Suspects          *[]RansomwareSuspects     `json:"suspects,omitempty"`
}

func (o *RansomwareReportProperties) GetEventTimeAsTime() (*time.Time, error) {
	if o.EventTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EventTime, "2006-01-02T15:04:05Z07:00")
}

func (o *RansomwareReportProperties) SetEventTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EventTime = &formatted
}
