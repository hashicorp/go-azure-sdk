package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRoleSetting struct {
	AdminEligibleSettings *[]GovernanceRuleSetting  `json:"adminEligibleSettings,omitempty"`
	AdminMemberSettings   *[]GovernanceRuleSetting  `json:"adminMemberSettings,omitempty"`
	Id                    *string                   `json:"id,omitempty"`
	IsDefault             *bool                     `json:"isDefault,omitempty"`
	LastUpdatedBy         *string                   `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDateTime   *string                   `json:"lastUpdatedDateTime,omitempty"`
	ODataType             *string                   `json:"@odata.type,omitempty"`
	Resource              *GovernanceResource       `json:"resource,omitempty"`
	ResourceId            *string                   `json:"resourceId,omitempty"`
	RoleDefinition        *GovernanceRoleDefinition `json:"roleDefinition,omitempty"`
	RoleDefinitionId      *string                   `json:"roleDefinitionId,omitempty"`
	UserEligibleSettings  *[]GovernanceRuleSetting  `json:"userEligibleSettings,omitempty"`
	UserMemberSettings    *[]GovernanceRuleSetting  `json:"userMemberSettings,omitempty"`
}
