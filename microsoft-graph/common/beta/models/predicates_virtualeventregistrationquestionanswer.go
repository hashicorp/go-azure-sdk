package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventRegistrationQuestionAnswerOperationPredicate struct {
	BooleanValue *bool
	DisplayName  *string
	ODataType    *string
	QuestionId   *string
	Value        *string
}

func (p VirtualEventRegistrationQuestionAnswerOperationPredicate) Matches(input VirtualEventRegistrationQuestionAnswer) bool {

	if p.BooleanValue != nil && (input.BooleanValue == nil || *p.BooleanValue != *input.BooleanValue) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.QuestionId != nil && (input.QuestionId == nil || *p.QuestionId != *input.QuestionId) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
