package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingCustomerInformationOperationPredicate struct {
	CustomerId              *string
	EmailAddress            *string
	Name                    *string
	Notes                   *string
	ODataType               *string
	Phone                   *string
	SmsNotificationsEnabled *bool
	TimeZone                *string
}

func (p BookingCustomerInformationOperationPredicate) Matches(input BookingCustomerInformation) bool {

	if p.CustomerId != nil && (input.CustomerId == nil || *p.CustomerId != *input.CustomerId) {
		return false
	}

	if p.EmailAddress != nil && (input.EmailAddress == nil || *p.EmailAddress != *input.EmailAddress) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Notes != nil && (input.Notes == nil || *p.Notes != *input.Notes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Phone != nil && (input.Phone == nil || *p.Phone != *input.Phone) {
		return false
	}

	if p.SmsNotificationsEnabled != nil && (input.SmsNotificationsEnabled == nil || *p.SmsNotificationsEnabled != *input.SmsNotificationsEnabled) {
		return false
	}

	if p.TimeZone != nil && (input.TimeZone == nil || *p.TimeZone != *input.TimeZone) {
		return false
	}

	return true
}
