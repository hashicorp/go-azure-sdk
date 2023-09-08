package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarGroupOperationPredicate struct {
	ChangeKey *string
	ClassId   *string
	Id        *string
	Name      *string
	ODataType *string
}

func (p CalendarGroupOperationPredicate) Matches(input CalendarGroup) bool {

	if p.ChangeKey != nil && (input.ChangeKey == nil || *p.ChangeKey != *input.ChangeKey) {
		return false
	}

	if p.ClassId != nil && (input.ClassId == nil || *p.ClassId != *input.ClassId) {
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

	return true
}
