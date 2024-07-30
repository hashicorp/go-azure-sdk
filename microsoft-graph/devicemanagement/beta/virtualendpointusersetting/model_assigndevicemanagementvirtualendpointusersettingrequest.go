package virtualendpointusersetting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementVirtualEndpointUserSettingRequest struct {
	Assignments *[]CloudPCUserSettingAssignment `json:"assignments,omitempty"`
}
