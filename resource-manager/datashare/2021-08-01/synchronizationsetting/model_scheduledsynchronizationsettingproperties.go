package synchronizationsetting

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduledSynchronizationSettingProperties struct {
	CreatedAt           *string            `json:"createdAt,omitempty"`
	ProvisioningState   *ProvisioningState `json:"provisioningState,omitempty"`
	RecurrenceInterval  RecurrenceInterval `json:"recurrenceInterval"`
	SynchronizationTime string             `json:"synchronizationTime"`
	UserName            *string            `json:"userName,omitempty"`
}
