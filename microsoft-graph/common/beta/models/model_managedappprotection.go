package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtection struct {
	AllowedDataIngestionLocations                  *[]ManagedAppProtectionAllowedDataIngestionLocations         `json:"allowedDataIngestionLocations,omitempty"`
	AllowedDataStorageLocations                    *[]ManagedAppProtectionAllowedDataStorageLocations           `json:"allowedDataStorageLocations,omitempty"`
	AllowedInboundDataTransferSources              *ManagedAppProtectionAllowedInboundDataTransferSources       `json:"allowedInboundDataTransferSources,omitempty"`
	AllowedOutboundClipboardSharingExceptionLength *int64                                                       `json:"allowedOutboundClipboardSharingExceptionLength,omitempty"`
	AllowedOutboundClipboardSharingLevel           *ManagedAppProtectionAllowedOutboundClipboardSharingLevel    `json:"allowedOutboundClipboardSharingLevel,omitempty"`
	AllowedOutboundDataTransferDestinations        *ManagedAppProtectionAllowedOutboundDataTransferDestinations `json:"allowedOutboundDataTransferDestinations,omitempty"`
	AppActionIfDeviceComplianceRequired            *ManagedAppProtectionAppActionIfDeviceComplianceRequired     `json:"appActionIfDeviceComplianceRequired,omitempty"`
	AppActionIfMaximumPinRetriesExceeded           *ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded    `json:"appActionIfMaximumPinRetriesExceeded,omitempty"`
	AppActionIfUnableToAuthenticateUser            *ManagedAppProtectionAppActionIfUnableToAuthenticateUser     `json:"appActionIfUnableToAuthenticateUser,omitempty"`
	BlockDataIngestionIntoOrganizationDocuments    *bool                                                        `json:"blockDataIngestionIntoOrganizationDocuments,omitempty"`
	ContactSyncBlocked                             *bool                                                        `json:"contactSyncBlocked,omitempty"`
	CreatedDateTime                                *string                                                      `json:"createdDateTime,omitempty"`
	DataBackupBlocked                              *bool                                                        `json:"dataBackupBlocked,omitempty"`
	Description                                    *string                                                      `json:"description,omitempty"`
	DeviceComplianceRequired                       *bool                                                        `json:"deviceComplianceRequired,omitempty"`
	DialerRestrictionLevel                         *ManagedAppProtectionDialerRestrictionLevel                  `json:"dialerRestrictionLevel,omitempty"`
	DisableAppPinIfDevicePinIsSet                  *bool                                                        `json:"disableAppPinIfDevicePinIsSet,omitempty"`
	DisplayName                                    *string                                                      `json:"displayName,omitempty"`
	FingerprintBlocked                             *bool                                                        `json:"fingerprintBlocked,omitempty"`
	GracePeriodToBlockAppsDuringOffClockHours      *string                                                      `json:"gracePeriodToBlockAppsDuringOffClockHours,omitempty"`
	Id                                             *string                                                      `json:"id,omitempty"`
	LastModifiedDateTime                           *string                                                      `json:"lastModifiedDateTime,omitempty"`
	ManagedBrowser                                 *ManagedAppProtectionManagedBrowser                          `json:"managedBrowser,omitempty"`
	ManagedBrowserToOpenLinksRequired              *bool                                                        `json:"managedBrowserToOpenLinksRequired,omitempty"`
	MaximumAllowedDeviceThreatLevel                *ManagedAppProtectionMaximumAllowedDeviceThreatLevel         `json:"maximumAllowedDeviceThreatLevel,omitempty"`
	MaximumPinRetries                              *int64                                                       `json:"maximumPinRetries,omitempty"`
	MaximumRequiredOsVersion                       *string                                                      `json:"maximumRequiredOsVersion,omitempty"`
	MaximumWarningOsVersion                        *string                                                      `json:"maximumWarningOsVersion,omitempty"`
	MaximumWipeOsVersion                           *string                                                      `json:"maximumWipeOsVersion,omitempty"`
	MinimumPinLength                               *int64                                                       `json:"minimumPinLength,omitempty"`
	MinimumRequiredAppVersion                      *string                                                      `json:"minimumRequiredAppVersion,omitempty"`
	MinimumRequiredOsVersion                       *string                                                      `json:"minimumRequiredOsVersion,omitempty"`
	MinimumWarningAppVersion                       *string                                                      `json:"minimumWarningAppVersion,omitempty"`
	MinimumWarningOsVersion                        *string                                                      `json:"minimumWarningOsVersion,omitempty"`
	MinimumWipeAppVersion                          *string                                                      `json:"minimumWipeAppVersion,omitempty"`
	MinimumWipeOsVersion                           *string                                                      `json:"minimumWipeOsVersion,omitempty"`
	MobileThreatDefensePartnerPriority             *ManagedAppProtectionMobileThreatDefensePartnerPriority      `json:"mobileThreatDefensePartnerPriority,omitempty"`
	MobileThreatDefenseRemediationAction           *ManagedAppProtectionMobileThreatDefenseRemediationAction    `json:"mobileThreatDefenseRemediationAction,omitempty"`
	NotificationRestriction                        *ManagedAppProtectionNotificationRestriction                 `json:"notificationRestriction,omitempty"`
	ODataType                                      *string                                                      `json:"@odata.type,omitempty"`
	OrganizationalCredentialsRequired              *bool                                                        `json:"organizationalCredentialsRequired,omitempty"`
	PeriodBeforePinReset                           *string                                                      `json:"periodBeforePinReset,omitempty"`
	PeriodOfflineBeforeAccessCheck                 *string                                                      `json:"periodOfflineBeforeAccessCheck,omitempty"`
	PeriodOfflineBeforeWipeIsEnforced              *string                                                      `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`
	PeriodOnlineBeforeAccessCheck                  *string                                                      `json:"periodOnlineBeforeAccessCheck,omitempty"`
	PinCharacterSet                                *ManagedAppProtectionPinCharacterSet                         `json:"pinCharacterSet,omitempty"`
	PinRequired                                    *bool                                                        `json:"pinRequired,omitempty"`
	PinRequiredInsteadOfBiometricTimeout           *string                                                      `json:"pinRequiredInsteadOfBiometricTimeout,omitempty"`
	PreviousPinBlockCount                          *int64                                                       `json:"previousPinBlockCount,omitempty"`
	PrintBlocked                                   *bool                                                        `json:"printBlocked,omitempty"`
	RoleScopeTagIds                                *[]string                                                    `json:"roleScopeTagIds,omitempty"`
	SaveAsBlocked                                  *bool                                                        `json:"saveAsBlocked,omitempty"`
	SimplePinBlocked                               *bool                                                        `json:"simplePinBlocked,omitempty"`
	Version                                        *string                                                      `json:"version,omitempty"`
}
