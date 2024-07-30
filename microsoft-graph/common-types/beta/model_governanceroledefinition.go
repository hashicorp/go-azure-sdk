package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRoleDefinition struct {
	DisplayName *string                `json:"displayName,omitempty"`
	ExternalId  *string                `json:"externalId,omitempty"`
	Id          *string                `json:"id,omitempty"`
	ODataType   *string                `json:"@odata.type,omitempty"`
	Resource    *GovernanceResource    `json:"resource,omitempty"`
	ResourceId  *string                `json:"resourceId,omitempty"`
	RoleSetting *GovernanceRoleSetting `json:"roleSetting,omitempty"`
	TemplateId  *string                `json:"templateId,omitempty"`
}
