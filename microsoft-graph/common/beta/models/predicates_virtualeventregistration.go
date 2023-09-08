package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationOperationPredicate struct {
	CancelationDateTime  *string
	Email                *string
	FirstName            *string
	Id                   *string
	LastName             *string
	ODataType            *string
	RegistrationDateTime *string
	UserId               *string
}

func (p VirtualEventRegistrationOperationPredicate) Matches(input VirtualEventRegistration) bool {

	if p.CancelationDateTime != nil && (input.CancelationDateTime == nil || *p.CancelationDateTime != *input.CancelationDateTime) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.FirstName != nil && (input.FirstName == nil || *p.FirstName != *input.FirstName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastName != nil && (input.LastName == nil || *p.LastName != *input.LastName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RegistrationDateTime != nil && (input.RegistrationDateTime == nil || *p.RegistrationDateTime != *input.RegistrationDateTime) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
