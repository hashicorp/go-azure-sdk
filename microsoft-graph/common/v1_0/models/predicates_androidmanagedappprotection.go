package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionOperationPredicate struct {
	ContactSyncBlocked                              *bool
	CreatedDateTime                                 *string
	CustomBrowserDisplayName                        *string
	CustomBrowserPackageId                          *string
	DataBackupBlocked                               *bool
	DeployedAppCount                                *int64
	Description                                     *string
	DeviceComplianceRequired                        *bool
	DisableAppEncryptionIfDeviceEncryptionIsEnabled *bool
	DisableAppPinIfDevicePinIsSet                   *bool
	DisplayName                                     *string
	EncryptAppData                                  *bool
	FingerprintBlocked                              *bool
	Id                                              *string
	IsAssigned                                      *bool
	LastModifiedDateTime                            *string
	ManagedBrowserToOpenLinksRequired               *bool
	MaximumPinRetries                               *int64
	MinimumPinLength                                *int64
	MinimumRequiredAppVersion                       *string
	MinimumRequiredOsVersion                        *string
	MinimumRequiredPatchVersion                     *string
	MinimumWarningAppVersion                        *string
	MinimumWarningOsVersion                         *string
	MinimumWarningPatchVersion                      *string
	ODataType                                       *string
	OrganizationalCredentialsRequired               *bool
	PeriodBeforePinReset                            *string
	PeriodOfflineBeforeAccessCheck                  *string
	PeriodOfflineBeforeWipeIsEnforced               *string
	PeriodOnlineBeforeAccessCheck                   *string
	PinRequired                                     *bool
	PrintBlocked                                    *bool
	SaveAsBlocked                                   *bool
	ScreenCaptureBlocked                            *bool
	SimplePinBlocked                                *bool
	Version                                         *string
}

