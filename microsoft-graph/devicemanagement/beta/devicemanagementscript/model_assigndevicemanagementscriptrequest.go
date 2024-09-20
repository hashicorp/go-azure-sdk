package devicemanagementscript

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementScriptRequest struct {
	DeviceManagementScriptAssignments      *[]beta.DeviceManagementScriptAssignment      `json:"deviceManagementScriptAssignments,omitempty"`
	DeviceManagementScriptGroupAssignments *[]beta.DeviceManagementScriptGroupAssignment `json:"deviceManagementScriptGroupAssignments,omitempty"`
}
