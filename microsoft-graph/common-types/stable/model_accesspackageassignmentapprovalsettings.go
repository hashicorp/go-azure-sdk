package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentApprovalSettings struct {
	IsApprovalRequiredForAdd    *bool                         `json:"isApprovalRequiredForAdd,omitempty"`
	IsApprovalRequiredForUpdate *bool                         `json:"isApprovalRequiredForUpdate,omitempty"`
	ODataType                   *string                       `json:"@odata.type,omitempty"`
	Stages                      *[]AccessPackageApprovalStage `json:"stages,omitempty"`
}
