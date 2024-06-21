package trigger

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
