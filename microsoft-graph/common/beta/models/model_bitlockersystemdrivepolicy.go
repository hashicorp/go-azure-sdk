package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BitLockerSystemDrivePolicy struct {
	EncryptionMethod                         *BitLockerSystemDrivePolicyEncryptionMethod                       `json:"encryptionMethod,omitempty"`
	MinimumPinLength                         *int64                                                            `json:"minimumPinLength,omitempty"`
	ODataType                                *string                                                           `json:"@odata.type,omitempty"`
	PrebootRecoveryEnableMessageAndUrl       *bool                                                             `json:"prebootRecoveryEnableMessageAndUrl,omitempty"`
	PrebootRecoveryMessage                   *string                                                           `json:"prebootRecoveryMessage,omitempty"`
	PrebootRecoveryUrl                       *string                                                           `json:"prebootRecoveryUrl,omitempty"`
	RecoveryOptions                          *BitLockerRecoveryOptions                                         `json:"recoveryOptions,omitempty"`
	StartupAuthenticationBlockWithoutTpmChip *bool                                                             `json:"startupAuthenticationBlockWithoutTpmChip,omitempty"`
	StartupAuthenticationRequired            *bool                                                             `json:"startupAuthenticationRequired,omitempty"`
	StartupAuthenticationTpmKeyUsage         *BitLockerSystemDrivePolicyStartupAuthenticationTpmKeyUsage       `json:"startupAuthenticationTpmKeyUsage,omitempty"`
	StartupAuthenticationTpmPinAndKeyUsage   *BitLockerSystemDrivePolicyStartupAuthenticationTpmPinAndKeyUsage `json:"startupAuthenticationTpmPinAndKeyUsage,omitempty"`
	StartupAuthenticationTpmPinUsage         *BitLockerSystemDrivePolicyStartupAuthenticationTpmPinUsage       `json:"startupAuthenticationTpmPinUsage,omitempty"`
	StartupAuthenticationTpmUsage            *BitLockerSystemDrivePolicyStartupAuthenticationTpmUsage          `json:"startupAuthenticationTpmUsage,omitempty"`
}
