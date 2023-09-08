package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationSettingStateOperationPredicate struct {
	CurrentValue        *string
	ErrorCode           *int64
	ErrorDescription    *string
	InstanceDisplayName *string
	ODataType           *string
	Setting             *string
	SettingName         *string
	UserEmail           *string
	UserId              *string
	UserName            *string
	UserPrincipalName   *string
}

func (p DeviceConfigurationSettingStateOperationPredicate) Matches(input DeviceConfigurationSettingState) bool {

	if p.CurrentValue != nil && (input.CurrentValue == nil || *p.CurrentValue != *input.CurrentValue) {
		return false
	}

	if p.ErrorCode != nil && (input.ErrorCode == nil || *p.ErrorCode != *input.ErrorCode) {
		return false
	}

	if p.ErrorDescription != nil && (input.ErrorDescription == nil || *p.ErrorDescription != *input.ErrorDescription) {
		return false
	}

	if p.InstanceDisplayName != nil && (input.InstanceDisplayName == nil || *p.InstanceDisplayName != *input.InstanceDisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Setting != nil && (input.Setting == nil || *p.Setting != *input.Setting) {
		return false
	}

	if p.SettingName != nil && (input.SettingName == nil || *p.SettingName != *input.SettingName) {
		return false
	}

	if p.UserEmail != nil && (input.UserEmail == nil || *p.UserEmail != *input.UserEmail) {
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
