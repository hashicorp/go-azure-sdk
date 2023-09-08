package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerFixedDrivePolicy struct {
	EncryptionMethod                *BitLockerFixedDrivePolicyEncryptionMethod `json:"encryptionMethod,omitempty"`
	ODataType                       *string                                    `json:"@odata.type,omitempty"`
	RecoveryOptions                 *BitLockerRecoveryOptions                  `json:"recoveryOptions,omitempty"`
	RequireEncryptionForWriteAccess *bool                                      `json:"requireEncryptionForWriteAccess,omitempty"`
}
