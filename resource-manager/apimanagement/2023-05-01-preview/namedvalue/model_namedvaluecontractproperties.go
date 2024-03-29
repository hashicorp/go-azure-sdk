package namedvalue

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamedValueContractProperties struct {
	DisplayName       string                      `json:"displayName"`
	KeyVault          *KeyVaultContractProperties `json:"keyVault,omitempty"`
	ProvisioningState *string                     `json:"provisioningState,omitempty"`
	Secret            *bool                       `json:"secret,omitempty"`
	Tags              *[]string                   `json:"tags,omitempty"`
	Value             *string                     `json:"value,omitempty"`
}
