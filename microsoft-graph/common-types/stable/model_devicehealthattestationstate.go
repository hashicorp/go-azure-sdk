package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationState struct {
	AttestationIdentityKey                   *string `json:"attestationIdentityKey,omitempty"`
	BitLockerStatus                          *string `json:"bitLockerStatus,omitempty"`
	BootAppSecurityVersion                   *string `json:"bootAppSecurityVersion,omitempty"`
	BootDebugging                            *string `json:"bootDebugging,omitempty"`
	BootManagerSecurityVersion               *string `json:"bootManagerSecurityVersion,omitempty"`
	BootManagerVersion                       *string `json:"bootManagerVersion,omitempty"`
	BootRevisionListInfo                     *string `json:"bootRevisionListInfo,omitempty"`
	CodeIntegrity                            *string `json:"codeIntegrity,omitempty"`
	CodeIntegrityCheckVersion                *string `json:"codeIntegrityCheckVersion,omitempty"`
	CodeIntegrityPolicy                      *string `json:"codeIntegrityPolicy,omitempty"`
	ContentNamespaceUrl                      *string `json:"contentNamespaceUrl,omitempty"`
	ContentVersion                           *string `json:"contentVersion,omitempty"`
	DataExcutionPolicy                       *string `json:"dataExcutionPolicy,omitempty"`
	DeviceHealthAttestationStatus            *string `json:"deviceHealthAttestationStatus,omitempty"`
	EarlyLaunchAntiMalwareDriverProtection   *string `json:"earlyLaunchAntiMalwareDriverProtection,omitempty"`
	HealthAttestationSupportedStatus         *string `json:"healthAttestationSupportedStatus,omitempty"`
	HealthStatusMismatchInfo                 *string `json:"healthStatusMismatchInfo,omitempty"`
	IssuedDateTime                           *string `json:"issuedDateTime,omitempty"`
	LastUpdateDateTime                       *string `json:"lastUpdateDateTime,omitempty"`
	ODataType                                *string `json:"@odata.type,omitempty"`
	OperatingSystemKernelDebugging           *string `json:"operatingSystemKernelDebugging,omitempty"`
	OperatingSystemRevListInfo               *string `json:"operatingSystemRevListInfo,omitempty"`
	Pcr0                                     *string `json:"pcr0,omitempty"`
	PcrHashAlgorithm                         *string `json:"pcrHashAlgorithm,omitempty"`
	ResetCount                               *int64  `json:"resetCount,omitempty"`
	RestartCount                             *int64  `json:"restartCount,omitempty"`
	SafeMode                                 *string `json:"safeMode,omitempty"`
	SecureBoot                               *string `json:"secureBoot,omitempty"`
	SecureBootConfigurationPolicyFingerPrint *string `json:"secureBootConfigurationPolicyFingerPrint,omitempty"`
	TestSigning                              *string `json:"testSigning,omitempty"`
	TpmVersion                               *string `json:"tpmVersion,omitempty"`
	VirtualSecureMode                        *string `json:"virtualSecureMode,omitempty"`
	WindowsPE                                *string `json:"windowsPE,omitempty"`
}
