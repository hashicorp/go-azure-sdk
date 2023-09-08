package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppProtection struct {
	AllowedInboundDataTransferSources       *WindowsManagedAppProtectionAllowedInboundDataTransferSources       `json:"allowedInboundDataTransferSources,omitempty"`
	AllowedOutboundClipboardSharingLevel    *WindowsManagedAppProtectionAllowedOutboundClipboardSharingLevel    `json:"allowedOutboundClipboardSharingLevel,omitempty"`
	AllowedOutboundDataTransferDestinations *WindowsManagedAppProtectionAllowedOutboundDataTransferDestinations `json:"allowedOutboundDataTransferDestinations,omitempty"`
	AppActionIfUnableToAuthenticateUser     *WindowsManagedAppProtectionAppActionIfUnableToAuthenticateUser     `json:"appActionIfUnableToAuthenticateUser,omitempty"`
	Apps                                    *[]ManagedMobileApp                                                 `json:"apps,omitempty"`
	Assignments                             *[]TargetedManagedAppPolicyAssignment                               `json:"assignments,omitempty"`
	CreatedDateTime                         *string                                                             `json:"createdDateTime,omitempty"`
	DeployedAppCount                        *int64                                                              `json:"deployedAppCount,omitempty"`
	Description                             *string                                                             `json:"description,omitempty"`
	DisplayName                             *string                                                             `json:"displayName,omitempty"`
	Id                                      *string                                                             `json:"id,omitempty"`
	IsAssigned                              *bool                                                               `json:"isAssigned,omitempty"`
	LastModifiedDateTime                    *string                                                             `json:"lastModifiedDateTime,omitempty"`
	MaximumAllowedDeviceThreatLevel         *WindowsManagedAppProtectionMaximumAllowedDeviceThreatLevel         `json:"maximumAllowedDeviceThreatLevel,omitempty"`
	MaximumRequiredOsVersion                *string                                                             `json:"maximumRequiredOsVersion,omitempty"`
	MaximumWarningOsVersion                 *string                                                             `json:"maximumWarningOsVersion,omitempty"`
	MaximumWipeOsVersion                    *string                                                             `json:"maximumWipeOsVersion,omitempty"`
	MinimumRequiredAppVersion               *string                                                             `json:"minimumRequiredAppVersion,omitempty"`
	MinimumRequiredOsVersion                *string                                                             `json:"minimumRequiredOsVersion,omitempty"`
	MinimumRequiredSdkVersion               *string                                                             `json:"minimumRequiredSdkVersion,omitempty"`
	MinimumWarningAppVersion                *string                                                             `json:"minimumWarningAppVersion,omitempty"`
	MinimumWarningOsVersion                 *string                                                             `json:"minimumWarningOsVersion,omitempty"`
	MinimumWipeAppVersion                   *string                                                             `json:"minimumWipeAppVersion,omitempty"`
	MinimumWipeOsVersion                    *string                                                             `json:"minimumWipeOsVersion,omitempty"`
	MinimumWipeSdkVersion                   *string                                                             `json:"minimumWipeSdkVersion,omitempty"`
	MobileThreatDefenseRemediationAction    *WindowsManagedAppProtectionMobileThreatDefenseRemediationAction    `json:"mobileThreatDefenseRemediationAction,omitempty"`
	ODataType                               *string                                                             `json:"@odata.type,omitempty"`
	PeriodOfflineBeforeAccessCheck          *string                                                             `json:"periodOfflineBeforeAccessCheck,omitempty"`
	PeriodOfflineBeforeWipeIsEnforced       *string                                                             `json:"periodOfflineBeforeWipeIsEnforced,omitempty"`
	PrintBlocked                            *bool                                                               `json:"printBlocked,omitempty"`
	RoleScopeTagIds                         *[]string                                                           `json:"roleScopeTagIds,omitempty"`
	Version                                 *string                                                             `json:"version,omitempty"`
}
