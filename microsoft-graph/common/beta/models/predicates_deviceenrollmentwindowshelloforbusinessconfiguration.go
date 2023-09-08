package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationOperationPredicate struct {
	CreatedDateTime             *string
	Description                 *string
	DisplayName                 *string
	EnhancedSignInSecurity      *int64
	Id                          *string
	LastModifiedDateTime        *string
	ODataType                   *string
	PinExpirationInDays         *int64
	PinMaximumLength            *int64
	PinMinimumLength            *int64
	PinPreviousBlockCount       *int64
	Priority                    *int64
	RemotePassportEnabled       *bool
	SecurityDeviceRequired      *bool
	UnlockWithBiometricsEnabled *bool
	Version                     *int64
}

func (p DeviceEnrollmentWindowsHelloForBusinessConfigurationOperationPredicate) Matches(input DeviceEnrollmentWindowsHelloForBusinessConfiguration) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EnhancedSignInSecurity != nil && (input.EnhancedSignInSecurity == nil || *p.EnhancedSignInSecurity != *input.EnhancedSignInSecurity) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PinExpirationInDays != nil && (input.PinExpirationInDays == nil || *p.PinExpirationInDays != *input.PinExpirationInDays) {
		return false
	}

	if p.PinMaximumLength != nil && (input.PinMaximumLength == nil || *p.PinMaximumLength != *input.PinMaximumLength) {
		return false
	}

	if p.PinMinimumLength != nil && (input.PinMinimumLength == nil || *p.PinMinimumLength != *input.PinMinimumLength) {
		return false
	}

	if p.PinPreviousBlockCount != nil && (input.PinPreviousBlockCount == nil || *p.PinPreviousBlockCount != *input.PinPreviousBlockCount) {
		return false
	}

	if p.Priority != nil && (input.Priority == nil || *p.Priority != *input.Priority) {
		return false
	}

	if p.RemotePassportEnabled != nil && (input.RemotePassportEnabled == nil || *p.RemotePassportEnabled != *input.RemotePassportEnabled) {
		return false
	}

	if p.SecurityDeviceRequired != nil && (input.SecurityDeviceRequired == nil || *p.SecurityDeviceRequired != *input.SecurityDeviceRequired) {
		return false
	}

	if p.UnlockWithBiometricsEnabled != nil && (input.UnlockWithBiometricsEnabled == nil || *p.UnlockWithBiometricsEnabled != *input.UnlockWithBiometricsEnabled) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
