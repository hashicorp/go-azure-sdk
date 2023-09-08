package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppAssignedDeviceLicenseOperationPredicate struct {
	DeviceName        *string
	Id                *string
	ManagedDeviceId   *string
	ODataType         *string
	UserEmailAddress  *string
	UserId            *string
	UserName          *string
	UserPrincipalName *string
}

func (p IosVppAppAssignedDeviceLicenseOperationPredicate) Matches(input IosVppAppAssignedDeviceLicense) bool {

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserEmailAddress != nil && (input.UserEmailAddress == nil || *p.UserEmailAddress != *input.UserEmailAddress) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserName != nil && (input.UserName == nil || *p.UserName != *input.UserName) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
