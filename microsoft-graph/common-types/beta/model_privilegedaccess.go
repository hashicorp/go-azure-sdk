package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccess struct {
	DisplayName            *string                            `json:"displayName,omitempty"`
	Id                     *string                            `json:"id,omitempty"`
	ODataType              *string                            `json:"@odata.type,omitempty"`
	Resources              *[]GovernanceResource              `json:"resources,omitempty"`
	RoleAssignmentRequests *[]GovernanceRoleAssignmentRequest `json:"roleAssignmentRequests,omitempty"`
	RoleAssignments        *[]GovernanceRoleAssignment        `json:"roleAssignments,omitempty"`
	RoleDefinitions        *[]GovernanceRoleDefinition        `json:"roleDefinitions,omitempty"`
	RoleSettings           *[]GovernanceRoleSetting           `json:"roleSettings,omitempty"`
}
