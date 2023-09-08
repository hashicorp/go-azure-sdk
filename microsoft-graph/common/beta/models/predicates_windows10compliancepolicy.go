package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10CompliancePolicyOperationPredicate struct {
	ActiveFirewallRequired                 *bool
	AntiSpywareRequired                    *bool
	AntivirusRequired                      *bool
	BitLockerEnabled                       *bool
	CodeIntegrityEnabled                   *bool
	ConfigurationManagerComplianceRequired *bool
	CreatedDateTime                        *string
	DefenderEnabled                        *bool
	DefenderVersion                        *string
	Description                            *string
	DeviceThreatProtectionEnabled          *bool
	DisplayName                            *string
	EarlyLaunchAntiMalwareDriverEnabled    *bool
	Id                                     *string
	LastModifiedDateTime                   *string
	MobileOsMaximumVersion                 *string
	MobileOsMinimumVersion                 *string
	ODataType                              *string
	OsMaximumVersion                       *string
	OsMinimumVersion                       *string
	PasswordBlockSimple                    *bool
	PasswordExpirationDays                 *int64
	PasswordMinimumCharacterSetCount       *int64
	PasswordMinimumLength                  *int64
	PasswordMinutesOfInactivityBeforeLock  *int64
	PasswordPreviousPasswordBlockCount     *int64
	PasswordRequired                       *bool
	PasswordRequiredToUnlockFromIdle       *bool
	RequireHealthyDeviceReport             *bool
	RtpEnabled                             *bool
	SecureBootEnabled                      *bool
	SignatureOutOfDate                     *bool
	StorageRequireEncryption               *bool
	TpmRequired                            *bool
	Version                                *int64
}

func (p Windows10CompliancePolicyOperationPredicate) Matches(input Windows10CompliancePolicy) bool {

	if p.ActiveFirewallRequired != nil && (input.ActiveFirewallRequired == nil || *p.ActiveFirewallRequired != *input.ActiveFirewallRequired) {
		return false
	}

	if p.AntiSpywareRequired != nil && (input.AntiSpywareRequired == nil || *p.AntiSpywareRequired != *input.AntiSpywareRequired) {
		return false
	}

	if p.AntivirusRequired != nil && (input.AntivirusRequired == nil || *p.AntivirusRequired != *input.AntivirusRequired) {
		return false
	}

	if p.BitLockerEnabled != nil && (input.BitLockerEnabled == nil || *p.BitLockerEnabled != *input.BitLockerEnabled) {
		return false
	}

	if p.CodeIntegrityEnabled != nil && (input.CodeIntegrityEnabled == nil || *p.CodeIntegrityEnabled != *input.CodeIntegrityEnabled) {
		return false
	}

	if p.ConfigurationManagerComplianceRequired != nil && (input.ConfigurationManagerComplianceRequired == nil || *p.ConfigurationManagerComplianceRequired != *input.ConfigurationManagerComplianceRequired) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DefenderEnabled != nil && (input.DefenderEnabled == nil || *p.DefenderEnabled != *input.DefenderEnabled) {
		return false
	}

	if p.DefenderVersion != nil && (input.DefenderVersion == nil || *p.DefenderVersion != *input.DefenderVersion) {
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

	if p.EarlyLaunchAntiMalwareDriverEnabled != nil && (input.EarlyLaunchAntiMalwareDriverEnabled == nil || *p.EarlyLaunchAntiMalwareDriverEnabled != *input.EarlyLaunchAntiMalwareDriverEnabled) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.MobileOsMaximumVersion != nil && (input.MobileOsMaximumVersion == nil || *p.MobileOsMaximumVersion != *input.MobileOsMaximumVersion) {
		return false
	}

	if p.MobileOsMinimumVersion != nil && (input.MobileOsMinimumVersion == nil || *p.MobileOsMinimumVersion != *input.MobileOsMinimumVersion) {
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

	if p.PasswordBlockSimple != nil && (input.PasswordBlockSimple == nil || *p.PasswordBlockSimple != *input.PasswordBlockSimple) {
		return false
	}

	if p.PasswordExpirationDays != nil && (input.PasswordExpirationDays == nil || *p.PasswordExpirationDays != *input.PasswordExpirationDays) {
		return false
	}

	if p.PasswordMinimumCharacterSetCount != nil && (input.PasswordMinimumCharacterSetCount == nil || *p.PasswordMinimumCharacterSetCount != *input.PasswordMinimumCharacterSetCount) {
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

	if p.PasswordRequiredToUnlockFromIdle != nil && (input.PasswordRequiredToUnlockFromIdle == nil || *p.PasswordRequiredToUnlockFromIdle != *input.PasswordRequiredToUnlockFromIdle) {
		return false
	}

	if p.RequireHealthyDeviceReport != nil && (input.RequireHealthyDeviceReport == nil || *p.RequireHealthyDeviceReport != *input.RequireHealthyDeviceReport) {
		return false
	}

	if p.RtpEnabled != nil && (input.RtpEnabled == nil || *p.RtpEnabled != *input.RtpEnabled) {
		return false
	}

	if p.SecureBootEnabled != nil && (input.SecureBootEnabled == nil || *p.SecureBootEnabled != *input.SecureBootEnabled) {
		return false
	}

	if p.SignatureOutOfDate != nil && (input.SignatureOutOfDate == nil || *p.SignatureOutOfDate != *input.SignatureOutOfDate) {
		return false
	}

	if p.StorageRequireEncryption != nil && (input.StorageRequireEncryption == nil || *p.StorageRequireEncryption != *input.StorageRequireEncryption) {
		return false
	}

	if p.TpmRequired != nil && (input.TpmRequired == nil || *p.TpmRequired != *input.TpmRequired) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
