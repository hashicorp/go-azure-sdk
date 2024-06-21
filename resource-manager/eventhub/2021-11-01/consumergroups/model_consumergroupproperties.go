package consumergroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConsumerGroupProperties struct {
	CreatedAt    *string `json:"createdAt,omitempty"`
	UpdatedAt    *string `json:"updatedAt,omitempty"`
	UserMetadata *string `json:"userMetadata,omitempty"`
}
