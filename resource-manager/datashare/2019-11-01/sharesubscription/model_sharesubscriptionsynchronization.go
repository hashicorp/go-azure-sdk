package sharesubscription

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
