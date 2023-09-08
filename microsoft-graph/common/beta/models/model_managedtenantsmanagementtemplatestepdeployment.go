package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateStepDeployment struct {
	CreatedByUserId     *string                                               `json:"createdByUserId,omitempty"`
	CreatedDateTime     *string                                               `json:"createdDateTime,omitempty"`
	Error               *ManagedTenantsGraphAPIErrorDetails                   `json:"error,omitempty"`
	Id                  *string                                               `json:"id,omitempty"`
	LastActionByUserId  *string                                               `json:"lastActionByUserId,omitempty"`
	LastActionDateTime  *string                                               `json:"lastActionDateTime,omitempty"`
	ODataType           *string                                               `json:"@odata.type,omitempty"`
	Status              *ManagedTenantsManagementTemplateStepDeploymentStatus `json:"status,omitempty"`
	TemplateStepVersion *ManagedTenantsManagementTemplateStepVersion          `json:"templateStepVersion,omitempty"`
	TenantId            *string                                               `json:"tenantId,omitempty"`
}
