package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalStageOperationPredicate struct {
	ApprovalStageTimeOutInDays      *int64
	EscalationTimeInMinutes         *int64
	IsApproverJustificationRequired *bool
	IsEscalationEnabled             *bool
	ODataType                       *string
}

func (p ApprovalStageOperationPredicate) Matches(input ApprovalStage) bool {

	if p.ApprovalStageTimeOutInDays != nil && (input.ApprovalStageTimeOutInDays == nil || *p.ApprovalStageTimeOutInDays != *input.ApprovalStageTimeOutInDays) {
		return false
	}

	if p.EscalationTimeInMinutes != nil && (input.EscalationTimeInMinutes == nil || *p.EscalationTimeInMinutes != *input.EscalationTimeInMinutes) {
		return false
	}

	if p.IsApproverJustificationRequired != nil && (input.IsApproverJustificationRequired == nil || *p.IsApproverJustificationRequired != *input.IsApproverJustificationRequired) {
		return false
	}

	if p.IsEscalationEnabled != nil && (input.IsEscalationEnabled == nil || *p.IsEscalationEnabled != *input.IsEscalationEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
