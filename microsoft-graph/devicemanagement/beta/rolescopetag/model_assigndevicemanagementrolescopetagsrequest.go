package rolescopetag

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementRoleScopeTagsRequest struct {
	Assignments *[]RoleScopeTagAutoAssignment `json:"assignments,omitempty"`
}
