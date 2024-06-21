package diskencryptionsets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EncryptionSetProperties struct {
	ActiveKey                         *KeyForDiskEncryptionSet   `json:"activeKey,omitempty"`
	AutoKeyRotationError              *ApiError                  `json:"autoKeyRotationError,omitempty"`
	EncryptionType                    *DiskEncryptionSetType     `json:"encryptionType,omitempty"`
	FederatedClientId                 *string                    `json:"federatedClientId,omitempty"`
	LastKeyRotationTimestamp          *string                    `json:"lastKeyRotationTimestamp,omitempty"`
	PreviousKeys                      *[]KeyForDiskEncryptionSet `json:"previousKeys,omitempty"`
	ProvisioningState                 *string                    `json:"provisioningState,omitempty"`
	RotationToLatestKeyVersionEnabled *bool                      `json:"rotationToLatestKeyVersionEnabled,omitempty"`
}
