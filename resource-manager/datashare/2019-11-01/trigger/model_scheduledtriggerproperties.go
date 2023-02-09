package trigger

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledTriggerProperties struct {
	CreatedAt           *string              `json:"createdAt,omitempty"`
	ProvisioningState   *ProvisioningState   `json:"provisioningState,omitempty"`
	RecurrenceInterval  RecurrenceInterval   `json:"recurrenceInterval"`
	SynchronizationMode *SynchronizationMode `json:"synchronizationMode,omitempty"`
	SynchronizationTime string               `json:"synchronizationTime"`
	TriggerStatus       *TriggerStatus       `json:"triggerStatus,omitempty"`
	UserName            *string              `json:"userName,omitempty"`
}

func (o *ScheduledTriggerProperties) GetCreatedAtAsTime() (*time.Time, error) {
	if o.CreatedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *ScheduledTriggerProperties) SetCreatedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedAt = &formatted
}

func (o *ScheduledTriggerProperties) GetSynchronizationTimeAsTime() (*time.Time, error) {
	return dates.ParseAsFormat(&o.SynchronizationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ScheduledTriggerProperties) SetSynchronizationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SynchronizationTime = formatted
}
