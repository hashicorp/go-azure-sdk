package artifact

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretValueReference struct {
	KeyVault      KeyVaultReference `json:"keyVault"`
	SecretName    string            `json:"secretName"`
	SecretVersion *string           `json:"secretVersion,omitempty"`
}
