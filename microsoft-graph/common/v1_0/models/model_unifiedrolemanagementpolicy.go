package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicy struct {
	Description           *string                            `json:"description,omitempty"`
	DisplayName           *string                            `json:"displayName,omitempty"`
	EffectiveRules        *[]UnifiedRoleManagementPolicyRule `json:"effectiveRules,omitempty"`
	Id                    *string                            `json:"id,omitempty"`
	IsOrganizationDefault *bool                              `json:"isOrganizationDefault,omitempty"`
	LastModifiedBy        *Identity                          `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime  *string                            `json:"lastModifiedDateTime,omitempty"`
	ODataType             *string                            `json:"@odata.type,omitempty"`
	Rules                 *[]UnifiedRoleManagementPolicyRule `json:"rules,omitempty"`
	ScopeId               *string                            `json:"scopeId,omitempty"`
	ScopeType             *string                            `json:"scopeType,omitempty"`
}
