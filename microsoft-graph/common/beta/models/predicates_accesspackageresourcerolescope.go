package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceRoleScopeOperationPredicate struct {
	CreatedBy        *string
	CreatedDateTime  *string
	Id               *string
	ModifiedBy       *string
	ModifiedDateTime *string
	ODataType        *string
}

func (p AccessPackageResourceRoleScopeOperationPredicate) Matches(input AccessPackageResourceRoleScope) bool {

	if p.CreatedBy != nil && (input.CreatedBy == nil || *p.CreatedBy != *input.CreatedBy) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ModifiedBy != nil && (input.ModifiedBy == nil || *p.ModifiedBy != *input.ModifiedBy) {
		return false
	}

	if p.ModifiedDateTime != nil && (input.ModifiedDateTime == nil || *p.ModifiedDateTime != *input.ModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
