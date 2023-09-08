package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageQuestionOperationPredicate struct {
	Id               *string
	IsAnswerEditable *bool
	IsRequired       *bool
	ODataType        *string
	Sequence         *int64
	Text             *string
}

func (p AccessPackageQuestionOperationPredicate) Matches(input AccessPackageQuestion) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsAnswerEditable != nil && (input.IsAnswerEditable == nil || *p.IsAnswerEditable != *input.IsAnswerEditable) {
		return false
	}

	if p.IsRequired != nil && (input.IsRequired == nil || *p.IsRequired != *input.IsRequired) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Sequence != nil && (input.Sequence == nil || *p.Sequence != *input.Sequence) {
		return false
	}

	if p.Text != nil && (input.Text == nil || *p.Text != *input.Text) {
		return false
	}

	return true
}
