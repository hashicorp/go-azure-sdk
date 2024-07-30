package devicecompliancescript

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementDeviceComplianceScriptRequest struct {
	DeviceHealthScriptAssignments *[]DeviceHealthScriptAssignment `json:"deviceHealthScriptAssignments,omitempty"`
}
