package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretArchiveReference struct {
	KeyVaultId    *string `json:"keyVaultId,omitempty"`
	SecretName    *string `json:"secretName,omitempty"`
	SecretVersion *string `json:"secretVersion,omitempty"`
}
