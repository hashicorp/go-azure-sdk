package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateCollectionTenantSummary struct {
	CompleteStepsCount                      *int64  `json:"completeStepsCount,omitempty"`
	CompleteUsersCount                      *int64  `json:"completeUsersCount,omitempty"`
	CreatedByUserId                         *string `json:"createdByUserId,omitempty"`
	CreatedDateTime                         *string `json:"createdDateTime,omitempty"`
	DismissedStepsCount                     *int64  `json:"dismissedStepsCount,omitempty"`
	ExcludedUsersCount                      *int64  `json:"excludedUsersCount,omitempty"`
	ExcludedUsersDistinctCount              *int64  `json:"excludedUsersDistinctCount,omitempty"`
	Id                                      *string `json:"id,omitempty"`
	IncompleteStepsCount                    *int64  `json:"incompleteStepsCount,omitempty"`
	IncompleteUsersCount                    *int64  `json:"incompleteUsersCount,omitempty"`
	IneligibleStepsCount                    *int64  `json:"ineligibleStepsCount,omitempty"`
	IsComplete                              *bool   `json:"isComplete,omitempty"`
	LastActionByUserId                      *string `json:"lastActionByUserId,omitempty"`
	LastActionDateTime                      *string `json:"lastActionDateTime,omitempty"`
	ManagementTemplateCollectionDisplayName *string `json:"managementTemplateCollectionDisplayName,omitempty"`
	ManagementTemplateCollectionId          *string `json:"managementTemplateCollectionId,omitempty"`
	ODataType                               *string `json:"@odata.type,omitempty"`
	RegressedStepsCount                     *int64  `json:"regressedStepsCount,omitempty"`
	RegressedUsersCount                     *int64  `json:"regressedUsersCount,omitempty"`
	TenantId                                *string `json:"tenantId,omitempty"`
	UnlicensedUsersCount                    *int64  `json:"unlicensedUsersCount,omitempty"`
}
