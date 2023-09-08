package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSimulationEventInfoOperationPredicate struct {
	Browser                 *string
	EventDateTime           *string
	EventName               *string
	IpAddress               *string
	ODataType               *string
	OsPlatformDeviceDetails *string
}

func (p UserSimulationEventInfoOperationPredicate) Matches(input UserSimulationEventInfo) bool {

	if p.Browser != nil && (input.Browser == nil || *p.Browser != *input.Browser) {
		return false
	}

	if p.EventDateTime != nil && (input.EventDateTime == nil || *p.EventDateTime != *input.EventDateTime) {
		return false
	}

	if p.EventName != nil && (input.EventName == nil || *p.EventName != *input.EventName) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OsPlatformDeviceDetails != nil && (input.OsPlatformDeviceDetails == nil || *p.OsPlatformDeviceDetails != *input.OsPlatformDeviceDetails) {
		return false
	}

	return true
}
