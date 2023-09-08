package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAnswerChoiceOperationPredicate struct {
	ActualValue *string
	ODataType   *string
	Text        *string
}

func (p AccessPackageAnswerChoiceOperationPredicate) Matches(input AccessPackageAnswerChoice) bool {

	if p.ActualValue != nil && (input.ActualValue == nil || *p.ActualValue != *input.ActualValue) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Text != nil && (input.Text == nil || *p.Text != *input.Text) {
		return false
	}

	return true
}
