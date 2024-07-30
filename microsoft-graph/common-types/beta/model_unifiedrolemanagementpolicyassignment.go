package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyAssignment struct {
	Id               *string                      `json:"id,omitempty"`
	ODataType        *string                      `json:"@odata.type,omitempty"`
	Policy           *UnifiedRoleManagementPolicy `json:"policy,omitempty"`
	PolicyId         *string                      `json:"policyId,omitempty"`
	RoleDefinitionId *string                      `json:"roleDefinitionId,omitempty"`
	ScopeId          *string                      `json:"scopeId,omitempty"`
	ScopeType        *string                      `json:"scopeType,omitempty"`
}
