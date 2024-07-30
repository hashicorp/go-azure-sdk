package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceEncryptionState struct {
	AdvancedBitLockerStates      *ManagedDeviceEncryptionStateAdvancedBitLockerStates      `json:"advancedBitLockerStates,omitempty"`
	DeviceName                   *string                                                   `json:"deviceName,omitempty"`
	DeviceType                   *ManagedDeviceEncryptionStateDeviceType                   `json:"deviceType,omitempty"`
	EncryptionPolicySettingState *ManagedDeviceEncryptionStateEncryptionPolicySettingState `json:"encryptionPolicySettingState,omitempty"`
	EncryptionReadinessState     *ManagedDeviceEncryptionStateEncryptionReadinessState     `json:"encryptionReadinessState,omitempty"`
	EncryptionState              *ManagedDeviceEncryptionStateEncryptionState              `json:"encryptionState,omitempty"`
	FileVaultStates              *ManagedDeviceEncryptionStateFileVaultStates              `json:"fileVaultStates,omitempty"`
	Id                           *string                                                   `json:"id,omitempty"`
	ODataType                    *string                                                   `json:"@odata.type,omitempty"`
	OsVersion                    *string                                                   `json:"osVersion,omitempty"`
	PolicyDetails                *[]EncryptionReportPolicyDetails                          `json:"policyDetails,omitempty"`
	TpmSpecificationVersion      *string                                                   `json:"tpmSpecificationVersion,omitempty"`
	UserPrincipalName            *string                                                   `json:"userPrincipalName,omitempty"`
}
