package assignmentfilter

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentFilter struct {
	AssignmentFilterManagementType *DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType `json:"assignmentFilterManagementType,omitempty"`
	CreatedDateTime                *string                                                               `json:"createdDateTime,omitempty"`
	Description                    *string                                                               `json:"description,omitempty"`
	DisplayName                    *string                                                               `json:"displayName,omitempty"`
	Id                             *string                                                               `json:"id,omitempty"`
	LastModifiedDateTime           *string                                                               `json:"lastModifiedDateTime,omitempty"`
	ODataType                      *string                                                               `json:"@odata.type,omitempty"`
	Payloads                       *[]PayloadByFilter                                                    `json:"payloads,omitempty"`
	Platform                       *DeviceAndAppManagementAssignmentFilterPlatform                       `json:"platform,omitempty"`
	RoleScopeTags                  *[]string                                                             `json:"roleScopeTags,omitempty"`
	Rule                           *string                                                               `json:"rule,omitempty"`
}
