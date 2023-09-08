package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidCompliancePolicyOperationPredicate struct {
	ConditionStatementId                               *string
	CreatedDateTime                                    *string
	Description                                        *string
	DeviceThreatProtectionEnabled                      *bool
	DisplayName                                        *string
	Id                                                 *string
	LastModifiedDateTime                               *string
	MinAndroidSecurityPatchLevel                       *string
	ODataType                                          *string
	OsMaximumVersion                                   *string
	OsMinimumVersion                                   *string
	PasswordExpirationDays                             *int64
	PasswordMinimumLength                              *int64
	PasswordMinutesOfInactivityBeforeLock              *int64
	PasswordPreviousPasswordBlockCount                 *int64
	PasswordRequired                                   *bool
	PasswordSignInFailureCountBeforeFactoryReset       *int64
	SecurityBlockDeviceAdministratorManagedDevices     *bool
	SecurityBlockJailbrokenDevices                     *bool
	SecurityDisableUsbDebugging                        *bool
	SecurityPreventInstallAppsFromUnknownSources       *bool
	SecurityRequireCompanyPortalAppIntegrity           *bool
	SecurityRequireGooglePlayServices                  *bool
	SecurityRequireSafetyNetAttestationBasicIntegrity  *bool
	SecurityRequireSafetyNetAttestationCertifiedDevice *bool
	SecurityRequireUpToDateSecurityProviders           *bool
	SecurityRequireVerifyApps                          *bool
	StorageRequireEncryption                           *bool
	Version                                            *int64
}

func (p AndroidCompliancePolicyOperationPredicate) Matches(input AndroidCompliancePolicy) bool {

	if p.ConditionStatementId != nil && (input.ConditionStatementId == nil || *p.ConditionStatementId != *input.ConditionStatementId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DeviceThreatProtectionEnabled != nil && (input.DeviceThreatProtectionEnabled == nil || *p.DeviceThreatProtectionEnabled != *input.DeviceThreatProtectionEnabled) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.MinAndroidSecurityPatchLevel != nil && (input.MinAndroidSecurityPatchLevel == nil || *p.MinAndroidSecurityPatchLevel != *input.MinAndroidSecurityPatchLevel) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OsMaximumVersion != nil && (input.OsMaximumVersion == nil || *p.OsMaximumVersion != *input.OsMaximumVersion) {
		return false
	}

	if p.OsMinimumVersion != nil && (input.OsMinimumVersion == nil || *p.OsMinimumVersion != *input.OsMinimumVersion) {
		return false
	}

	if p.PasswordExpirationDays != nil && (input.PasswordExpirationDays == nil || *p.PasswordExpirationDays != *input.PasswordExpirationDays) {
		return false
	}

	if p.PasswordMinimumLength != nil && (input.PasswordMinimumLength == nil || *p.PasswordMinimumLength != *input.PasswordMinimumLength) {
		return false
	}

	if p.PasswordMinutesOfInactivityBeforeLock != nil && (input.PasswordMinutesOfInactivityBeforeLock == nil || *p.PasswordMinutesOfInactivityBeforeLock != *input.PasswordMinutesOfInactivityBeforeLock) {
		return false
	}

	if p.PasswordPreviousPasswordBlockCount != nil && (input.PasswordPreviousPasswordBlockCount == nil || *p.PasswordPreviousPasswordBlockCount != *input.PasswordPreviousPasswordBlockCount) {
		return false
	}

	if p.PasswordRequired != nil && (input.PasswordRequired == nil || *p.PasswordRequired != *input.PasswordRequired) {
		return false
	}

	if p.PasswordSignInFailureCountBeforeFactoryReset != nil && (input.PasswordSignInFailureCountBeforeFactoryReset == nil || *p.PasswordSignInFailureCountBeforeFactoryReset != *input.PasswordSignInFailureCountBeforeFactoryReset) {
		return false
	}

	if p.SecurityBlockDeviceAdministratorManagedDevices != nil && (input.SecurityBlockDeviceAdministratorManagedDevices == nil || *p.SecurityBlockDeviceAdministratorManagedDevices != *input.SecurityBlockDeviceAdministratorManagedDevices) {
		return false
	}

	if p.SecurityBlockJailbrokenDevices != nil && (input.SecurityBlockJailbrokenDevices == nil || *p.SecurityBlockJailbrokenDevices != *input.SecurityBlockJailbrokenDevices) {
		return false
	}

	if p.SecurityDisableUsbDebugging != nil && (input.SecurityDisableUsbDebugging == nil || *p.SecurityDisableUsbDebugging != *input.SecurityDisableUsbDebugging) {
		return false
	}

	if p.SecurityPreventInstallAppsFromUnknownSources != nil && (input.SecurityPreventInstallAppsFromUnknownSources == nil || *p.SecurityPreventInstallAppsFromUnknownSources != *input.SecurityPreventInstallAppsFromUnknownSources) {
		return false
	}

	if p.SecurityRequireCompanyPortalAppIntegrity != nil && (input.SecurityRequireCompanyPortalAppIntegrity == nil || *p.SecurityRequireCompanyPortalAppIntegrity != *input.SecurityRequireCompanyPortalAppIntegrity) {
		return false
	}

	if p.SecurityRequireGooglePlayServices != nil && (input.SecurityRequireGooglePlayServices == nil || *p.SecurityRequireGooglePlayServices != *input.SecurityRequireGooglePlayServices) {
		return false
	}

	if p.SecurityRequireSafetyNetAttestationBasicIntegrity != nil && (input.SecurityRequireSafetyNetAttestationBasicIntegrity == nil || *p.SecurityRequireSafetyNetAttestationBasicIntegrity != *input.SecurityRequireSafetyNetAttestationBasicIntegrity) {
		return false
	}

	if p.SecurityRequireSafetyNetAttestationCertifiedDevice != nil && (input.SecurityRequireSafetyNetAttestationCertifiedDevice == nil || *p.SecurityRequireSafetyNetAttestationCertifiedDevice != *input.SecurityRequireSafetyNetAttestationCertifiedDevice) {
		return false
	}

	if p.SecurityRequireUpToDateSecurityProviders != nil && (input.SecurityRequireUpToDateSecurityProviders == nil || *p.SecurityRequireUpToDateSecurityProviders != *input.SecurityRequireUpToDateSecurityProviders) {
		return false
	}

	if p.SecurityRequireVerifyApps != nil && (input.SecurityRequireVerifyApps == nil || *p.SecurityRequireVerifyApps != *input.SecurityRequireVerifyApps) {
		return false
	}

	if p.StorageRequireEncryption != nil && (input.StorageRequireEncryption == nil || *p.StorageRequireEncryption != *input.StorageRequireEncryption) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
