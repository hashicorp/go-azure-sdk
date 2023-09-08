package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegistrationPolicyOperationPredicate struct {
	Description     *string
	DisplayName     *string
	Id              *string
	ODataType       *string
	UserDeviceQuota *int64
}

func (p DeviceRegistrationPolicyOperationPredicate) Matches(input DeviceRegistrationPolicy) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserDeviceQuota != nil && (input.UserDeviceQuota == nil || *p.UserDeviceQuota != *input.UserDeviceQuota) {
		return false
	}

	return true
}
