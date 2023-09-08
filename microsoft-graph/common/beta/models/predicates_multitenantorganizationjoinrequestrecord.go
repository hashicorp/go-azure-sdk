package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestRecordOperationPredicate struct {
	AddedByTenantId *string
	Id              *string
	ODataType       *string
}

func (p MultiTenantOrganizationJoinRequestRecordOperationPredicate) Matches(input MultiTenantOrganizationJoinRequestRecord) bool {

	if p.AddedByTenantId != nil && (input.AddedByTenantId == nil || *p.AddedByTenantId != *input.AddedByTenantId) {
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
