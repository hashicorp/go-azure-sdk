package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationStateOperationPredicate struct {
	AttestationIdentityKey                   *string
	BitLockerStatus                          *string
	BootAppSecurityVersion                   *string
	BootDebugging                            *string
	BootManagerSecurityVersion               *string
	BootManagerVersion                       *string
	BootRevisionListInfo                     *string
	CodeIntegrity                            *string
	CodeIntegrityCheckVersion                *string
	CodeIntegrityPolicy                      *string
	ContentNamespaceUrl                      *string
	ContentVersion                           *string
	DataExcutionPolicy                       *string
	DeviceHealthAttestationStatus            *string
	EarlyLaunchAntiMalwareDriverProtection   *string
	HealthAttestationSupportedStatus         *string
	HealthStatusMismatchInfo                 *string
	IssuedDateTime                           *string
	LastUpdateDateTime                       *string
	ODataType                                *string
	OperatingSystemKernelDebugging           *string
	OperatingSystemRevListInfo               *string
	Pcr0                                     *string
	PcrHashAlgorithm                         *string
	ResetCount                               *int64
	RestartCount                             *int64
	SafeMode                                 *string
	SecureBoot                               *string
	SecureBootConfigurationPolicyFingerPrint *string
	TestSigning                              *string
	TpmVersion                               *string
	VirtualSecureMode                        *string
	WindowsPE                                *string
}

func (p DeviceHealthAttestationStateOperationPredicate) Matches(input DeviceHealthAttestationState) bool {

	if p.AttestationIdentityKey != nil && (input.AttestationIdentityKey == nil || *p.AttestationIdentityKey != *input.AttestationIdentityKey) {
		return false
	}

	if p.BitLockerStatus != nil && (input.BitLockerStatus == nil || *p.BitLockerStatus != *input.BitLockerStatus) {
		return false
	}

	if p.BootAppSecurityVersion != nil && (input.BootAppSecurityVersion == nil || *p.BootAppSecurityVersion != *input.BootAppSecurityVersion) {
		return false
	}

	if p.BootDebugging != nil && (input.BootDebugging == nil || *p.BootDebugging != *input.BootDebugging) {
		return false
	}

	if p.BootManagerSecurityVersion != nil && (input.BootManagerSecurityVersion == nil || *p.BootManagerSecurityVersion != *input.BootManagerSecurityVersion) {
		return false
	}

	if p.BootManagerVersion != nil && (input.BootManagerVersion == nil || *p.BootManagerVersion != *input.BootManagerVersion) {
		return false
	}

	if p.BootRevisionListInfo != nil && (input.BootRevisionListInfo == nil || *p.BootRevisionListInfo != *input.BootRevisionListInfo) {
		return false
	}

	if p.CodeIntegrity != nil && (input.CodeIntegrity == nil || *p.CodeIntegrity != *input.CodeIntegrity) {
		return false
	}

	if p.CodeIntegrityCheckVersion != nil && (input.CodeIntegrityCheckVersion == nil || *p.CodeIntegrityCheckVersion != *input.CodeIntegrityCheckVersion) {
		return false
	}

	if p.CodeIntegrityPolicy != nil && (input.CodeIntegrityPolicy == nil || *p.CodeIntegrityPolicy != *input.CodeIntegrityPolicy) {
		return false
	}

	if p.ContentNamespaceUrl != nil && (input.ContentNamespaceUrl == nil || *p.ContentNamespaceUrl != *input.ContentNamespaceUrl) {
		return false
	}

	if p.ContentVersion != nil && (input.ContentVersion == nil || *p.ContentVersion != *input.ContentVersion) {
		return false
	}

	if p.DataExcutionPolicy != nil && (input.DataExcutionPolicy == nil || *p.DataExcutionPolicy != *input.DataExcutionPolicy) {
		return false
	}

	if p.DeviceHealthAttestationStatus != nil && (input.DeviceHealthAttestationStatus == nil || *p.DeviceHealthAttestationStatus != *input.DeviceHealthAttestationStatus) {
		return false
	}

	if p.EarlyLaunchAntiMalwareDriverProtection != nil && (input.EarlyLaunchAntiMalwareDriverProtection == nil || *p.EarlyLaunchAntiMalwareDriverProtection != *input.EarlyLaunchAntiMalwareDriverProtection) {
		return false
	}

	if p.HealthAttestationSupportedStatus != nil && (input.HealthAttestationSupportedStatus == nil || *p.HealthAttestationSupportedStatus != *input.HealthAttestationSupportedStatus) {
		return false
	}

	if p.HealthStatusMismatchInfo != nil && (input.HealthStatusMismatchInfo == nil || *p.HealthStatusMismatchInfo != *input.HealthStatusMismatchInfo) {
		return false
	}

	if p.IssuedDateTime != nil && (input.IssuedDateTime == nil || *p.IssuedDateTime != *input.IssuedDateTime) {
		return false
	}

	if p.LastUpdateDateTime != nil && (input.LastUpdateDateTime == nil || *p.LastUpdateDateTime != *input.LastUpdateDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OperatingSystemKernelDebugging != nil && (input.OperatingSystemKernelDebugging == nil || *p.OperatingSystemKernelDebugging != *input.OperatingSystemKernelDebugging) {
		return false
	}

	if p.OperatingSystemRevListInfo != nil && (input.OperatingSystemRevListInfo == nil || *p.OperatingSystemRevListInfo != *input.OperatingSystemRevListInfo) {
		return false
	}

	if p.Pcr0 != nil && (input.Pcr0 == nil || *p.Pcr0 != *input.Pcr0) {
		return false
	}

	if p.PcrHashAlgorithm != nil && (input.PcrHashAlgorithm == nil || *p.PcrHashAlgorithm != *input.PcrHashAlgorithm) {
		return false
	}

	if p.ResetCount != nil && (input.ResetCount == nil || *p.ResetCount != *input.ResetCount) {
		return false
	}

	if p.RestartCount != nil && (input.RestartCount == nil || *p.RestartCount != *input.RestartCount) {
		return false
	}

	if p.SafeMode != nil && (input.SafeMode == nil || *p.SafeMode != *input.SafeMode) {
		return false
	}

	if p.SecureBoot != nil && (input.SecureBoot == nil || *p.SecureBoot != *input.SecureBoot) {
		return false
	}

	if p.SecureBootConfigurationPolicyFingerPrint != nil && (input.SecureBootConfigurationPolicyFingerPrint == nil || *p.SecureBootConfigurationPolicyFingerPrint != *input.SecureBootConfigurationPolicyFingerPrint) {
		return false
	}

	if p.TestSigning != nil && (input.TestSigning == nil || *p.TestSigning != *input.TestSigning) {
		return false
	}

	if p.TpmVersion != nil && (input.TpmVersion == nil || *p.TpmVersion != *input.TpmVersion) {
		return false
	}

	if p.VirtualSecureMode != nil && (input.VirtualSecureMode == nil || *p.VirtualSecureMode != *input.VirtualSecureMode) {
		return false
	}

	if p.WindowsPE != nil && (input.WindowsPE == nil || *p.WindowsPE != *input.WindowsPE) {
		return false
	}

	return true
}
