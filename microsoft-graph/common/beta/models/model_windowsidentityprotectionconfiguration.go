package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsIdentityProtectionConfiguration struct {
	Assignments                                  *[]DeviceConfigurationAssignment                                   `json:"assignments,omitempty"`
	CreatedDateTime                              *string                                                            `json:"createdDateTime,omitempty"`
	Description                                  *string                                                            `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode  *DeviceManagementApplicabilityRuleDeviceMode                       `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition   *DeviceManagementApplicabilityRuleOsEdition                        `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion   *DeviceManagementApplicabilityRuleOsVersion                        `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                  *[]SettingStateDeviceSummary                                       `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                         *DeviceConfigurationDeviceOverview                                 `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                               *[]DeviceConfigurationDeviceStatus                                 `json:"deviceStatuses,omitempty"`
	DisplayName                                  *string                                                            `json:"displayName,omitempty"`
	EnhancedAntiSpoofingForFacialFeaturesEnabled *bool                                                              `json:"enhancedAntiSpoofingForFacialFeaturesEnabled,omitempty"`
	GroupAssignments                             *[]DeviceConfigurationGroupAssignment                              `json:"groupAssignments,omitempty"`
	Id                                           *string                                                            `json:"id,omitempty"`
	LastModifiedDateTime                         *string                                                            `json:"lastModifiedDateTime,omitempty"`
	ODataType                                    *string                                                            `json:"@odata.type,omitempty"`
	PinExpirationInDays                          *int64                                                             `json:"pinExpirationInDays,omitempty"`
	PinLowercaseCharactersUsage                  *WindowsIdentityProtectionConfigurationPinLowercaseCharactersUsage `json:"pinLowercaseCharactersUsage,omitempty"`
	PinMaximumLength                             *int64                                                             `json:"pinMaximumLength,omitempty"`
	PinMinimumLength                             *int64                                                             `json:"pinMinimumLength,omitempty"`
	PinPreviousBlockCount                        *int64                                                             `json:"pinPreviousBlockCount,omitempty"`
	PinRecoveryEnabled                           *bool                                                              `json:"pinRecoveryEnabled,omitempty"`
	PinSpecialCharactersUsage                    *WindowsIdentityProtectionConfigurationPinSpecialCharactersUsage   `json:"pinSpecialCharactersUsage,omitempty"`
	PinUppercaseCharactersUsage                  *WindowsIdentityProtectionConfigurationPinUppercaseCharactersUsage `json:"pinUppercaseCharactersUsage,omitempty"`
	RoleScopeTagIds                              *[]string                                                          `json:"roleScopeTagIds,omitempty"`
	SecurityDeviceRequired                       *bool                                                              `json:"securityDeviceRequired,omitempty"`
	SupportsScopeTags                            *bool                                                              `json:"supportsScopeTags,omitempty"`
	UnlockWithBiometricsEnabled                  *bool                                                              `json:"unlockWithBiometricsEnabled,omitempty"`
	UseCertificatesForOnPremisesAuthEnabled      *bool                                                              `json:"useCertificatesForOnPremisesAuthEnabled,omitempty"`
	UseSecurityKeyForSignin                      *bool                                                              `json:"useSecurityKeyForSignin,omitempty"`
	UserStatusOverview                           *DeviceConfigurationUserOverview                                   `json:"userStatusOverview,omitempty"`
	UserStatuses                                 *[]DeviceConfigurationUserStatus                                   `json:"userStatuses,omitempty"`
	Version                                      *int64                                                             `json:"version,omitempty"`
	WindowsHelloForBusinessBlocked               *bool                                                              `json:"windowsHelloForBusinessBlocked,omitempty"`
}
