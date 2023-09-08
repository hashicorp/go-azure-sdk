package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookWorksheetOperationPredicate struct {
	Id         *string
	Name       *string
	ODataType  *string
	Position   *int64
	Visibility *string
}

func (p WorkbookWorksheetOperationPredicate) Matches(input WorkbookWorksheet) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Position != nil && (input.Position == nil || *p.Position != *input.Position) {
		return false
	}

	if p.Visibility != nil && (input.Visibility == nil || *p.Visibility != *input.Visibility) {
		return false
	}

	return true
}
