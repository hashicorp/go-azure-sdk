package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ODataErrorsMainErrorOperationPredicate struct {
	Code    *string
	Message *string
	Target  *string
}

func (p ODataErrorsMainErrorOperationPredicate) Matches(input ODataErrorsMainError) bool {

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.Target != nil && (input.Target == nil || *p.Target != *input.Target) {
		return false
	}

	return true
}
