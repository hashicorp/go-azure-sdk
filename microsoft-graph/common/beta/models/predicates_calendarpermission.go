package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarPermissionOperationPredicate struct {
	Id                   *string
	IsInsideOrganization *bool
	IsRemovable          *bool
	ODataType            *string
}

func (p CalendarPermissionOperationPredicate) Matches(input CalendarPermission) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsInsideOrganization != nil && (input.IsInsideOrganization == nil || *p.IsInsideOrganization != *input.IsInsideOrganization) {
		return false
	}

	if p.IsRemovable != nil && (input.IsRemovable == nil || *p.IsRemovable != *input.IsRemovable) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
