package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateStepTenantSummaryOperationPredicate struct {
	AssignedTenantsCount                    *int64
	CompliantTenantsCount                   *int64
	CreatedByUserId                         *string
	CreatedDateTime                         *string
	DismissedTenantsCount                   *int64
	Id                                      *string
	IneligibleTenantsCount                  *int64
	LastActionByUserId                      *string
	LastActionDateTime                      *string
	ManagementTemplateCollectionDisplayName *string
	ManagementTemplateCollectionId          *string
	ManagementTemplateDisplayName           *string
	ManagementTemplateId                    *string
	ManagementTemplateStepDisplayName       *string
	ManagementTemplateStepId                *string
	NotCompliantTenantsCount                *int64
	ODataType                               *string
}

func (p ManagedTenantsManagementTemplateStepTenantSummaryOperationPredicate) Matches(input ManagedTenantsManagementTemplateStepTenantSummary) bool {

	if p.AssignedTenantsCount != nil && (input.AssignedTenantsCount == nil || *p.AssignedTenantsCount != *input.AssignedTenantsCount) {
		return false
	}

	if p.CompliantTenantsCount != nil && (input.CompliantTenantsCount == nil || *p.CompliantTenantsCount != *input.CompliantTenantsCount) {
		return false
	}

	if p.CreatedByUserId != nil && (input.CreatedByUserId == nil || *p.CreatedByUserId != *input.CreatedByUserId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DismissedTenantsCount != nil && (input.DismissedTenantsCount == nil || *p.DismissedTenantsCount != *input.DismissedTenantsCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IneligibleTenantsCount != nil && (input.IneligibleTenantsCount == nil || *p.IneligibleTenantsCount != *input.IneligibleTenantsCount) {
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

	if p.ManagementTemplateDisplayName != nil && (input.ManagementTemplateDisplayName == nil || *p.ManagementTemplateDisplayName != *input.ManagementTemplateDisplayName) {
		return false
	}

	if p.ManagementTemplateId != nil && (input.ManagementTemplateId == nil || *p.ManagementTemplateId != *input.ManagementTemplateId) {
		return false
	}

	if p.ManagementTemplateStepDisplayName != nil && (input.ManagementTemplateStepDisplayName == nil || *p.ManagementTemplateStepDisplayName != *input.ManagementTemplateStepDisplayName) {
		return false
	}

	if p.ManagementTemplateStepId != nil && (input.ManagementTemplateStepId == nil || *p.ManagementTemplateStepId != *input.ManagementTemplateStepId) {
		return false
	}

	if p.NotCompliantTenantsCount != nil && (input.NotCompliantTenantsCount == nil || *p.NotCompliantTenantsCount != *input.NotCompliantTenantsCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
