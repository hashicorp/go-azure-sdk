package devicemanagement

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementAssignmentFiltersStatusDetailRequest struct {
	AssignmentFilterIds *[]string `json:"assignmentFilterIds,omitempty"`
	ManagedDeviceId     *string   `json:"managedDeviceId,omitempty"`
	PayloadId           *string   `json:"payloadId,omitempty"`
	Skip                *int64    `json:"skip,omitempty"`
	Top                 *int64    `json:"top,omitempty"`
	UserId              *string   `json:"userId,omitempty"`
}
