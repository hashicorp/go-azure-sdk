package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantGroup struct {
	AllTenantsIncluded *bool                                 `json:"allTenantsIncluded,omitempty"`
	DisplayName        *string                               `json:"displayName,omitempty"`
	Id                 *string                               `json:"id,omitempty"`
	ManagementActions  *[]ManagedTenantsManagementActionInfo `json:"managementActions,omitempty"`
	ManagementIntents  *[]ManagedTenantsManagementIntentInfo `json:"managementIntents,omitempty"`
	ODataType          *string                               `json:"@odata.type,omitempty"`
	TenantIds          *[]string                             `json:"tenantIds,omitempty"`
}
