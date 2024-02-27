package sharesubscription

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShareSubscriptionSynchronization struct {
	DurationMs          *int64               `json:"durationMs,omitempty"`
	EndTime             *string              `json:"endTime,omitempty"`
	Message             *string              `json:"message,omitempty"`
	StartTime           *string              `json:"startTime,omitempty"`
	Status              *string              `json:"status,omitempty"`
	SynchronizationId   string               `json:"synchronizationId"`
	SynchronizationMode *SynchronizationMode `json:"synchronizationMode,omitempty"`
}

func (o *ShareSubscriptionSynchronization) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ShareSubscriptionSynchronization) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}
