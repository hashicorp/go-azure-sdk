package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementIntent struct {
	DisplayName         *string                                         `json:"displayName,omitempty"`
	Id                  *string                                         `json:"id,omitempty"`
	IsGlobal            *bool                                           `json:"isGlobal,omitempty"`
	ManagementTemplates *[]ManagedTenantsManagementTemplateDetailedInfo `json:"managementTemplates,omitempty"`
	ODataType           *string                                         `json:"@odata.type,omitempty"`
}
