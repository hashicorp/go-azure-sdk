package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppDeviceDetailsOperationPredicate struct {
	AppVersion      *string
	ClientApp       *string
	DeviceId        *string
	ODataType       *string
	OperatingSystem *string
}

func (p AuthenticationAppDeviceDetailsOperationPredicate) Matches(input AuthenticationAppDeviceDetails) bool {

	if p.AppVersion != nil && (input.AppVersion == nil || *p.AppVersion != *input.AppVersion) {
		return false
	}

	if p.ClientApp != nil && (input.ClientApp == nil || *p.ClientApp != *input.ClientApp) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OperatingSystem != nil && (input.OperatingSystem == nil || *p.OperatingSystem != *input.OperatingSystem) {
		return false
	}

	return true
}
