package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RbacApplicationMultiple struct {
	Id                 *string                          `json:"id,omitempty"`
	ODataType          *string                          `json:"@odata.type,omitempty"`
	ResourceNamespaces *[]UnifiedRbacResourceNamespace  `json:"resourceNamespaces,omitempty"`
	RoleAssignments    *[]UnifiedRoleAssignmentMultiple `json:"roleAssignments,omitempty"`
	RoleDefinitions    *[]UnifiedRoleDefinition         `json:"roleDefinitions,omitempty"`
}
