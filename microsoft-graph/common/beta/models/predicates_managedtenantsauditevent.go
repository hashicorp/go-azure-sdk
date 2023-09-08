package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAuditEventOperationPredicate struct {
	Activity          *string
	ActivityDateTime  *string
	ActivityId        *string
	Category          *string
	HttpVerb          *string
	Id                *string
	InitiatedByAppId  *string
	InitiatedByUpn    *string
	InitiatedByUserId *string
	IpAddress         *string
	ODataType         *string
	RequestBody       *string
	RequestUrl        *string
	TenantIds         *string
	TenantNames       *string
}

func (p ManagedTenantsAuditEventOperationPredicate) Matches(input ManagedTenantsAuditEvent) bool {

	if p.Activity != nil && (input.Activity == nil || *p.Activity != *input.Activity) {
		return false
	}

	if p.ActivityDateTime != nil && (input.ActivityDateTime == nil || *p.ActivityDateTime != *input.ActivityDateTime) {
		return false
	}

	if p.ActivityId != nil && (input.ActivityId == nil || *p.ActivityId != *input.ActivityId) {
		return false
	}

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.HttpVerb != nil && (input.HttpVerb == nil || *p.HttpVerb != *input.HttpVerb) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InitiatedByAppId != nil && (input.InitiatedByAppId == nil || *p.InitiatedByAppId != *input.InitiatedByAppId) {
		return false
	}

	if p.InitiatedByUpn != nil && (input.InitiatedByUpn == nil || *p.InitiatedByUpn != *input.InitiatedByUpn) {
		return false
	}

	if p.InitiatedByUserId != nil && (input.InitiatedByUserId == nil || *p.InitiatedByUserId != *input.InitiatedByUserId) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequestBody != nil && (input.RequestBody == nil || *p.RequestBody != *input.RequestBody) {
		return false
	}

	if p.RequestUrl != nil && (input.RequestUrl == nil || *p.RequestUrl != *input.RequestUrl) {
		return false
	}

	if p.TenantIds != nil && (input.TenantIds == nil || *p.TenantIds != *input.TenantIds) {
		return false
	}

	if p.TenantNames != nil && (input.TenantNames == nil || *p.TenantNames != *input.TenantNames) {
		return false
	}

	return true
}
