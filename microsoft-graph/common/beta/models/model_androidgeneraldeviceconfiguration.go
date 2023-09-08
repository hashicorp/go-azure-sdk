package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidGeneralDeviceConfiguration struct {
	AppsBlockClipboardSharing                      *bool                                                        `json:"appsBlockClipboardSharing,omitempty"`
	AppsBlockCopyPaste                             *bool                                                        `json:"appsBlockCopyPaste,omitempty"`
	AppsBlockYouTube                               *bool                                                        `json:"appsBlockYouTube,omitempty"`
	AppsHideList                                   *[]AppListItem                                               `json:"appsHideList,omitempty"`
	AppsInstallAllowList                           *[]AppListItem                                               `json:"appsInstallAllowList,omitempty"`
	AppsLaunchBlockList                            *[]AppListItem                                               `json:"appsLaunchBlockList,omitempty"`
	Assignments                                    *[]DeviceConfigurationAssignment                             `json:"assignments,omitempty"`
	BluetoothBlocked                               *bool                                                        `json:"bluetoothBlocked,omitempty"`
	CameraBlocked                                  *bool                                                        `json:"cameraBlocked,omitempty"`
	CellularBlockDataRoaming                       *bool                                                        `json:"cellularBlockDataRoaming,omitempty"`
	CellularBlockMessaging                         *bool                                                        `json:"cellularBlockMessaging,omitempty"`
	CellularBlockVoiceRoaming                      *bool                                                        `json:"cellularBlockVoiceRoaming,omitempty"`
	CellularBlockWiFiTethering                     *bool                                                        `json:"cellularBlockWiFiTethering,omitempty"`
	CompliantAppListType                           *AndroidGeneralDeviceConfigurationCompliantAppListType       `json:"compliantAppListType,omitempty"`
	CompliantAppsList                              *[]AppListItem                                               `json:"compliantAppsList,omitempty"`
	CreatedDateTime                                *string                                                      `json:"createdDateTime,omitempty"`
	DateAndTimeBlockChanges                        *bool                                                        `json:"dateAndTimeBlockChanges,omitempty"`
	Description                                    *string                                                      `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode    *DeviceManagementApplicabilityRuleDeviceMode                 `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition     *DeviceManagementApplicabilityRuleOsEdition                  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion     *DeviceManagementApplicabilityRuleOsVersion                  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                    *[]SettingStateDeviceSummary                                 `json:"deviceSettingStateSummaries,omitempty"`
	DeviceSharingAllowed                           *bool                                                        `json:"deviceSharingAllowed,omitempty"`
	DeviceStatusOverview                           *DeviceConfigurationDeviceOverview                           `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                 *[]DeviceConfigurationDeviceStatus                           `json:"deviceStatuses,omitempty"`
	DiagnosticDataBlockSubmission                  *bool                                                        `json:"diagnosticDataBlockSubmission,omitempty"`
	DisplayName                                    *string                                                      `json:"displayName,omitempty"`
	FactoryResetBlocked                            *bool                                                        `json:"factoryResetBlocked,omitempty"`
	GoogleAccountBlockAutoSync                     *bool                                                        `json:"googleAccountBlockAutoSync,omitempty"`
	GooglePlayStoreBlocked                         *bool                                                        `json:"googlePlayStoreBlocked,omitempty"`
	GroupAssignments                               *[]DeviceConfigurationGroupAssignment                        `json:"groupAssignments,omitempty"`
	Id                                             *string                                                      `json:"id,omitempty"`
	KioskModeApps                                  *[]AppListItem                                               `json:"kioskModeApps,omitempty"`
	KioskModeBlockSleepButton                      *bool                                                        `json:"kioskModeBlockSleepButton,omitempty"`
	KioskModeBlockVolumeButtons                    *bool                                                        `json:"kioskModeBlockVolumeButtons,omitempty"`
	LastModifiedDateTime                           *string                                                      `json:"lastModifiedDateTime,omitempty"`
	LocationServicesBlocked                        *bool                                                        `json:"locationServicesBlocked,omitempty"`
	NfcBlocked                                     *bool                                                        `json:"nfcBlocked,omitempty"`
	ODataType                                      *string                                                      `json:"@odata.type,omitempty"`
	PasswordBlockFingerprintUnlock                 *bool                                                        `json:"passwordBlockFingerprintUnlock,omitempty"`
	PasswordBlockTrustAgents                       *bool                                                        `json:"passwordBlockTrustAgents,omitempty"`
	PasswordExpirationDays                         *int64                                                       `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumLength                          *int64                                                       `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeScreenTimeout *int64                                                       `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasswordPreviousPasswordBlockCount             *int64                                                       `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequired                               *bool                                                        `json:"passwordRequired,omitempty"`
	PasswordRequiredType                           *AndroidGeneralDeviceConfigurationPasswordRequiredType       `json:"passwordRequiredType,omitempty"`
	PasswordSignInFailureCountBeforeFactoryReset   *int64                                                       `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	PowerOffBlocked                                *bool                                                        `json:"powerOffBlocked,omitempty"`
	RequiredPasswordComplexity                     *AndroidGeneralDeviceConfigurationRequiredPasswordComplexity `json:"requiredPasswordComplexity,omitempty"`
	RoleScopeTagIds                                *[]string                                                    `json:"roleScopeTagIds,omitempty"`
	ScreenCaptureBlocked                           *bool                                                        `json:"screenCaptureBlocked,omitempty"`
	SecurityRequireVerifyApps                      *bool                                                        `json:"securityRequireVerifyApps,omitempty"`
	StorageBlockGoogleBackup                       *bool                                                        `json:"storageBlockGoogleBackup,omitempty"`
	StorageBlockRemovableStorage                   *bool                                                        `json:"storageBlockRemovableStorage,omitempty"`
	StorageRequireDeviceEncryption                 *bool                                                        `json:"storageRequireDeviceEncryption,omitempty"`
	StorageRequireRemovableStorageEncryption       *bool                                                        `json:"storageRequireRemovableStorageEncryption,omitempty"`
	SupportsScopeTags                              *bool                                                        `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                             *DeviceConfigurationUserOverview                             `json:"userStatusOverview,omitempty"`
	UserStatuses                                   *[]DeviceConfigurationUserStatus                             `json:"userStatuses,omitempty"`
	Version                                        *int64                                                       `json:"version,omitempty"`
	VoiceAssistantBlocked                          *bool                                                        `json:"voiceAssistantBlocked,omitempty"`
	VoiceDialingBlocked                            *bool                                                        `json:"voiceDialingBlocked,omitempty"`
	WebBrowserBlockAutofill                        *bool                                                        `json:"webBrowserBlockAutofill,omitempty"`
	WebBrowserBlockJavaScript                      *bool                                                        `json:"webBrowserBlockJavaScript,omitempty"`
	WebBrowserBlockPopups                          *bool                                                        `json:"webBrowserBlockPopups,omitempty"`
	WebBrowserBlocked                              *bool                                                        `json:"webBrowserBlocked,omitempty"`
	WebBrowserCookieSettings                       *AndroidGeneralDeviceConfigurationWebBrowserCookieSettings   `json:"webBrowserCookieSettings,omitempty"`
	WiFiBlocked                                    *bool                                                        `json:"wiFiBlocked,omitempty"`
}
