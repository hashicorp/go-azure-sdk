package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalSettings struct {
	ApprovalMode                     *string          `json:"approvalMode,omitempty"`
	ApprovalStages                   *[]ApprovalStage `json:"approvalStages,omitempty"`
	IsApprovalRequired               *bool            `json:"isApprovalRequired,omitempty"`
	IsApprovalRequiredForExtension   *bool            `json:"isApprovalRequiredForExtension,omitempty"`
	IsRequestorJustificationRequired *bool            `json:"isRequestorJustificationRequired,omitempty"`
	ODataType                        *string          `json:"@odata.type,omitempty"`
}
