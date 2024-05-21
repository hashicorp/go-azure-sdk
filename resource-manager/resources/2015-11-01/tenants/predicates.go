package tenants

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantIdDescriptionOperationPredicate struct {
	Id       *string
	TenantId *string
}

func (p TenantIdDescriptionOperationPredicate) Matches(input TenantIdDescription) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
