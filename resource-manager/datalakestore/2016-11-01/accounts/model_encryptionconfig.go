package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionConfig struct {
	KeyVaultMetaInfo *KeyVaultMetaInfo    `json:"keyVaultMetaInfo"`
	Type             EncryptionConfigType `json:"type"`
}
