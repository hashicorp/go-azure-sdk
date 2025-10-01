package bms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionDetails struct {
	EncryptionEnabled *bool   `json:"encryptionEnabled,omitempty"`
	KekURL            *string `json:"kekUrl,omitempty"`
	KekVaultId        *string `json:"kekVaultId,omitempty"`
	SecretKeyURL      *string `json:"secretKeyUrl,omitempty"`
	SecretKeyVaultId  *string `json:"secretKeyVaultId,omitempty"`
}
