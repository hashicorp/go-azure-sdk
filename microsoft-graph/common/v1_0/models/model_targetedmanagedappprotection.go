package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtection struct {
	AllowedDataStorageLocations             *[]TargetedManagedAppProtectionAllowedDataStorageLocations           `json:"allowedDataStorageLocations,omitempty"`
	AllowedInboundDataTransferSources       *TargetedManagedAppProtectionAllowedInboundDataTransferSources       `json:"allowedInboundDataTransferSources,omitempty"`
	AllowedOutboundClipboardSharingLevel    *TargetedManagedAppProtectionAllowedOutboundClipboardSharingLevel    `json:"allowedOutboundClipboardSharingLevel,omitempty"`
	AllowedOutboundDataTransferDestinations *TargetedManagedAppProtectionAllowedOutboundDataTransferDestinations `json:"allowedOutboundDataTransferDestinations,omitempty"`
	Assignments                             *[]TargetedManagedAppPolicyAssignment                                `json:"assignments,omitempty"`
	ContactSyncBlocked                      *bool                                                                `json:"contactSyncBlocked,omitempty"`
	CreatedDateTime                         *string                                                              `json:"createdDateTime,omitempty"`
	DataBackupBlocked                       *bool                                                                `json:"dataBackupBlocked,omitempty"`
	Description                             *string                                                              `json:"description,omitempty"`
	DeviceComplianceRequired                *bool                                                                `json:"deviceComplianceRequired,omitempty"`
	DisableAppPinIfDevicePinIsSet           *bool                                                                `json:"disableAppPinIfDevicePinIsSet,omitempty"`
	DisplayName                             *string                                                              `json:"displayName,omitempty"`
	FingerprintBlocked                      *bool                                                                `json:"fingerprintBlocked,omitempty"`
	Id                                      *string                                                              `json:"id,omitempty"`
	IsAssigned                              *bool                                                                `json:"isAssigned,omitempty"`
	LastModifiedDateTime                    *string                                                              `json:"lastModifiedDateTime,omitempty"`
	ManagedBrowser                          *TargetedManagedAppProtectionManagedBrowser                          `json:"managedBrowser,omitempty"`
	ManagedBrowserToOpenLinksRequired       *bool                                                                `json:"managedBrowserToOpenLinksRequired,omitempty"`
	MaximumPinRetries                       *int64                                                               `json:"maximumPinRetries,omitempty"`
	MinimumPinLength                        *int64                                                               `json:"minimumPinLength,omitempty"`
	MinimumRequiredAppVersion               *string                                                              `json:"minimumRequiredAppVersion,omitempty"`
	MinimumRequiredOsVersion                *string                                                              `json:"minimumRequiredOsVersion,omitempty"`
	MinimumWarningAppVersion                *string                                                              `json:"minimumWarningAppVersion,omitempty"`
	MinimumWarningOsVersion                 *string                                                              `json:"minimumWarningOsVersion,omitempty"`
	ODataType                               *string                                                              `json:"@odata.type,omitempty"`
	OrganizationalCredentialsRequired       *bool                                                                `json:"organizationalCredentialsRequired,omitempty"`
	PeriodBeforePinReset                    *string                                                              `json:"periodBeforePinReset,omitempty"`
	PeriodOfflineBeforeAccessCheck          *string                                                              `json:"periodOfflineBeforeAccessCheck,omitempty"`
	PeriodOfflineBeforeWipeIsEnforced       *string                                                              `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`
	PeriodOnlineBeforeAccessCheck           *string                                                              `json:"periodOnlineBeforeAccessCheck,omitempty"`
	PinCharacterSet                         *TargetedManagedAppProtectionPinCharacterSet                         `json:"pinCharacterSet,omitempty"`
	PinRequired                             *bool                                                                `json:"pinRequired,omitempty"`
	PrintBlocked                            *bool                                                                `json:"printBlocked,omitempty"`
	SaveAsBlocked                           *bool                                                                `json:"saveAsBlocked,omitempty"`
	SimplePinBlocked                        *bool                                                                `json:"simplePinBlocked,omitempty"`
	Version                                 *string                                                              `json:"version,omitempty"`
}
