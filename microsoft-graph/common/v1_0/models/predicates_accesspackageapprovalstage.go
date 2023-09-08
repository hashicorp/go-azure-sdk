package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageApprovalStageOperationPredicate struct {
	DurationBeforeAutomaticDenial   *string
	DurationBeforeEscalation        *string
	IsApproverJustificationRequired *bool
	IsEscalationEnabled             *bool
	ODataType                       *string
}

func (p AccessPackageApprovalStageOperationPredicate) Matches(input AccessPackageApprovalStage) bool {

	if p.DurationBeforeAutomaticDenial != nil && (input.DurationBeforeAutomaticDenial == nil || *p.DurationBeforeAutomaticDenial != *input.DurationBeforeAutomaticDenial) {
		return false
	}

	if p.DurationBeforeEscalation != nil && (input.DurationBeforeEscalation == nil || *p.DurationBeforeEscalation != *input.DurationBeforeEscalation) {
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
