package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantGroupOperationPredicate struct {
	AllTenantsIncluded *bool
	DisplayName        *string
	Id                 *string
	ODataType          *string
}

func (p ManagedTenantsTenantGroupOperationPredicate) Matches(input ManagedTenantsTenantGroup) bool {

	if p.AllTenantsIncluded != nil && (input.AllTenantsIncluded == nil || *p.AllTenantsIncluded != *input.AllTenantsIncluded) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
