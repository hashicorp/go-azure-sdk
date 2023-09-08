package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionTenantDeploymentStatusOperationPredicate struct {
	Id            *string
	ODataType     *string
	TenantGroupId *string
	TenantId      *string
}

func (p ManagedTenantsManagementActionTenantDeploymentStatusOperationPredicate) Matches(input ManagedTenantsManagementActionTenantDeploymentStatus) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantGroupId != nil && (input.TenantGroupId == nil || *p.TenantGroupId != *input.TenantGroupId) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
