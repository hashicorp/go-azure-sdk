package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationOperationPredicate struct {
	Description               *string
	EndDateTime               *string
	Id                        *string
	ODataType                 *string
	RegistrationPageViewCount *int64
	RegistrationPageWebUrl    *string
	StartDateTime             *string
	Subject                   *string
}

func (p MeetingRegistrationOperationPredicate) Matches(input MeetingRegistration) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RegistrationPageViewCount != nil && (input.RegistrationPageViewCount == nil || *p.RegistrationPageViewCount != *input.RegistrationPageViewCount) {
		return false
	}

	if p.RegistrationPageWebUrl != nil && (input.RegistrationPageWebUrl == nil || *p.RegistrationPageWebUrl != *input.RegistrationPageWebUrl) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	return true
}
