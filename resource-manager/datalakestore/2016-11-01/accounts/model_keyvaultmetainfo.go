package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyVaultMetaInfo struct {
	EncryptionKeyName    string `json:"encryptionKeyName"`
	EncryptionKeyVersion string `json:"encryptionKeyVersion"`
	KeyVaultResourceId   string `json:"keyVaultResourceId"`
}
