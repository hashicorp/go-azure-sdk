package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementRoleAssignment struct {
	Description    *string                                        `json:"description,omitempty"`
	DisplayName    *string                                        `json:"displayName,omitempty"`
	Id             *string                                        `json:"id,omitempty"`
	Members        *[]string                                      `json:"members,omitempty"`
	ODataType      *string                                        `json:"@odata.type,omitempty"`
	ResourceScopes *[]string                                      `json:"resourceScopes,omitempty"`
	RoleDefinition *RoleDefinition                                `json:"roleDefinition,omitempty"`
	RoleScopeTags  *[]RoleScopeTag                                `json:"roleScopeTags,omitempty"`
	ScopeMembers   *[]string                                      `json:"scopeMembers,omitempty"`
	ScopeType      *DeviceAndAppManagementRoleAssignmentScopeType `json:"scopeType,omitempty"`
}
