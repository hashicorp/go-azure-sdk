package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantInfoOperationPredicate struct {
	ODataType *string
	TenantId  *string
}

func (p ManagedTenantsTenantInfoOperationPredicate) Matches(input ManagedTenantsTenantInfo) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
