package devicecustomattributeshellscript

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementDeviceCustomAttributeShellScriptRequest struct {
	DeviceManagementScriptAssignments      *[]DeviceManagementScriptAssignment      `json:"deviceManagementScriptAssignments,omitempty"`
	DeviceManagementScriptGroupAssignments *[]DeviceManagementScriptGroupAssignment `json:"deviceManagementScriptGroupAssignments,omitempty"`
}
