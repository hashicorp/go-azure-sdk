package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateStepTenantSummary struct {
	AssignedTenantsCount                    *int64  `json:"assignedTenantsCount,omitempty"`
	CompliantTenantsCount                   *int64  `json:"compliantTenantsCount,omitempty"`
	CreatedByUserId                         *string `json:"createdByUserId,omitempty"`
	CreatedDateTime                         *string `json:"createdDateTime,omitempty"`
	DismissedTenantsCount                   *int64  `json:"dismissedTenantsCount,omitempty"`
	Id                                      *string `json:"id,omitempty"`
	IneligibleTenantsCount                  *int64  `json:"ineligibleTenantsCount,omitempty"`
	LastActionByUserId                      *string `json:"lastActionByUserId,omitempty"`
	LastActionDateTime                      *string `json:"lastActionDateTime,omitempty"`
	ManagementTemplateCollectionDisplayName *string `json:"managementTemplateCollectionDisplayName,omitempty"`
	ManagementTemplateCollectionId          *string `json:"managementTemplateCollectionId,omitempty"`
	ManagementTemplateDisplayName           *string `json:"managementTemplateDisplayName,omitempty"`
	ManagementTemplateId                    *string `json:"managementTemplateId,omitempty"`
	ManagementTemplateStepDisplayName       *string `json:"managementTemplateStepDisplayName,omitempty"`
	ManagementTemplateStepId                *string `json:"managementTemplateStepId,omitempty"`
	NotCompliantTenantsCount                *int64  `json:"notCompliantTenantsCount,omitempty"`
	ODataType                               *string `json:"@odata.type,omitempty"`
}
