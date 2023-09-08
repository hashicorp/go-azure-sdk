package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantGenericErrorOperationPredicate struct {
	Error     *string
	ODataType *string
	TenantId  *string
}

func (p ManagedTenantsManagedTenantGenericErrorOperationPredicate) Matches(input ManagedTenantsManagedTenantGenericError) bool {

	if p.Error != nil && (input.Error == nil || *p.Error != *input.Error) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
