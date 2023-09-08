package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceReferenceOperationPredicate struct {
	Id        *string
	ODataType *string
	Type      *string
	WebUrl    *string
}

func (p ResourceReferenceOperationPredicate) Matches(input ResourceReference) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
