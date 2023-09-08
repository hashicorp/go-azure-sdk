package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosCompliancePolicyOperationPredicate struct {
	CreatedDateTime                       *string
	Description                           *string
	DeviceThreatProtectionEnabled         *bool
	DisplayName                           *string
	Id                                    *string
	LastModifiedDateTime                  *string
	ManagedEmailProfileRequired           *bool
	ODataType                             *string
	OsMaximumVersion                      *string
	OsMinimumVersion                      *string
	PasscodeBlockSimple                   *bool
	PasscodeExpirationDays                *int64
	PasscodeMinimumCharacterSetCount      *int64
	PasscodeMinimumLength                 *int64
	PasscodeMinutesOfInactivityBeforeLock *int64
	PasscodePreviousPasscodeBlockCount    *int64
	PasscodeRequired                      *bool
	SecurityBlockJailbrokenDevices        *bool
	Version                               *int64
}

func (p IosCompliancePolicyOperationPredicate) Matches(input IosCompliancePolicy) bool {

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

	if p.ManagedEmailProfileRequired != nil && (input.ManagedEmailProfileRequired == nil || *p.ManagedEmailProfileRequired != *input.ManagedEmailProfileRequired) {
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

	if p.PasscodeBlockSimple != nil && (input.PasscodeBlockSimple == nil || *p.PasscodeBlockSimple != *input.PasscodeBlockSimple) {
		return false
	}

	if p.PasscodeExpirationDays != nil && (input.PasscodeExpirationDays == nil || *p.PasscodeExpirationDays != *input.PasscodeExpirationDays) {
		return false
	}

	if p.PasscodeMinimumCharacterSetCount != nil && (input.PasscodeMinimumCharacterSetCount == nil || *p.PasscodeMinimumCharacterSetCount != *input.PasscodeMinimumCharacterSetCount) {
		return false
	}

	if p.PasscodeMinimumLength != nil && (input.PasscodeMinimumLength == nil || *p.PasscodeMinimumLength != *input.PasscodeMinimumLength) {
		return false
	}

	if p.PasscodeMinutesOfInactivityBeforeLock != nil && (input.PasscodeMinutesOfInactivityBeforeLock == nil || *p.PasscodeMinutesOfInactivityBeforeLock != *input.PasscodeMinutesOfInactivityBeforeLock) {
		return false
	}

	if p.PasscodePreviousPasscodeBlockCount != nil && (input.PasscodePreviousPasscodeBlockCount == nil || *p.PasscodePreviousPasscodeBlockCount != *input.PasscodePreviousPasscodeBlockCount) {
		return false
	}

	if p.PasscodeRequired != nil && (input.PasscodeRequired == nil || *p.PasscodeRequired != *input.PasscodeRequired) {
		return false
	}

	if p.SecurityBlockJailbrokenDevices != nil && (input.SecurityBlockJailbrokenDevices == nil || *p.SecurityBlockJailbrokenDevices != *input.SecurityBlockJailbrokenDevices) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
