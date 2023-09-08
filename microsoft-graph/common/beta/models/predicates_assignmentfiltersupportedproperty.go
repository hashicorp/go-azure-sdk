package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterSupportedPropertyOperationPredicate struct {
	DataType                *string
	IsCollection            *bool
	Name                    *string
	ODataType               *string
	PropertyRegexConstraint *string
}

func (p AssignmentFilterSupportedPropertyOperationPredicate) Matches(input AssignmentFilterSupportedProperty) bool {

	if p.DataType != nil && (input.DataType == nil || *p.DataType != *input.DataType) {
		return false
	}

	if p.IsCollection != nil && (input.IsCollection == nil || *p.IsCollection != *input.IsCollection) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PropertyRegexConstraint != nil && (input.PropertyRegexConstraint == nil || *p.PropertyRegexConstraint != *input.PropertyRegexConstraint) {
		return false
	}

	return true
}
