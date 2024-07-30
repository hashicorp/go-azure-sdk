package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionInfo struct {
	ManagementActionId        *string `json:"managementActionId,omitempty"`
	ManagementTemplateId      *string `json:"managementTemplateId,omitempty"`
	ManagementTemplateVersion *int64  `json:"managementTemplateVersion,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
}
