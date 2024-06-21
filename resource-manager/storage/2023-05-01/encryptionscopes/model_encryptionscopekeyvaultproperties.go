package encryptionscopes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionScopeKeyVaultProperties struct {
	CurrentVersionedKeyIdentifier *string `json:"currentVersionedKeyIdentifier,omitempty"`
	KeyUri                        *string `json:"keyUri,omitempty"`
	LastKeyRotationTimestamp      *string `json:"lastKeyRotationTimestamp,omitempty"`
}
