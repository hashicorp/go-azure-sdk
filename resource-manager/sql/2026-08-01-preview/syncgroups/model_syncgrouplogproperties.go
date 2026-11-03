package syncgroups

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupLogProperties struct {
	Details         *string           `json:"details,omitempty"`
	OperationStatus *string           `json:"operationStatus,omitempty"`
	Source          *string           `json:"source,omitempty"`
	Timestamp       *string           `json:"timestamp,omitempty"`
	TracingId       *string           `json:"tracingId,omitempty"`
	Type            *SyncGroupLogType `json:"type,omitempty"`
}

func (o *SyncGroupLogProperties) GetTimestampAsTime() (*time.Time, error) {
	if o.Timestamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Timestamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SyncGroupLogProperties) SetTimestampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Timestamp = &formatted
}
