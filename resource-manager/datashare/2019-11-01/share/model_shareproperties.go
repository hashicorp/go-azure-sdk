package share

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShareProperties struct {
	CreatedAt         *string            `json:"createdAt,omitempty"`
	Description       *string            `json:"description,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	ShareKind         *ShareKind         `json:"shareKind,omitempty"`
	Terms             *string            `json:"terms,omitempty"`
	UserEmail         *string            `json:"userEmail,omitempty"`
	UserName          *string            `json:"userName,omitempty"`
}
