package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SBSubscriptionProperties struct {
	AccessedAt                                *string                   `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                          *string                   `json:"autoDeleteOnIdle,omitempty"`
	ClientAffineProperties                    *SBClientAffineProperties `json:"clientAffineProperties,omitempty"`
	CountDetails                              *MessageCountDetails      `json:"countDetails,omitempty"`
	CreatedAt                                 *string                   `json:"createdAt,omitempty"`
	DeadLetteringOnFilterEvaluationExceptions *bool                     `json:"deadLetteringOnFilterEvaluationExceptions,omitempty"`
	DeadLetteringOnMessageExpiration          *bool                     `json:"deadLetteringOnMessageExpiration,omitempty"`
	DefaultMessageTimeToLive                  *string                   `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow       *string                   `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations                   *bool                     `json:"enableBatchedOperations,omitempty"`
	ForwardDeadLetteredMessagesTo             *string                   `json:"forwardDeadLetteredMessagesTo,omitempty"`
	ForwardTo                                 *string                   `json:"forwardTo,omitempty"`
	IsClientAffine                            *bool                     `json:"isClientAffine,omitempty"`
	LockDuration                              *string                   `json:"lockDuration,omitempty"`
	MaxDeliveryCount                          *int64                    `json:"maxDeliveryCount,omitempty"`
	MessageCount                              *int64                    `json:"messageCount,omitempty"`
	RequiresSession                           *bool                     `json:"requiresSession,omitempty"`
	Status                                    *EntityStatus             `json:"status,omitempty"`
	UpdatedAt                                 *string                   `json:"updatedAt,omitempty"`
}
