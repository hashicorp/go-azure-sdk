package registries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyVaultProperties struct {
	Identity                 *string `json:"identity,omitempty"`
	KeyIdentifier            *string `json:"keyIdentifier,omitempty"`
	KeyRotationEnabled       *bool   `json:"keyRotationEnabled,omitempty"`
	LastKeyRotationTimestamp *string `json:"lastKeyRotationTimestamp,omitempty"`
	VersionedKeyIdentifier   *string `json:"versionedKeyIdentifier,omitempty"`
}
