package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberOperationPredicate struct {
	AddedByTenantId *string
	AddedDateTime   *string
	DeletedDateTime *string
	DisplayName     *string
	Id              *string
	JoinedDateTime  *string
	ODataType       *string
	TenantId        *string
}

func (p MultiTenantOrganizationMemberOperationPredicate) Matches(input MultiTenantOrganizationMember) bool {

	if p.AddedByTenantId != nil && (input.AddedByTenantId == nil || *p.AddedByTenantId != *input.AddedByTenantId) {
		return false
	}

	if p.AddedDateTime != nil && (input.AddedDateTime == nil || *p.AddedDateTime != *input.AddedDateTime) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.JoinedDateTime != nil && (input.JoinedDateTime == nil || *p.JoinedDateTime != *input.JoinedDateTime) {
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
