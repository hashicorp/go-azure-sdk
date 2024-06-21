package sharesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledSourceShareSynchronizationSettingProperties struct {
	RecurrenceInterval  *RecurrenceInterval `json:"recurrenceInterval,omitempty"`
	SynchronizationTime *string             `json:"synchronizationTime,omitempty"`
}
