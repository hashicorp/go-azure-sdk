package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptUserStateOperationPredicate struct {
	ErrorDeviceCount   *int64
	Id                 *string
	ODataType          *string
	SuccessDeviceCount *int64
	UserPrincipalName  *string
}

func (p DeviceManagementScriptUserStateOperationPredicate) Matches(input DeviceManagementScriptUserState) bool {

	if p.ErrorDeviceCount != nil && (input.ErrorDeviceCount == nil || *p.ErrorDeviceCount != *input.ErrorDeviceCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SuccessDeviceCount != nil && (input.SuccessDeviceCount == nil || *p.SuccessDeviceCount != *input.SuccessDeviceCount) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
