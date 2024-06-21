package queues

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SBQueueProperties struct {
	AccessedAt                          *string              `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string              `json:"autoDeleteOnIdle,omitempty"`
	CountDetails                        *MessageCountDetails `json:"countDetails,omitempty"`
	CreatedAt                           *string              `json:"createdAt,omitempty"`
	DeadLetteringOnMessageExpiration    *bool                `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive            *string              `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string              `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                `json:"enablePartitioning,omitempty"`
	ForwardDeadLetteredMessagesTo       *string              `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                           *string              `json:"forwardTo,omitempty"`
	LockDuration                        *string              `json:"lockDuration,omitempty"`
	MaxDeliveryCount                    *int64               `json:"maxDeliveryCount,omitempty"`
	MaxMessageSizeInKilobytes           *int64               `json:"maxMessageSizeInKilobytes,omitempty"`
	MaxSizeInMegabytes                  *int64               `json:"maxSizeInMegabytes,omitempty"`
	MessageCount                        *int64               `json:"messageCount,omitempty"`
	RequiresDuplicateDetection          *bool                `json:"requiresDuplicateDetection,omitempty"`
	RequiresSession                     *bool                `json:"requiresSession,omitempty"`
	SizeInBytes                         *int64               `json:"sizeInBytes,omitempty"`
	Status                              *EntityStatus        `json:"status,omitempty"`
	UpdatedAt                           *string              `json:"updatedAt,omitempty"`
}
