package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentDeviceStateOperationPredicate struct {
	DeviceDisplayName    *string
	DeviceId             *string
	Id                   *string
	LastReportedDateTime *string
	ODataType            *string
	UserName             *string
	UserPrincipalName    *string
}

func (p DeviceManagementIntentDeviceStateOperationPredicate) Matches(input DeviceManagementIntentDeviceState) bool {

	if p.DeviceDisplayName != nil && (input.DeviceDisplayName == nil || *p.DeviceDisplayName != *input.DeviceDisplayName) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastReportedDateTime != nil && (input.LastReportedDateTime == nil || *p.LastReportedDateTime != *input.LastReportedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
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
