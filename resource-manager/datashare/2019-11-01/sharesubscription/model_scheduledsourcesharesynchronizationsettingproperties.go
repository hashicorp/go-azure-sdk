package sharesubscription

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledSourceShareSynchronizationSettingProperties struct {
	RecurrenceInterval  *RecurrenceInterval `json:"recurrenceInterval,omitempty"`
	SynchronizationTime *string             `json:"synchronizationTime,omitempty"`
}

func (o *ScheduledSourceShareSynchronizationSettingProperties) GetSynchronizationTimeAsTime() (*time.Time, error) {
	if o.SynchronizationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SynchronizationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ScheduledSourceShareSynchronizationSettingProperties) SetSynchronizationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SynchronizationTime = &formatted
}
