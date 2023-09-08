package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationResultOperationPredicate struct {
	Message          *string
	ODataType        *string
	RuleName         *string
	ValidationPassed *bool
}

func (p ValidationResultOperationPredicate) Matches(input ValidationResult) bool {

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RuleName != nil && (input.RuleName == nil || *p.RuleName != *input.RuleName) {
		return false
	}

	if p.ValidationPassed != nil && (input.ValidationPassed == nil || *p.ValidationPassed != *input.ValidationPassed) {
		return false
	}

	return true
}
