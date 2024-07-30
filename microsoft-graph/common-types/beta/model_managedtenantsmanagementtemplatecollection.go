package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateCollection struct {
	CreatedByUserId     *string                             `json:"createdByUserId,omitempty"`
	CreatedDateTime     *string                             `json:"createdDateTime,omitempty"`
	Description         *string                             `json:"description,omitempty"`
	DisplayName         *string                             `json:"displayName,omitempty"`
	Id                  *string                             `json:"id,omitempty"`
	LastActionByUserId  *string                             `json:"lastActionByUserId,omitempty"`
	LastActionDateTime  *string                             `json:"lastActionDateTime,omitempty"`
	ManagementTemplates *[]ManagedTenantsManagementTemplate `json:"managementTemplates,omitempty"`
	ODataType           *string                             `json:"@odata.type,omitempty"`
}
