package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignedRoleDetails struct {
	ODataType         *string   `json:"@odata.type,omitempty"`
	RoleAssignmentIds *[]string `json:"roleAssignmentIds,omitempty"`
	RoleDefinitionIds *[]string `json:"roleDefinitionIds,omitempty"`
}
