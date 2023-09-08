package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsDeviceAccountActionParameterOperationPredicate struct {
	CalendarSyncEnabled              *bool
	DeviceAccountEmail               *string
	ExchangeServer                   *string
	ODataType                        *string
	PasswordRotationEnabled          *bool
	SessionInitiationProtocalAddress *string
}

func (p UpdateWindowsDeviceAccountActionParameterOperationPredicate) Matches(input UpdateWindowsDeviceAccountActionParameter) bool {

	if p.CalendarSyncEnabled != nil && (input.CalendarSyncEnabled == nil || *p.CalendarSyncEnabled != *input.CalendarSyncEnabled) {
		return false
	}

	if p.DeviceAccountEmail != nil && (input.DeviceAccountEmail == nil || *p.DeviceAccountEmail != *input.DeviceAccountEmail) {
		return false
	}

	if p.ExchangeServer != nil && (input.ExchangeServer == nil || *p.ExchangeServer != *input.ExchangeServer) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PasswordRotationEnabled != nil && (input.PasswordRotationEnabled == nil || *p.PasswordRotationEnabled != *input.PasswordRotationEnabled) {
		return false
	}

	if p.SessionInitiationProtocalAddress != nil && (input.SessionInitiationProtocalAddress == nil || *p.SessionInitiationProtocalAddress != *input.SessionInitiationProtocalAddress) {
		return false
	}

	return true
}
