package windowsdriverupdateprofile

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignDeviceManagementWindowsDriverUpdateProfileRequest struct {
	Assignments *[]WindowsDriverUpdateProfileAssignment `json:"assignments,omitempty"`
}
