package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateStep struct {
	AcceptedVersion    *ManagedTenantsManagementTemplateStepVersion   `json:"acceptedVersion,omitempty"`
	Category           *ManagedTenantsManagementTemplateStepCategory  `json:"category,omitempty"`
	CreatedByUserId    *string                                        `json:"createdByUserId,omitempty"`
	CreatedDateTime    *string                                        `json:"createdDateTime,omitempty"`
	Description        *string                                        `json:"description,omitempty"`
	DisplayName        *string                                        `json:"displayName,omitempty"`
	Id                 *string                                        `json:"id,omitempty"`
	LastActionByUserId *string                                        `json:"lastActionByUserId,omitempty"`
	LastActionDateTime *string                                        `json:"lastActionDateTime,omitempty"`
	ManagementTemplate *ManagedTenantsManagementTemplate              `json:"managementTemplate,omitempty"`
	ODataType          *string                                        `json:"@odata.type,omitempty"`
	PortalLink         *ActionUrl                                     `json:"portalLink,omitempty"`
	Priority           *int64                                         `json:"priority,omitempty"`
	Versions           *[]ManagedTenantsManagementTemplateStepVersion `json:"versions,omitempty"`
}
