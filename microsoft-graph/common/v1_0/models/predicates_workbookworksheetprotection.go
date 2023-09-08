package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookWorksheetProtectionOperationPredicate struct {
	Id        *string
	ODataType *string
	Protected *bool
}

func (p WorkbookWorksheetProtectionOperationPredicate) Matches(input WorkbookWorksheetProtection) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Protected != nil && (input.Protected == nil || *p.Protected != *input.Protected) {
		return false
	}

	return true
}
