package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtection struct {
	AllowedDataStorageLocations                     *[]DefaultManagedAppProtectionAllowedDataStorageLocations           `json:"allowedDataStorageLocations,omitempty"`
	AllowedInboundDataTransferSources               *DefaultManagedAppProtectionAllowedInboundDataTransferSources       `json:"allowedInboundDataTransferSources,omitempty"`
	AllowedOutboundClipboardSharingLevel            *DefaultManagedAppProtectionAllowedOutboundClipboardSharingLevel    `json:"allowedOutboundClipboardSharingLevel,omitempty"`
	AllowedOutboundDataTransferDestinations         *DefaultManagedAppProtectionAllowedOutboundDataTransferDestinations `json:"allowedOutboundDataTransferDestinations,omitempty"`
	AppDataEncryptionType                           *DefaultManagedAppProtectionAppDataEncryptionType                   `json:"appDataEncryptionType,omitempty"`
	Apps                                            *[]ManagedMobileApp                                                 `json:"apps,omitempty"`
	ContactSyncBlocked                              *bool                                                               `json:"contactSyncBlocked,omitempty"`
	CreatedDateTime                                 *string                                                             `json:"createdDateTime,omitempty"`
	CustomSettings                                  *[]KeyValuePair                                                     `json:"customSettings,omitempty"`
	DataBackupBlocked                               *bool                                                               `json:"dataBackupBlocked,omitempty"`
	DeployedAppCount                                *int64                                                              `json:"deployedAppCount,omitempty"`
	DeploymentSummary                               *ManagedAppPolicyDeploymentSummary                                  `json:"deploymentSummary,omitempty"`
	Description                                     *string                                                             `json:"description,omitempty"`
	DeviceComplianceRequired                        *bool                                                               `json:"deviceComplianceRequired,omitempty"`
	DisableAppEncryptionIfDeviceEncryptionIsEnabled *bool                                                               `json:"disableAppEncryptionIfDeviceEncryptionIsEnabled,omitempty"`
	DisableAppPinIfDevicePinIsSet                   *bool                                                               `json:"disableAppPinIfDevicePinIsSet,omitempty"`
	DisplayName                                     *string                                                             `json:"displayName,omitempty"`
	EncryptAppData                                  *bool                                                               `json:"encryptAppData,omitempty"`
	FaceIdBlocked                                   *bool                                                               `json:"faceIdBlocked,omitempty"`
	FingerprintBlocked                              *bool                                                               `json:"fingerprintBlocked,omitempty"`
	Id                                              *string                                                             `json:"id,omitempty"`
	LastModifiedDateTime                            *string                                                             `json:"lastModifiedDateTime,omitempty"`
	ManagedBrowser                                  *DefaultManagedAppProtectionManagedBrowser                          `json:"managedBrowser,omitempty"`
	ManagedBrowserToOpenLinksRequired               *bool                                                               `json:"managedBrowserToOpenLinksRequired,omitempty"`
	MaximumPinRetries                               *int64                                                              `json:"maximumPinRetries,omitempty"`
	MinimumPinLength                                *int64                                                              `json:"minimumPinLength,omitempty"`
	MinimumRequiredAppVersion                       *string                                                             `json:"minimumRequiredAppVersion,omitempty"`
	MinimumRequiredOsVersion                        *string                                                             `json:"minimumRequiredOsVersion,omitempty"`
	MinimumRequiredPatchVersion                     *string                                                             `json:"minimumRequiredPatchVersion,omitempty"`
	MinimumRequiredSdkVersion                       *string                                                             `json:"minimumRequiredSdkVersion,omitempty"`
	MinimumWarningAppVersion                        *string                                                             `json:"minimumWarningAppVersion,omitempty"`
	MinimumWarningOsVersion                         *string                                                             `json:"minimumWarningOsVersion,omitempty"`
	MinimumWarningPatchVersion                      *string                                                             `json:"minimumWarningPatchVersion,omitempty"`
	ODataType                                       *string                                                             `json:"@odata.type,omitempty"`
	OrganizationalCredentialsRequired               *bool                                                               `json:"organizationalCredentialsRequired,omitempty"`
	PeriodBeforePinReset                            *string                                                             `json:"periodBeforePinReset,omitempty"`
	PeriodOfflineBeforeAccessCheck                  *string                                                             `json:"periodOfflineBeforeAccessCheck,omitempty"`
	PeriodOfflineBeforeWipeIsEnforced               *string                                                             `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`
	PeriodOnlineBeforeAccessCheck                   *string                                                             `json:"periodOnlineBeforeAccessCheck,omitempty"`
	PinCharacterSet                                 *DefaultManagedAppProtectionPinCharacterSet                         `json:"pinCharacterSet,omitempty"`
	PinRequired                                     *bool                                                               `json:"pinRequired,omitempty"`
	PrintBlocked                                    *bool                                                               `json:"printBlocked,omitempty"`
	SaveAsBlocked                                   *bool                                                               `json:"saveAsBlocked,omitempty"`
	ScreenCaptureBlocked                            *bool                                                               `json:"screenCaptureBlocked,omitempty"`
	SimplePinBlocked                                *bool                                                               `json:"simplePinBlocked,omitempty"`
	Version                                         *string                                                             `json:"version,omitempty"`
}
