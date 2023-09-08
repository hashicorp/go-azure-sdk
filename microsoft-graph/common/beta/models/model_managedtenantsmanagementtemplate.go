package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplate struct {
	Category                      *ManagedTenantsManagementTemplateCategory     `json:"category,omitempty"`
	CreatedByUserId               *string                                       `json:"createdByUserId,omitempty"`
	CreatedDateTime               *string                                       `json:"createdDateTime,omitempty"`
	Description                   *string                                       `json:"description,omitempty"`
	DisplayName                   *string                                       `json:"displayName,omitempty"`
	Id                            *string                                       `json:"id,omitempty"`
	InformationLinks              *[]ActionUrl                                  `json:"informationLinks,omitempty"`
	LastActionByUserId            *string                                       `json:"lastActionByUserId,omitempty"`
	LastActionDateTime            *string                                       `json:"lastActionDateTime,omitempty"`
	ManagementTemplateCollections *[]ManagedTenantsManagementTemplateCollection `json:"managementTemplateCollections,omitempty"`
	ManagementTemplateSteps       *[]ManagedTenantsManagementTemplateStep       `json:"managementTemplateSteps,omitempty"`
	ODataType                     *string                                       `json:"@odata.type,omitempty"`
	Parameters                    *[]ManagedTenantsTemplateParameter            `json:"parameters,omitempty"`
	Priority                      *int64                                        `json:"priority,omitempty"`
	Provider                      *ManagedTenantsManagementTemplateProvider     `json:"provider,omitempty"`
	UserImpact                    *string                                       `json:"userImpact,omitempty"`
	Version                       *int64                                        `json:"version,omitempty"`
	WorkloadActions               *[]ManagedTenantsWorkloadAction               `json:"workloadActions,omitempty"`
}
