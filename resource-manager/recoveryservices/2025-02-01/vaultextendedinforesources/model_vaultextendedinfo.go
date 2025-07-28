package vaultextendedinforesources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VaultExtendedInfo struct {
	Algorithm               *string `json:"algorithm,omitempty"`
	EncryptionKey           *string `json:"encryptionKey,omitempty"`
	EncryptionKeyThumbprint *string `json:"encryptionKeyThumbprint,omitempty"`
	IntegrityKey            *string `json:"integrityKey,omitempty"`
}
