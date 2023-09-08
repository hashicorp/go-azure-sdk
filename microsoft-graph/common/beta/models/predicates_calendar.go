package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarOperationPredicate struct {
	CalendarGroupId     *string
	CanEdit             *bool
	CanShare            *bool
	CanViewPrivateItems *bool
	ChangeKey           *string
	HexColor            *string
	Id                  *string
	IsDefaultCalendar   *bool
	IsRemovable         *bool
	IsShared            *bool
	IsSharedWithMe      *bool
	IsTallyingResponses *bool
	Name                *string
	ODataType           *string
}

func (p CalendarOperationPredicate) Matches(input Calendar) bool {

	if p.CalendarGroupId != nil && (input.CalendarGroupId == nil || *p.CalendarGroupId != *input.CalendarGroupId) {
		return false
	}

	if p.CanEdit != nil && (input.CanEdit == nil || *p.CanEdit != *input.CanEdit) {
		return false
	}

	if p.CanShare != nil && (input.CanShare == nil || *p.CanShare != *input.CanShare) {
		return false
	}

	if p.CanViewPrivateItems != nil && (input.CanViewPrivateItems == nil || *p.CanViewPrivateItems != *input.CanViewPrivateItems) {
		return false
	}

	if p.ChangeKey != nil && (input.ChangeKey == nil || *p.ChangeKey != *input.ChangeKey) {
		return false
	}

	if p.HexColor != nil && (input.HexColor == nil || *p.HexColor != *input.HexColor) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDefaultCalendar != nil && (input.IsDefaultCalendar == nil || *p.IsDefaultCalendar != *input.IsDefaultCalendar) {
		return false
	}

	if p.IsRemovable != nil && (input.IsRemovable == nil || *p.IsRemovable != *input.IsRemovable) {
		return false
	}

	if p.IsShared != nil && (input.IsShared == nil || *p.IsShared != *input.IsShared) {
		return false
	}

	if p.IsSharedWithMe != nil && (input.IsSharedWithMe == nil || *p.IsSharedWithMe != *input.IsSharedWithMe) {
		return false
	}

	if p.IsTallyingResponses != nil && (input.IsTallyingResponses == nil || *p.IsTallyingResponses != *input.IsTallyingResponses) {
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
