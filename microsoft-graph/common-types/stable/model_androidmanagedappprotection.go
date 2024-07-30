package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtection struct {
	AllowedDataStorageLocations                     *AndroidManagedAppProtectionAllowedDataStorageLocations             `json:"allowedDataStorageLocations,omitempty"`
	AllowedInboundDataTransferSources               *AndroidManagedAppProtectionAllowedInboundDataTransferSources       `json:"allowedInboundDataTransferSources,omitempty"`
	AllowedOutboundClipboardSharingLevel            *AndroidManagedAppProtectionAllowedOutboundClipboardSharingLevel    `json:"allowedOutboundClipboardSharingLevel,omitempty"`
	AllowedOutboundDataTransferDestinations         *AndroidManagedAppProtectionAllowedOutboundDataTransferDestinations `json:"allowedOutboundDataTransferDestinations,omitempty"`
	Apps                                            *[]ManagedMobileApp                                                 `json:"apps,omitempty"`
	Assignments                                     *[]TargetedManagedAppPolicyAssignment                               `json:"assignments,omitempty"`
	ContactSyncBlocked                              *bool                                                               `json:"contactSyncBlocked,omitempty"`
	CreatedDateTime                                 *string                                                             `json:"createdDateTime,omitempty"`
	CustomBrowserDisplayName                        *string                                                             `json:"customBrowserDisplayName,omitempty"`
	CustomBrowserPackageId                          *string                                                             `json:"customBrowserPackageId,omitempty"`
	DataBackupBlocked                               *bool                                                               `json:"dataBackupBlocked,omitempty"`
	DeployedAppCount                                *int64                                                              `json:"deployedAppCount,omitempty"`
	DeploymentSummary                               *ManagedAppPolicyDeploymentSummary                                  `json:"deploymentSummary,omitempty"`
	Description                                     *string                                                             `json:"description,omitempty"`
	DeviceComplianceRequired                        *bool                                                               `json:"deviceComplianceRequired,omitempty"`
	DisableAppEncryptionIfDeviceEncryptionIsEnabled *bool                                                               `json:"disableAppEncryptionIfDeviceEncryptionIsEnabled,omitempty"`
	DisableAppPinIfDevicePinIsSet                   *bool                                                               `json:"disableAppPinIfDevicePinIsSet,omitempty"`
	DisplayName                                     *string                                                             `json:"displayName,omitempty"`
	EncryptAppData                                  *bool                                                               `json:"encryptAppData,omitempty"`
	FingerprintBlocked                              *bool                                                               `json:"fingerprintBlocked,omitempty"`
	Id                                              *string                                                             `json:"id,omitempty"`
	IsAssigned                                      *bool                                                               `json:"isAssigned,omitempty"`
	LastModifiedDateTime                            *string                                                             `json:"lastModifiedDateTime,omitempty"`
	ManagedBrowser                                  *AndroidManagedAppProtectionManagedBrowser                          `json:"managedBrowser,omitempty"`
	ManagedBrowserToOpenLinksRequired               *bool                                                               `json:"managedBrowserToOpenLinksRequired,omitempty"`
	MaximumPinRetries                               *int64                                                              `json:"maximumPinRetries,omitempty"`
	MinimumPinLength                                *int64                                                              `json:"minimumPinLength,omitempty"`
	MinimumRequiredAppVersion                       *string                                                             `json:"minimumRequiredAppVersion,omitempty"`
	MinimumRequiredOsVersion                        *string                                                             `json:"minimumRequiredOsVersion,omitempty"`
	MinimumRequiredPatchVersion                     *string                                                             `json:"minimumRequiredPatchVersion,omitempty"`
	MinimumWarningAppVersion                        *string                                                             `json:"minimumWarningAppVersion,omitempty"`
	MinimumWarningOsVersion                         *string                                                             `json:"minimumWarningOsVersion,omitempty"`
	MinimumWarningPatchVersion                      *string                                                             `json:"minimumWarningPatchVersion,omitempty"`
	ODataType                                       *string                                                             `json:"@odata.type,omitempty"`
	OrganizationalCredentialsRequired               *bool                                                               `json:"organizationalCredentialsRequired,omitempty"`
	PeriodBeforePinReset                            *string                                                             `json:"periodBeforePinReset,omitempty"`
	PeriodOfflineBeforeAccessCheck                  *string                                                             `json:"periodOfflineBeforeAccessCheck,omitempty"`
	PeriodOfflineBeforeWipeIsEnforced               *string                                                             `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`
	PeriodOnlineBeforeAccessCheck                   *string                                                             `json:"periodOnlineBeforeAccessCheck,omitempty"`
	PinCharacterSet                                 *AndroidManagedAppProtectionPinCharacterSet                         `json:"pinCharacterSet,omitempty"`
	PinRequired                                     *bool                                                               `json:"pinRequired,omitempty"`
	PrintBlocked                                    *bool                                                               `json:"printBlocked,omitempty"`
	SaveAsBlocked                                   *bool                                                               `json:"saveAsBlocked,omitempty"`
	ScreenCaptureBlocked                            *bool                                                               `json:"screenCaptureBlocked,omitempty"`
	SimplePinBlocked                                *bool                                                               `json:"simplePinBlocked,omitempty"`
	Version                                         *string                                                             `json:"version,omitempty"`
}
