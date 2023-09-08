package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfiguration struct {
	AllowedGoogleAccountDomains                               *[]string                                                                      `json:"allowedGoogleAccountDomains,omitempty"`
	Assignments                                               *[]DeviceConfigurationAssignment                                               `json:"assignments,omitempty"`
	BlockUnifiedPasswordForWorkProfile                        *bool                                                                          `json:"blockUnifiedPasswordForWorkProfile,omitempty"`
	CreatedDateTime                                           *string                                                                        `json:"createdDateTime,omitempty"`
	Description                                               *string                                                                        `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode               *DeviceManagementApplicabilityRuleDeviceMode                                   `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition                *DeviceManagementApplicabilityRuleOsEdition                                    `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion                *DeviceManagementApplicabilityRuleOsVersion                                    `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                               *[]SettingStateDeviceSummary                                                   `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                                      *DeviceConfigurationDeviceOverview                                             `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                            *[]DeviceConfigurationDeviceStatus                                             `json:"deviceStatuses,omitempty"`
	DisplayName                                               *string                                                                        `json:"displayName,omitempty"`
	GroupAssignments                                          *[]DeviceConfigurationGroupAssignment                                          `json:"groupAssignments,omitempty"`
	Id                                                        *string                                                                        `json:"id,omitempty"`
	LastModifiedDateTime                                      *string                                                                        `json:"lastModifiedDateTime,omitempty"`
	ODataType                                                 *string                                                                        `json:"@odata.type,omitempty"`
	PasswordBlockFaceUnlock                                   *bool                                                                          `json:"passwordBlockFaceUnlock,omitempty"`
	PasswordBlockFingerprintUnlock                            *bool                                                                          `json:"passwordBlockFingerprintUnlock,omitempty"`
	PasswordBlockIrisUnlock                                   *bool                                                                          `json:"passwordBlockIrisUnlock,omitempty"`
	PasswordBlockTrustAgents                                  *bool                                                                          `json:"passwordBlockTrustAgents,omitempty"`
	PasswordExpirationDays                                    *int64                                                                         `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumLength                                     *int64                                                                         `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeScreenTimeout            *int64                                                                         `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasswordPreviousPasswordBlockCount                        *int64                                                                         `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequiredType                                      *AndroidForWorkGeneralDeviceConfigurationPasswordRequiredType                  `json:"passwordRequiredType,omitempty"`
	PasswordSignInFailureCountBeforeFactoryReset              *int64                                                                         `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	RequiredPasswordComplexity                                *AndroidForWorkGeneralDeviceConfigurationRequiredPasswordComplexity            `json:"requiredPasswordComplexity,omitempty"`
	RoleScopeTagIds                                           *[]string                                                                      `json:"roleScopeTagIds,omitempty"`
	SecurityRequireVerifyApps                                 *bool                                                                          `json:"securityRequireVerifyApps,omitempty"`
	SupportsScopeTags                                         *bool                                                                          `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                                        *DeviceConfigurationUserOverview                                               `json:"userStatusOverview,omitempty"`
	UserStatuses                                              *[]DeviceConfigurationUserStatus                                               `json:"userStatuses,omitempty"`
	Version                                                   *int64                                                                         `json:"version,omitempty"`
	VpnAlwaysOnPackageIdentifier                              *string                                                                        `json:"vpnAlwaysOnPackageIdentifier,omitempty"`
	VpnEnableAlwaysOnLockdownMode                             *bool                                                                          `json:"vpnEnableAlwaysOnLockdownMode,omitempty"`
	WorkProfileAccountUse                                     *AndroidForWorkGeneralDeviceConfigurationWorkProfileAccountUse                 `json:"workProfileAccountUse,omitempty"`
	WorkProfileAllowWidgets                                   *bool                                                                          `json:"workProfileAllowWidgets,omitempty"`
	WorkProfileBlockAddingAccounts                            *bool                                                                          `json:"workProfileBlockAddingAccounts,omitempty"`
	WorkProfileBlockCamera                                    *bool                                                                          `json:"workProfileBlockCamera,omitempty"`
	WorkProfileBlockCrossProfileCallerId                      *bool                                                                          `json:"workProfileBlockCrossProfileCallerId,omitempty"`
	WorkProfileBlockCrossProfileContactsSearch                *bool                                                                          `json:"workProfileBlockCrossProfileContactsSearch,omitempty"`
	WorkProfileBlockCrossProfileCopyPaste                     *bool                                                                          `json:"workProfileBlockCrossProfileCopyPaste,omitempty"`
	WorkProfileBlockNotificationsWhileDeviceLocked            *bool                                                                          `json:"workProfileBlockNotificationsWhileDeviceLocked,omitempty"`
	WorkProfileBlockPersonalAppInstallsFromUnknownSources     *bool                                                                          `json:"workProfileBlockPersonalAppInstallsFromUnknownSources,omitempty"`
	WorkProfileBlockScreenCapture                             *bool                                                                          `json:"workProfileBlockScreenCapture,omitempty"`
	WorkProfileBluetoothEnableContactSharing                  *bool                                                                          `json:"workProfileBluetoothEnableContactSharing,omitempty"`
	WorkProfileDataSharingType                                *AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType            `json:"workProfileDataSharingType,omitempty"`
	WorkProfileDefaultAppPermissionPolicy                     *AndroidForWorkGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy `json:"workProfileDefaultAppPermissionPolicy,omitempty"`
	WorkProfilePasswordBlockFaceUnlock                        *bool                                                                          `json:"workProfilePasswordBlockFaceUnlock,omitempty"`
	WorkProfilePasswordBlockFingerprintUnlock                 *bool                                                                          `json:"workProfilePasswordBlockFingerprintUnlock,omitempty"`
	WorkProfilePasswordBlockIrisUnlock                        *bool                                                                          `json:"workProfilePasswordBlockIrisUnlock,omitempty"`
	WorkProfilePasswordBlockTrustAgents                       *bool                                                                          `json:"workProfilePasswordBlockTrustAgents,omitempty"`
	WorkProfilePasswordExpirationDays                         *int64                                                                         `json:"workProfilePasswordExpirationDays,omitempty"`
	WorkProfilePasswordMinLetterCharacters                    *int64                                                                         `json:"workProfilePasswordMinLetterCharacters,omitempty"`
	WorkProfilePasswordMinLowerCaseCharacters                 *int64                                                                         `json:"workProfilePasswordMinLowerCaseCharacters,omitempty"`
	WorkProfilePasswordMinNonLetterCharacters                 *int64                                                                         `json:"workProfilePasswordMinNonLetterCharacters,omitempty"`
	WorkProfilePasswordMinNumericCharacters                   *int64                                                                         `json:"workProfilePasswordMinNumericCharacters,omitempty"`
	WorkProfilePasswordMinSymbolCharacters                    *int64                                                                         `json:"workProfilePasswordMinSymbolCharacters,omitempty"`
	WorkProfilePasswordMinUpperCaseCharacters                 *int64                                                                         `json:"workProfilePasswordMinUpperCaseCharacters,omitempty"`
	WorkProfilePasswordMinimumLength                          *int64                                                                         `json:"workProfilePasswordMinimumLength,omitempty"`
	WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout *int64                                                                         `json:"workProfilePasswordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	WorkProfilePasswordPreviousPasswordBlockCount             *int64                                                                         `json:"workProfilePasswordPreviousPasswordBlockCount,omitempty"`
	WorkProfilePasswordRequiredType                           *AndroidForWorkGeneralDeviceConfigurationWorkProfilePasswordRequiredType       `json:"workProfilePasswordRequiredType,omitempty"`
	WorkProfilePasswordSignInFailureCountBeforeFactoryReset   *int64                                                                         `json:"workProfilePasswordSignInFailureCountBeforeFactoryReset,omitempty"`
	WorkProfileRequirePassword                                *bool                                                                          `json:"workProfileRequirePassword,omitempty"`
	WorkProfileRequiredPasswordComplexity                     *AndroidForWorkGeneralDeviceConfigurationWorkProfileRequiredPasswordComplexity `json:"workProfileRequiredPasswordComplexity,omitempty"`
}
