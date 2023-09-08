package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalStage struct {
	ApprovalStageTimeOutInDays      *int64     `json:"approvalStageTimeOutInDays,omitempty"`
	EscalationApprovers             *[]UserSet `json:"escalationApprovers,omitempty"`
	EscalationTimeInMinutes         *int64     `json:"escalationTimeInMinutes,omitempty"`
	IsApproverJustificationRequired *bool      `json:"isApproverJustificationRequired,omitempty"`
	IsEscalationEnabled             *bool      `json:"isEscalationEnabled,omitempty"`
	ODataType                       *string    `json:"@odata.type,omitempty"`
	PrimaryApprovers                *[]UserSet `json:"primaryApprovers,omitempty"`
}
