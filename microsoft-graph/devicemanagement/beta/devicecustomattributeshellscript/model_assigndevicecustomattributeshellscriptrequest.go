package devicecustomattributeshellscript

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceCustomAttributeShellScriptRequest struct {
	DeviceManagementScriptAssignments      *[]beta.DeviceManagementScriptAssignment      `json:"deviceManagementScriptAssignments,omitempty"`
	DeviceManagementScriptGroupAssignments *[]beta.DeviceManagementScriptGroupAssignment `json:"deviceManagementScriptGroupAssignments,omitempty"`
}
