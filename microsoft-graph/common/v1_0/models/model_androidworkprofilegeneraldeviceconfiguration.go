package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfiguration struct {
	Assignments                                               *[]DeviceConfigurationAssignment                                                   `json:"assignments,omitempty"`
	CreatedDateTime                                           *string                                                                            `json:"createdDateTime,omitempty"`
	Description                                               *string                                                                            `json:"description,omitempty"`
	DeviceSettingStateSummaries                               *[]SettingStateDeviceSummary                                                       `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                                      *DeviceConfigurationDeviceOverview                                                 `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                            *[]DeviceConfigurationDeviceStatus                                                 `json:"deviceStatuses,omitempty"`
	DisplayName                                               *string                                                                            `json:"displayName,omitempty"`
	Id                                                        *string                                                                            `json:"id,omitempty"`
	LastModifiedDateTime                                      *string                                                                            `json:"lastModifiedDateTime,omitempty"`
	ODataType                                                 *string                                                                            `json:"@odata.type,omitempty"`
	PasswordBlockFingerprintUnlock                            *bool                                                                              `json:"passwordBlockFingerprintUnlock,omitempty"`
	PasswordBlockTrustAgents                                  *bool                                                                              `json:"passwordBlockTrustAgents,omitempty"`
	PasswordExpirationDays                                    *int64                                                                             `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumLength                                     *int64                                                                             `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeScreenTimeout            *int64                                                                             `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasswordPreviousPasswordBlockCount                        *int64                                                                             `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequiredType                                      *AndroidWorkProfileGeneralDeviceConfigurationPasswordRequiredType                  `json:"passwordRequiredType,omitempty"`
	PasswordSignInFailureCountBeforeFactoryReset              *int64                                                                             `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	SecurityRequireVerifyApps                                 *bool                                                                              `json:"securityRequireVerifyApps,omitempty"`
	UserStatusOverview                                        *DeviceConfigurationUserOverview                                                   `json:"userStatusOverview,omitempty"`
	UserStatuses                                              *[]DeviceConfigurationUserStatus                                                   `json:"userStatuses,omitempty"`
	Version                                                   *int64                                                                             `json:"version,omitempty"`
	WorkProfileBlockAddingAccounts                            *bool                                                                              `json:"workProfileBlockAddingAccounts,omitempty"`
	WorkProfileBlockCamera                                    *bool                                                                              `json:"workProfileBlockCamera,omitempty"`
	WorkProfileBlockCrossProfileCallerId                      *bool                                                                              `json:"workProfileBlockCrossProfileCallerId,omitempty"`
	WorkProfileBlockCrossProfileContactsSearch                *bool                                                                              `json:"workProfileBlockCrossProfileContactsSearch,omitempty"`
	WorkProfileBlockCrossProfileCopyPaste                     *bool                                                                              `json:"workProfileBlockCrossProfileCopyPaste,omitempty"`
	WorkProfileBlockNotificationsWhileDeviceLocked            *bool                                                                              `json:"workProfileBlockNotificationsWhileDeviceLocked,omitempty"`
	WorkProfileBlockScreenCapture                             *bool                                                                              `json:"workProfileBlockScreenCapture,omitempty"`
	WorkProfileBluetoothEnableContactSharing                  *bool                                                                              `json:"workProfileBluetoothEnableContactSharing,omitempty"`
	WorkProfileDataSharingType                                *AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType            `json:"workProfileDataSharingType,omitempty"`
	WorkProfileDefaultAppPermissionPolicy                     *AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDefaultAppPermissionPolicy `json:"workProfileDefaultAppPermissionPolicy,omitempty"`
	WorkProfilePasswordBlockFingerprintUnlock                 *bool                                                                              `json:"workProfilePasswordBlockFingerprintUnlock,omitempty"`
	WorkProfilePasswordBlockTrustAgents                       *bool                                                                              `json:"workProfilePasswordBlockTrustAgents,omitempty"`
	WorkProfilePasswordExpirationDays                         *int64                                                                             `json:"workProfilePasswordExpirationDays,omitempty"`
	WorkProfilePasswordMinLetterCharacters                    *int64                                                                             `json:"workProfilePasswordMinLetterCharacters,omitempty"`
	WorkProfilePasswordMinLowerCaseCharacters                 *int64                                                                             `json:"workProfilePasswordMinLowerCaseCharacters,omitempty"`
	WorkProfilePasswordMinNonLetterCharacters                 *int64                                                                             `json:"workProfilePasswordMinNonLetterCharacters,omitempty"`
	WorkProfilePasswordMinNumericCharacters                   *int64                                                                             `json:"workProfilePasswordMinNumericCharacters,omitempty"`
	WorkProfilePasswordMinSymbolCharacters                    *int64                                                                             `json:"workProfilePasswordMinSymbolCharacters,omitempty"`
	WorkProfilePasswordMinUpperCaseCharacters                 *int64                                                                             `json:"workProfilePasswordMinUpperCaseCharacters,omitempty"`
	WorkProfilePasswordMinimumLength                          *int64                                                                             `json:"workProfilePasswordMinimumLength,omitempty"`
	WorkProfilePasswordMinutesOfInactivityBeforeScreenTimeout *int64                                                                             `json:"workProfilePasswordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	WorkProfilePasswordPreviousPasswordBlockCount             *int64                                                                             `json:"workProfilePasswordPreviousPasswordBlockCount,omitempty"`
	WorkProfilePasswordRequiredType                           *AndroidWorkProfileGeneralDeviceConfigurationWorkProfilePasswordRequiredType       `json:"workProfilePasswordRequiredType,omitempty"`
	WorkProfilePasswordSignInFailureCountBeforeFactoryReset   *int64                                                                             `json:"workProfilePasswordSignInFailureCountBeforeFactoryReset,omitempty"`
	WorkProfileRequirePassword                                *bool                                                                              `json:"workProfileRequirePassword,omitempty"`
}
