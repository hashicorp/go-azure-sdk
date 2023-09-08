package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllLicensedUsersAssignmentTarget struct {
	DeviceAndAppManagementAssignmentFilterId   *string                                                                     `json:"deviceAndAppManagementAssignmentFilterId,omitempty"`
	DeviceAndAppManagementAssignmentFilterType *AllLicensedUsersAssignmentTargetDeviceAndAppManagementAssignmentFilterType `json:"deviceAndAppManagementAssignmentFilterType,omitempty"`
	ODataType                                  *string                                                                     `json:"@odata.type,omitempty"`
}
