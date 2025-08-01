package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiskEncryptionSettings struct {
	DiskEncryptionKey *KeyVaultSecretReference `json:"diskEncryptionKey,omitempty"`
	Enabled           *bool                    `json:"enabled,omitempty"`
	KeyEncryptionKey  *KeyVaultKeyReference    `json:"keyEncryptionKey,omitempty"`
}
