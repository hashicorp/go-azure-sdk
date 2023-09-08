package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingRegistrationQuestionOperationPredicate struct {
	DisplayName *string
	Id          *string
	IsRequired  *bool
	ODataType   *string
}

func (p MeetingRegistrationQuestionOperationPredicate) Matches(input MeetingRegistrationQuestion) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsRequired != nil && (input.IsRequired == nil || *p.IsRequired != *input.IsRequired) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
