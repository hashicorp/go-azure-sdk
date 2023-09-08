package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementRoleDefinition struct {
	Description             *string           `json:"description,omitempty"`
	DisplayName             *string           `json:"displayName,omitempty"`
	Id                      *string           `json:"id,omitempty"`
	IsBuiltIn               *bool             `json:"isBuiltIn,omitempty"`
	IsBuiltInRoleDefinition *bool             `json:"isBuiltInRoleDefinition,omitempty"`
	ODataType               *string           `json:"@odata.type,omitempty"`
	Permissions             *[]RolePermission `json:"permissions,omitempty"`
	RoleAssignments         *[]RoleAssignment `json:"roleAssignments,omitempty"`
	RolePermissions         *[]RolePermission `json:"rolePermissions,omitempty"`
	RoleScopeTagIds         *[]string         `json:"roleScopeTagIds,omitempty"`
}
