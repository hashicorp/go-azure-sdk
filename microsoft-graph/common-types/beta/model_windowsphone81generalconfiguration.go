package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81GeneralConfiguration struct {
	ApplyOnlyToWindowsPhone81                      *bool                                                   `json:"applyOnlyToWindowsPhone81,omitempty"`
	AppsBlockCopyPaste                             *bool                                                   `json:"appsBlockCopyPaste,omitempty"`
	Assignments                                    *[]DeviceConfigurationAssignment                        `json:"assignments,omitempty"`
	BluetoothBlocked                               *bool                                                   `json:"bluetoothBlocked,omitempty"`
	CameraBlocked                                  *bool                                                   `json:"cameraBlocked,omitempty"`
	CellularBlockWifiTethering                     *bool                                                   `json:"cellularBlockWifiTethering,omitempty"`
	CompliantAppListType                           *WindowsPhone81GeneralConfigurationCompliantAppListType `json:"compliantAppListType,omitempty"`
	CompliantAppsList                              *[]AppListItem                                          `json:"compliantAppsList,omitempty"`
	CreatedDateTime                                *string                                                 `json:"createdDateTime,omitempty"`
	Description                                    *string                                                 `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode    *DeviceManagementApplicabilityRuleDeviceMode            `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition     *DeviceManagementApplicabilityRuleOsEdition             `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion     *DeviceManagementApplicabilityRuleOsVersion             `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                    *[]SettingStateDeviceSummary                            `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                           *DeviceConfigurationDeviceOverview                      `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                 *[]DeviceConfigurationDeviceStatus                      `json:"deviceStatuses,omitempty"`
	DiagnosticDataBlockSubmission                  *bool                                                   `json:"diagnosticDataBlockSubmission,omitempty"`
	DisplayName                                    *string                                                 `json:"displayName,omitempty"`
	EmailBlockAddingAccounts                       *bool                                                   `json:"emailBlockAddingAccounts,omitempty"`
	GroupAssignments                               *[]DeviceConfigurationGroupAssignment                   `json:"groupAssignments,omitempty"`
	Id                                             *string                                                 `json:"id,omitempty"`
	LastModifiedDateTime                           *string                                                 `json:"lastModifiedDateTime,omitempty"`
	LocationServicesBlocked                        *bool                                                   `json:"locationServicesBlocked,omitempty"`
	MicrosoftAccountBlocked                        *bool                                                   `json:"microsoftAccountBlocked,omitempty"`
	NfcBlocked                                     *bool                                                   `json:"nfcBlocked,omitempty"`
	ODataType                                      *string                                                 `json:"@odata.type,omitempty"`
	PasswordBlockSimple                            *bool                                                   `json:"passwordBlockSimple,omitempty"`
	PasswordExpirationDays                         *int64                                                  `json:"passwordExpirationDays,omitempty"`
	PasswordMinimumCharacterSetCount               *int64                                                  `json:"passwordMinimumCharacterSetCount,omitempty"`
	PasswordMinimumLength                          *int64                                                  `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeScreenTimeout *int64                                                  `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasswordPreviousPasswordBlockCount             *int64                                                  `json:"passwordPreviousPasswordBlockCount,omitempty"`
	PasswordRequired                               *bool                                                   `json:"passwordRequired,omitempty"`
	PasswordRequiredType                           *WindowsPhone81GeneralConfigurationPasswordRequiredType `json:"passwordRequiredType,omitempty"`
	PasswordSignInFailureCountBeforeFactoryReset   *int64                                                  `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	RoleScopeTagIds                                *[]string                                               `json:"roleScopeTagIds,omitempty"`
	ScreenCaptureBlocked                           *bool                                                   `json:"screenCaptureBlocked,omitempty"`
	StorageBlockRemovableStorage                   *bool                                                   `json:"storageBlockRemovableStorage,omitempty"`
	StorageRequireEncryption                       *bool                                                   `json:"storageRequireEncryption,omitempty"`
	SupportsScopeTags                              *bool                                                   `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                             *DeviceConfigurationUserOverview                        `json:"userStatusOverview,omitempty"`
	UserStatuses                                   *[]DeviceConfigurationUserStatus                        `json:"userStatuses,omitempty"`
	Version                                        *int64                                                  `json:"version,omitempty"`
	WebBrowserBlocked                              *bool                                                   `json:"webBrowserBlocked,omitempty"`
	WifiBlockAutomaticConnectHotspots              *bool                                                   `json:"wifiBlockAutomaticConnectHotspots,omitempty"`
	WifiBlockHotspotReporting                      *bool                                                   `json:"wifiBlockHotspotReporting,omitempty"`
	WifiBlocked                                    *bool                                                   `json:"wifiBlocked,omitempty"`
	WindowsStoreBlocked                            *bool                                                   `json:"windowsStoreBlocked,omitempty"`
}
