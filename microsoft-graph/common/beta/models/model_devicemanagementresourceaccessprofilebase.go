package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementResourceAccessProfileBase struct {
	Assignments          *[]DeviceManagementResourceAccessProfileAssignment `json:"assignments,omitempty"`
	CreationDateTime     *string                                            `json:"creationDateTime,omitempty"`
	Description          *string                                            `json:"description,omitempty"`
	DisplayName          *string                                            `json:"displayName,omitempty"`
	Id                   *string                                            `json:"id,omitempty"`
	LastModifiedDateTime *string                                            `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                            `json:"@odata.type,omitempty"`
	RoleScopeTagIds      *[]string                                          `json:"roleScopeTagIds,omitempty"`
	Version              *int64                                             `json:"version,omitempty"`
}
