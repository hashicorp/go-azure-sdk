package users

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecretStore struct {
	KeyVaultId  *string `json:"keyVaultId,omitempty"`
	KeyVaultUri *string `json:"keyVaultUri,omitempty"`
}
