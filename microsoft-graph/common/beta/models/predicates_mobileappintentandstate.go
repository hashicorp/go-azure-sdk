package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppIntentAndStateOperationPredicate struct {
	Id                      *string
	ManagedDeviceIdentifier *string
	ODataType               *string
	UserId                  *string
}

func (p MobileAppIntentAndStateOperationPredicate) Matches(input MobileAppIntentAndState) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ManagedDeviceIdentifier != nil && (input.ManagedDeviceIdentifier == nil || *p.ManagedDeviceIdentifier != *input.ManagedDeviceIdentifier) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
