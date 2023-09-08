package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantContactInformationOperationPredicate struct {
	Email     *string
	Name      *string
	Notes     *string
	ODataType *string
	Phone     *string
	Title     *string
}

func (p ManagedTenantsTenantContactInformationOperationPredicate) Matches(input ManagedTenantsTenantContactInformation) bool {

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.Notes != nil && (input.Notes == nil || *p.Notes != *input.Notes) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Phone != nil && (input.Phone == nil || *p.Phone != *input.Phone) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
