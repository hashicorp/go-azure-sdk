package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantAlertOperationPredicate struct {
	AlertRuleDisplayName *string
	AssignedToUserId     *string
	CorrelationCount     *int64
	CorrelationId        *string
	CreatedByUserId      *string
	CreatedDateTime      *string
	Id                   *string
	LastActionByUserId   *string
	LastActionDateTime   *string
	Message              *string
	ODataType            *string
	TenantId             *string
	Title                *string
}

func (p ManagedTenantsManagedTenantAlertOperationPredicate) Matches(input ManagedTenantsManagedTenantAlert) bool {

	if p.AlertRuleDisplayName != nil && (input.AlertRuleDisplayName == nil || *p.AlertRuleDisplayName != *input.AlertRuleDisplayName) {
		return false
	}

	if p.AssignedToUserId != nil && (input.AssignedToUserId == nil || *p.AssignedToUserId != *input.AssignedToUserId) {
		return false
	}

	if p.CorrelationCount != nil && (input.CorrelationCount == nil || *p.CorrelationCount != *input.CorrelationCount) {
		return false
	}

	if p.CorrelationId != nil && (input.CorrelationId == nil || *p.CorrelationId != *input.CorrelationId) {
		return false
	}

	if p.CreatedByUserId != nil && (input.CreatedByUserId == nil || *p.CreatedByUserId != *input.CreatedByUserId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastActionByUserId != nil && (input.LastActionByUserId == nil || *p.LastActionByUserId != *input.LastActionByUserId) {
		return false
	}

	if p.LastActionDateTime != nil && (input.LastActionDateTime == nil || *p.LastActionDateTime != *input.LastActionDateTime) {
		return false
	}

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
