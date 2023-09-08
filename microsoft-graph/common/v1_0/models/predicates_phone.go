package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoneOperationPredicate struct {
	Language  *string
	Number    *string
	ODataType *string
	Region    *string
}

func (p PhoneOperationPredicate) Matches(input Phone) bool {

	if p.Language != nil && (input.Language == nil || *p.Language != *input.Language) {
		return false
	}

	if p.Number != nil && (input.Number == nil || *p.Number != *input.Number) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Region != nil && (input.Region == nil || *p.Region != *input.Region) {
		return false
	}

	return true
}
