package deviceconfiguration

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementDeviceConfigurationsRequest struct {
	Assignments                         *[]DeviceConfigurationAssignment      `json:"assignments,omitempty"`
	DeviceConfigurationGroupAssignments *[]DeviceConfigurationGroupAssignment `json:"deviceConfigurationGroupAssignments,omitempty"`
}
