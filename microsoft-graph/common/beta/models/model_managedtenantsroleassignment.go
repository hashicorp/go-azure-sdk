package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsRoleAssignment struct {
	AssignmentType *ManagedTenantsRoleAssignmentAssignmentType `json:"assignmentType,omitempty"`
	ODataType      *string                                     `json:"@odata.type,omitempty"`
	Roles          *[]ManagedTenantsRoleDefinition             `json:"roles,omitempty"`
}
