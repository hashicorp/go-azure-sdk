package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionOperationPredicate struct {
	ContactSyncBlocked                *bool
	CreatedDateTime                   *string
	DataBackupBlocked                 *bool
	Description                       *string
	DeviceComplianceRequired          *bool
	DisableAppPinIfDevicePinIsSet     *bool
	DisplayName                       *string
	FingerprintBlocked                *bool
	Id                                *string
	LastModifiedDateTime              *string
	ManagedBrowserToOpenLinksRequired *bool
	MaximumPinRetries                 *int64
	MinimumPinLength                  *int64
	MinimumRequiredAppVersion         *string
	MinimumRequiredOsVersion          *string
	MinimumWarningAppVersion          *string
	MinimumWarningOsVersion           *string
	ODataType                         *string
	OrganizationalCredentialsRequired *bool
	PeriodBeforePinReset              *string
	PeriodOfflineBeforeAccessCheck    *string
	PeriodOfflineBeforeWipeIsEnforced *string
	PeriodOnlineBeforeAccessCheck     *string
	PinRequired                       *bool
	PrintBlocked                      *bool
	SaveAsBlocked                     *bool
	SimplePinBlocked                  *bool
	Version                           *string
}

func (p ManagedAppProtectionOperationPredicate) Matches(input ManagedAppProtection) bool {

	if p.ContactSyncBlocked != nil && (input.ContactSyncBlocked == nil || *p.ContactSyncBlocked != *input.ContactSyncBlocked) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DataBackupBlocked != nil && (input.DataBackupBlocked == nil || *p.DataBackupBlocked != *input.DataBackupBlocked) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DeviceComplianceRequired != nil && (input.DeviceComplianceRequired == nil || *p.DeviceComplianceRequired != *input.DeviceComplianceRequired) {
		return false
	}

	if p.DisableAppPinIfDevicePinIsSet != nil && (input.DisableAppPinIfDevicePinIsSet == nil || *p.DisableAppPinIfDevicePinIsSet != *input.DisableAppPinIfDevicePinIsSet) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.FingerprintBlocked != nil && (input.FingerprintBlocked == nil || *p.FingerprintBlocked != *input.FingerprintBlocked) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
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

	if p.MinimumWarningAppVersion != nil && (input.MinimumWarningAppVersion == nil || *p.MinimumWarningAppVersion != *input.MinimumWarningAppVersion) {
		return false
	}

	if p.MinimumWarningOsVersion != nil && (input.MinimumWarningOsVersion == nil || *p.MinimumWarningOsVersion != *input.MinimumWarningOsVersion) {
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

	if p.SimplePinBlocked != nil && (input.SimplePinBlocked == nil || *p.SimplePinBlocked != *input.SimplePinBlocked) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
