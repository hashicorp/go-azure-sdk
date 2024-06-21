package topics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SBTopicProperties struct {
	AccessedAt                          *string              `json:"accessedAt,omitempty"`
	AutoDeleteOnIdle                    *string              `json:"autoDeleteOnIdle,omitempty"`
	CountDetails                        *MessageCountDetails `json:"countDetails,omitempty"`
	CreatedAt                           *string              `json:"createdAt,omitempty"`
	DefaultMessageTimeToLive            *string              `json:"defaultMessageTimeToLive,omitempty"`
	DuplicateDetectionHistoryTimeWindow *string              `json:"duplicateDetectionHistoryTimeWindow,omitempty"`
	EnableBatchedOperations             *bool                `json:"enableBatchedOperations,omitempty"`
	EnableExpress                       *bool                `json:"enableExpress,omitempty"`
	EnablePartitioning                  *bool                `json:"enablePartitioning,omitempty"`
	MaxMessageSizeInKilobytes           *int64               `json:"maxMessageSizeInKilobytes,omitempty"`
	MaxSizeInMegabytes                  *int64               `json:"maxSizeInMegabytes,omitempty"`
	RequiresDuplicateDetection          *bool                `json:"requiresDuplicateDetection,omitempty"`
	SizeInBytes                         *int64               `json:"sizeInBytes,omitempty"`
	Status                              *EntityStatus        `json:"status,omitempty"`
	SubscriptionCount                   *int64               `json:"subscriptionCount,omitempty"`
	SupportOrdering                     *bool                `json:"supportOrdering,omitempty"`
	UpdatedAt                           *string              `json:"updatedAt,omitempty"`
}
