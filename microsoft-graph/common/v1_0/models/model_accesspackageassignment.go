package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignment struct {
	AccessPackage                   *AccessPackage                    `json:"accessPackage,omitempty"`
	AssignmentPolicy                *AccessPackageAssignmentPolicy    `json:"assignmentPolicy,omitempty"`
	CustomExtensionCalloutInstances *[]CustomExtensionCalloutInstance `json:"customExtensionCalloutInstances,omitempty"`
	ExpiredDateTime                 *string                           `json:"expiredDateTime,omitempty"`
	Id                              *string                           `json:"id,omitempty"`
	ODataType                       *string                           `json:"@odata.type,omitempty"`
	Schedule                        *EntitlementManagementSchedule    `json:"schedule,omitempty"`
	State                           *AccessPackageAssignmentState     `json:"state,omitempty"`
	Status                          *string                           `json:"status,omitempty"`
	Target                          *AccessPackageSubject             `json:"target,omitempty"`
}
