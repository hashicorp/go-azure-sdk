package account

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountProperties struct {
	CreatedAt         *string            `json:"createdAt,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	UserEmail         *string            `json:"userEmail,omitempty"`
	UserName          *string            `json:"userName,omitempty"`
}
