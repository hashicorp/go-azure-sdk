package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateCollectionTenantSummaryOperationPredicate struct {
	CompleteStepsCount                      *int64
	CompleteUsersCount                      *int64
	CreatedByUserId                         *string
	CreatedDateTime                         *string
	DismissedStepsCount                     *int64
	ExcludedUsersCount                      *int64
	ExcludedUsersDistinctCount              *int64
	Id                                      *string
	IncompleteStepsCount                    *int64
	IncompleteUsersCount                    *int64
	IneligibleStepsCount                    *int64
	IsComplete                              *bool
	LastActionByUserId                      *string
	LastActionDateTime                      *string
	ManagementTemplateCollectionDisplayName *string
	ManagementTemplateCollectionId          *string
	ODataType                               *string
	RegressedStepsCount                     *int64
	RegressedUsersCount                     *int64
	TenantId                                *string
	UnlicensedUsersCount                    *int64
}

func (p ManagedTenantsManagementTemplateCollectionTenantSummaryOperationPredicate) Matches(input ManagedTenantsManagementTemplateCollectionTenantSummary) bool {

	if p.CompleteStepsCount != nil && (input.CompleteStepsCount == nil || *p.CompleteStepsCount != *input.CompleteStepsCount) {
		return false
	}

	if p.CompleteUsersCount != nil && (input.CompleteUsersCount == nil || *p.CompleteUsersCount != *input.CompleteUsersCount) {
		return false
	}

	if p.CreatedByUserId != nil && (input.CreatedByUserId == nil || *p.CreatedByUserId != *input.CreatedByUserId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DismissedStepsCount != nil && (input.DismissedStepsCount == nil || *p.DismissedStepsCount != *input.DismissedStepsCount) {
		return false
	}

	if p.ExcludedUsersCount != nil && (input.ExcludedUsersCount == nil || *p.ExcludedUsersCount != *input.ExcludedUsersCount) {
		return false
	}

	if p.ExcludedUsersDistinctCount != nil && (input.ExcludedUsersDistinctCount == nil || *p.ExcludedUsersDistinctCount != *input.ExcludedUsersDistinctCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IncompleteStepsCount != nil && (input.IncompleteStepsCount == nil || *p.IncompleteStepsCount != *input.IncompleteStepsCount) {
		return false
	}

	if p.IncompleteUsersCount != nil && (input.IncompleteUsersCount == nil || *p.IncompleteUsersCount != *input.IncompleteUsersCount) {
		return false
	}

	if p.IneligibleStepsCount != nil && (input.IneligibleStepsCount == nil || *p.IneligibleStepsCount != *input.IneligibleStepsCount) {
		return false
	}

	if p.IsComplete != nil && (input.IsComplete == nil || *p.IsComplete != *input.IsComplete) {
		return false
	}

	if p.LastActionByUserId != nil && (input.LastActionByUserId == nil || *p.LastActionByUserId != *input.LastActionByUserId) {
		return false
	}

	if p.LastActionDateTime != nil && (input.LastActionDateTime == nil || *p.LastActionDateTime != *input.LastActionDateTime) {
		return false
	}

	if p.ManagementTemplateCollectionDisplayName != nil && (input.ManagementTemplateCollectionDisplayName == nil || *p.ManagementTemplateCollectionDisplayName != *input.ManagementTemplateCollectionDisplayName) {
		return false
	}

	if p.ManagementTemplateCollectionId != nil && (input.ManagementTemplateCollectionId == nil || *p.ManagementTemplateCollectionId != *input.ManagementTemplateCollectionId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RegressedStepsCount != nil && (input.RegressedStepsCount == nil || *p.RegressedStepsCount != *input.RegressedStepsCount) {
		return false
	}

	if p.RegressedUsersCount != nil && (input.RegressedUsersCount == nil || *p.RegressedUsersCount != *input.RegressedUsersCount) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.UnlicensedUsersCount != nil && (input.UnlicensedUsersCount == nil || *p.UnlicensedUsersCount != *input.UnlicensedUsersCount) {
		return false
	}

	return true
}
