package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookFormatProtectionOperationPredicate struct {
	FormulaHidden *bool
	Id            *string
	Locked        *bool
	ODataType     *string
}

func (p WorkbookFormatProtectionOperationPredicate) Matches(input WorkbookFormatProtection) bool {

	if p.FormulaHidden != nil && (input.FormulaHidden == nil || *p.FormulaHidden != *input.FormulaHidden) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Locked != nil && (input.Locked == nil || *p.Locked != *input.Locked) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