func (p AndroidManagedAppProtectionOperationPredicate) Matches(input AndroidManagedAppProtection) bool {

	if p.ContactSyncBlocked != nil && (input.ContactSyncBlocked == nil || *p.ContactSyncBlocked != *input.ContactSyncBlocked) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.CustomBrowserDisplayName != nil && (input.CustomBrowserDisplayName == nil || *p.CustomBrowserDisplayName != *input.CustomBrowserDisplayName) {
		return false
	}

	if p.CustomBrowserPackageId != nil && (input.CustomBrowserPackageId == nil || *p.CustomBrowserPackageId != *input.CustomBrowserPackageId) {
		return false
	}

	if p.DataBackupBlocked != nil && (input.DataBackupBlocked == nil || *p.DataBackupBlocked != *input.DataBackupBlocked) {
		return false
	}

	if p.DeployedAppCount != nil && (input.DeployedAppCount == nil || *p.DeployedAppCount != *input.DeployedAppCount) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DeviceComplianceRequired != nil && (input.DeviceComplianceRequired == nil || *p.DeviceComplianceRequired != *input.DeviceComplianceRequired) {
		return false
	}

	if p.DisableAppEncryptionIfDeviceEncryptionIsEnabled != nil && (input.DisableAppEncryptionIfDeviceEncryptionIsEnabled == nil || *p.DisableAppEncryptionIfDeviceEncryptionIsEnabled != *input.DisableAppEncryptionIfDeviceEncryptionIsEnabled) {
		return false
	}

	if p.DisableAppPinIfDevicePinIsSet != nil && (input.DisableAppPinIfDevicePinIsSet == nil || *p.DisableAppPinIfDevicePinIsSet != *input.DisableAppPinIfDevicePinIsSet) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EncryptAppData != nil && (input.EncryptAppData == nil || *p.EncryptAppData != *input.EncryptAppData) {
		return false
	}

	if p.FingerprintBlocked != nil && (input.FingerprintBlocked == nil || *p.FingerprintBlocked != *input.FingerprintBlocked) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAssigned != nil && (input.IsAssigned == nil || *p.IsAssigned != *input.IsAssigned) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ManagedBrowserToOpenLinksRequired != nil && (input.ManagedBrowserToOpenLinksRequired == nil || *p.ManagedBrowserToOpenLinksRequired != *input.ManagedBrowserToOpenLinksRequired) {
		return false
	}

	if p.MaximumPinRetries != nil && (input.MaximumPinRetries == nil || *p.MaximumPinRetries != *input.MaximumPinRetries) {
		return false
	}

	if p.MinimumPinLength != nil && (input.MinimumPinLength == nil || *p.MinimumPinLength != *input.MinimumPinLength) {
		return false
	}

	if p.MinimumRequiredAppVersion != nil && (input.MinimumRequiredAppVersion == nil || *p.MinimumRequiredAppVersion != *input.MinimumRequiredAppVersion) {
		return false
	}

	if p.MinimumRequiredOsVersion != nil && (input.MinimumRequiredOsVersion == nil || *p.MinimumRequiredOsVersion != *input.MinimumRequiredOsVersion) {
		return false
	}

	if p.MinimumRequiredPatchVersion != nil && (input.MinimumRequiredPatchVersion == nil || *p.MinimumRequiredPatchVersion != *input.MinimumRequiredPatchVersion) {
		return false
	}

	if p.MinimumWarningAppVersion != nil && (input.MinimumWarningAppVersion == nil || *p.MinimumWarningAppVersion != *input.MinimumWarningAppVersion) {
		return false
	}

	if p.MinimumWarningOsVersion != nil && (input.MinimumWarningOsVersion == nil || *p.MinimumWarningOsVersion != *input.MinimumWarningOsVersion) {
		return false
	}

	if p.MinimumWarningPatchVersion != nil && (input.MinimumWarningPatchVersion == nil || *p.MinimumWarningPatchVersion != *input.MinimumWarningPatchVersion) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OrganizationalCredentialsRequired != nil && (input.OrganizationalCredentialsRequired == nil || *p.OrganizationalCredentialsRequired != *input.OrganizationalCredentialsRequired) {
		return false
	}

	if p.PeriodBeforePinReset != nil && (input.PeriodBeforePinReset == nil || *p.PeriodBeforePinReset != *input.PeriodBeforePinReset) {
		return false
	}

	if p.PeriodOfflineBeforeAccessCheck != nil && (input.PeriodOfflineBeforeAccessCheck == nil || *p.PeriodOfflineBeforeAccessCheck != *input.PeriodOfflineBeforeAccessCheck) {
		return false
	}

	if p.PeriodOfflineBeforeWipeIsEnforced != nil && (input.PeriodOfflineBeforeWipeIsEnforced == nil || *p.PeriodOfflineBeforeWipeIsEnforced != *input.PeriodOfflineBeforeWipeIsEnforced) {
		return false
	}

	if p.PeriodOnlineBeforeAccessCheck != nil && (input.PeriodOnlineBeforeAccessCheck == nil || *p.PeriodOnlineBeforeAccessCheck != *input.PeriodOnlineBeforeAccessCheck) {
		return false
	}

	if p.PinRequired != nil && (input.PinRequired == nil || *p.PinRequired != *input.PinRequired) {
		return false
	}

	if p.PrintBlocked != nil && (input.PrintBlocked == nil || *p.PrintBlocked != *input.PrintBlocked) {
		return false
	}

	if p.SaveAsBlocked != nil && (input.SaveAsBlocked == nil || *p.SaveAsBlocked != *input.SaveAsBlocked) {
		return false
	}

	if p.ScreenCaptureBlocked != nil && (input.ScreenCaptureBlocked == nil || *p.ScreenCaptureBlocked != *input.ScreenCaptureBlocked) {
		return false
	}

	if p.SimplePinBlocked != nil && (input.SimplePinBlocked == nil || *p.SimplePinBlocked != *input.SimplePinBlocked) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
