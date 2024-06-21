package users

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserProperties struct {
	CreatedDate       *string          `json:"createdDate,omitempty"`
	Identity          *UserIdentity    `json:"identity,omitempty"`
	ProvisioningState *string          `json:"provisioningState,omitempty"`
	SecretStore       *UserSecretStore `json:"secretStore,omitempty"`
	UniqueIdentifier  *string          `json:"uniqueIdentifier,omitempty"`
}
