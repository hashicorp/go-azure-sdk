package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionTenantDeploymentStatus struct {
	Id            *string                                           `json:"id,omitempty"`
	ODataType     *string                                           `json:"@odata.type,omitempty"`
	Statuses      *[]ManagedTenantsManagementActionDeploymentStatus `json:"statuses,omitempty"`
	TenantGroupId *string                                           `json:"tenantGroupId,omitempty"`
	TenantId      *string                                           `json:"tenantId,omitempty"`
}
