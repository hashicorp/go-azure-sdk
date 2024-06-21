package credentialsets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialSetProperties struct {
	AuthCredentials   *[]AuthCredential  `json:"authCredentials,omitempty"`
	CreationDate      *string            `json:"creationDate,omitempty"`
	LoginServer       *string            `json:"loginServer,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}
