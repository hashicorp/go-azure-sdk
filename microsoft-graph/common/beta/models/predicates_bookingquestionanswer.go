package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingQuestionAnswerOperationPredicate struct {
	Answer     *string
	IsRequired *bool
	ODataType  *string
	Question   *string
	QuestionId *string
}

func (p BookingQuestionAnswerOperationPredicate) Matches(input BookingQuestionAnswer) bool {

	if p.Answer != nil && (input.Answer == nil || *p.Answer != *input.Answer) {
		return false
	}

	if p.IsRequired != nil && (input.IsRequired == nil || *p.IsRequired != *input.IsRequired) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Question != nil && (input.Question == nil || *p.Question != *input.Question) {
		return false
	}

	if p.QuestionId != nil && (input.QuestionId == nil || *p.QuestionId != *input.QuestionId) {
		return false
	}

	return true
}
