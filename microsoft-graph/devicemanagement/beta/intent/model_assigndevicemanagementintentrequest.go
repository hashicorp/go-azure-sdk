package intent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementIntentRequest struct {
	Assignments *[]DeviceManagementIntentAssignment `json:"assignments,omitempty"`
}
