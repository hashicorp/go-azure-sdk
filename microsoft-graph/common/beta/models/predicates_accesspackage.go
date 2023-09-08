package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageOperationPredicate struct {
	CatalogId           *string
	CreatedBy           *string
	CreatedDateTime     *string
	Description         *string
	DisplayName         *string
	Id                  *string
	IsHidden            *bool
	IsRoleScopesVisible *bool
	ModifiedBy          *string
	ModifiedDateTime    *string
	ODataType           *string
}

func (p AccessPackageOperationPredicate) Matches(input AccessPackage) bool {

	if p.CatalogId != nil && (input.CatalogId == nil || *p.CatalogId != *input.CatalogId) {
		return false
	}

	if p.CreatedBy != nil && (input.CreatedBy == nil || *p.CreatedBy != *input.CreatedBy) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsHidden != nil && (input.IsHidden == nil || *p.IsHidden != *input.IsHidden) {
		return false
	}

	if p.IsRoleScopesVisible != nil && (input.IsRoleScopesVisible == nil || *p.IsRoleScopesVisible != *input.IsRoleScopesVisible) {
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
