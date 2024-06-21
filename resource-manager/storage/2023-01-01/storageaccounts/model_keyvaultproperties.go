package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyVaultProperties struct {
	CurrentVersionedKeyExpirationTimestamp *string `json:"currentVersionedKeyExpirationTimestamp,omitempty"`
	CurrentVersionedKeyIdentifier          *string `json:"currentVersionedKeyIdentifier,omitempty"`
	Keyname                                *string `json:"keyname,omitempty"`
	Keyvaulturi                            *string `json:"keyvaulturi,omitempty"`
	Keyversion                             *string `json:"keyversion,omitempty"`
	LastKeyRotationTimestamp               *string `json:"lastKeyRotationTimestamp,omitempty"`
}
