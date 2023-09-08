package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermColumnOperationPredicate struct {
	AllowMultipleValues    *bool
	ODataType              *string
	ShowFullyQualifiedName *bool
}

func (p TermColumnOperationPredicate) Matches(input TermColumn) bool {

	if p.AllowMultipleValues != nil && (input.AllowMultipleValues == nil || *p.AllowMultipleValues != *input.AllowMultipleValues) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ShowFullyQualifiedName != nil && (input.ShowFullyQualifiedName == nil || *p.ShowFullyQualifiedName != *input.ShowFullyQualifiedName) {
		return false
	}

	return true
}
