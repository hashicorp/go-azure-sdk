package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestRequirements struct {
	AllowCustomAssignmentSchedule *bool                          `json:"allowCustomAssignmentSchedule,omitempty"`
	IsApprovalRequiredForAdd      *bool                          `json:"isApprovalRequiredForAdd,omitempty"`
	IsApprovalRequiredForUpdate   *bool                          `json:"isApprovalRequiredForUpdate,omitempty"`
	ODataType                     *string                        `json:"@odata.type,omitempty"`
	PolicyDescription             *string                        `json:"policyDescription,omitempty"`
	PolicyDisplayName             *string                        `json:"policyDisplayName,omitempty"`
	PolicyId                      *string                        `json:"policyId,omitempty"`
	Questions                     *[]AccessPackageQuestion       `json:"questions,omitempty"`
	Schedule                      *EntitlementManagementSchedule `json:"schedule,omitempty"`
}
