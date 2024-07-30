package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfiguration struct {
	Assignments                       *[]EnrollmentConfigurationAssignment                                                   `json:"assignments,omitempty"`
	CreatedDateTime                   *string                                                                                `json:"createdDateTime,omitempty"`
	Description                       *string                                                                                `json:"description,omitempty"`
	DeviceEnrollmentConfigurationType *DeviceEnrollmentWindowsHelloForBusinessConfigurationDeviceEnrollmentConfigurationType `json:"deviceEnrollmentConfigurationType,omitempty"`
	DisplayName                       *string                                                                                `json:"displayName,omitempty"`
	EnhancedBiometricsState           *DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState           `json:"enhancedBiometricsState,omitempty"`
	EnhancedSignInSecurity            *int64                                                                                 `json:"enhancedSignInSecurity,omitempty"`
	Id                                *string                                                                                `json:"id,omitempty"`
	LastModifiedDateTime              *string                                                                                `json:"lastModifiedDateTime,omitempty"`
	ODataType                         *string                                                                                `json:"@odata.type,omitempty"`
	PinExpirationInDays               *int64                                                                                 `json:"pinExpirationInDays,omitempty"`
	PinLowercaseCharactersUsage       *DeviceEnrollmentWindowsHelloForBusinessConfigurationPinLowercaseCharactersUsage       `json:"pinLowercaseCharactersUsage,omitempty"`
	PinMaximumLength                  *int64                                                                                 `json:"pinMaximumLength,omitempty"`
	PinMinimumLength                  *int64                                                                                 `json:"pinMinimumLength,omitempty"`
	PinPreviousBlockCount             *int64                                                                                 `json:"pinPreviousBlockCount,omitempty"`
	PinSpecialCharactersUsage         *DeviceEnrollmentWindowsHelloForBusinessConfigurationPinSpecialCharactersUsage         `json:"pinSpecialCharactersUsage,omitempty"`
	PinUppercaseCharactersUsage       *DeviceEnrollmentWindowsHelloForBusinessConfigurationPinUppercaseCharactersUsage       `json:"pinUppercaseCharactersUsage,omitempty"`
	Priority                          *int64                                                                                 `json:"priority,omitempty"`
	RemotePassportEnabled             *bool                                                                                  `json:"remotePassportEnabled,omitempty"`
	RoleScopeTagIds                   *[]string                                                                              `json:"roleScopeTagIds,omitempty"`
	SecurityDeviceRequired            *bool                                                                                  `json:"securityDeviceRequired,omitempty"`
	SecurityKeyForSignIn              *DeviceEnrollmentWindowsHelloForBusinessConfigurationSecurityKeyForSignIn              `json:"securityKeyForSignIn,omitempty"`
	State                             *DeviceEnrollmentWindowsHelloForBusinessConfigurationState                             `json:"state,omitempty"`
	UnlockWithBiometricsEnabled       *bool                                                                                  `json:"unlockWithBiometricsEnabled,omitempty"`
	Version                           *int64                                                                                 `json:"version,omitempty"`
}
