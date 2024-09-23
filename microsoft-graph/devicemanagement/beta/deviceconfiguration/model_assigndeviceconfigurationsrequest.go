package deviceconfiguration

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceConfigurationsRequest struct {
	Assignments                         *[]beta.DeviceConfigurationAssignment      `json:"assignments,omitempty"`
	DeviceConfigurationGroupAssignments *[]beta.DeviceConfigurationGroupAssignment `json:"deviceConfigurationGroupAssignments,omitempty"`
}
