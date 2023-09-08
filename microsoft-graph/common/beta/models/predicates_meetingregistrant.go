package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrantOperationPredicate struct {
	Email                *string
	FirstName            *string
	Id                   *string
	JoinWebUrl           *string
	LastName             *string
	ODataType            *string
	RegistrationDateTime *string
}

func (p MeetingRegistrantOperationPredicate) Matches(input MeetingRegistrant) bool {

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.FirstName != nil && (input.FirstName == nil || *p.FirstName != *input.FirstName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.JoinWebUrl != nil && (input.JoinWebUrl == nil || *p.JoinWebUrl != *input.JoinWebUrl) {
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

	return true
}
