package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateDetailedInfo struct {
	Category             *ManagedTenantsManagementTemplateDetailedInfoCategory `json:"category,omitempty"`
	DisplayName          *string                                               `json:"displayName,omitempty"`
	ManagementTemplateId *string                                               `json:"managementTemplateId,omitempty"`
	ODataType            *string                                               `json:"@odata.type,omitempty"`
	Version              *int64                                                `json:"version,omitempty"`
}
