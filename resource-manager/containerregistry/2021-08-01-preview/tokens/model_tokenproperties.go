package tokens

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenProperties struct {
	CreationDate      *string                     `json:"creationDate,omitempty"`
	Credentials       *TokenCredentialsProperties `json:"credentials,omitempty"`
	ProvisioningState *ProvisioningState          `json:"provisioningState,omitempty"`
	ScopeMapId        *string                     `json:"scopeMapId,omitempty"`
	Status            *TokenStatus                `json:"status,omitempty"`
}
