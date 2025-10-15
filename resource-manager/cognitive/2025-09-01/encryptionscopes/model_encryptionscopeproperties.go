package encryptionscopes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionScopeProperties struct {
	KeySource          *KeySource                        `json:"keySource,omitempty"`
	KeyVaultProperties *KeyVaultProperties               `json:"keyVaultProperties,omitempty"`
	ProvisioningState  *EncryptionScopeProvisioningState `json:"provisioningState,omitempty"`
	State              *EncryptionScopeState             `json:"state,omitempty"`
}
