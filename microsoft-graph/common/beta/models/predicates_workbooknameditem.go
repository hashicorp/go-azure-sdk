package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookNamedItemOperationPredicate struct {
	Comment   *string
	Id        *string
	Name      *string
	ODataType *string
	Scope     *string
	Type      *string
	Visible   *bool
}

func (p WorkbookNamedItemOperationPredicate) Matches(input WorkbookNamedItem) bool {

	if p.Comment != nil && (input.Comment == nil || *p.Comment != *input.Comment) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Scope != nil && (input.Scope == nil || *p.Scope != *input.Scope) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	if p.Visible != nil && (input.Visible == nil || *p.Visible != *input.Visible) {
		return false
	}

	return true
}
