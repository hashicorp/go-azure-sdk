package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceExchangeAccessStateSummaryOperationPredicate struct {
	AllowedDeviceCount     *int64
	BlockedDeviceCount     *int64
	ODataType              *string
	QuarantinedDeviceCount *int64
	UnavailableDeviceCount *int64
	UnknownDeviceCount     *int64
}

func (p DeviceExchangeAccessStateSummaryOperationPredicate) Matches(input DeviceExchangeAccessStateSummary) bool {

	if p.AllowedDeviceCount != nil && (input.AllowedDeviceCount == nil || *p.AllowedDeviceCount != *input.AllowedDeviceCount) {
		return false
	}

	if p.BlockedDeviceCount != nil && (input.BlockedDeviceCount == nil || *p.BlockedDeviceCount != *input.BlockedDeviceCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.QuarantinedDeviceCount != nil && (input.QuarantinedDeviceCount == nil || *p.QuarantinedDeviceCount != *input.QuarantinedDeviceCount) {
		return false
	}

	if p.UnavailableDeviceCount != nil && (input.UnavailableDeviceCount == nil || *p.UnavailableDeviceCount != *input.UnavailableDeviceCount) {
		return false
	}

	if p.UnknownDeviceCount != nil && (input.UnknownDeviceCount == nil || *p.UnknownDeviceCount != *input.UnknownDeviceCount) {
		return false
	}

	return true
}
