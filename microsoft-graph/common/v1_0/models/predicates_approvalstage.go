package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApprovalStageOperationPredicate struct {
	AssignedToMe     *bool
	DisplayName      *string
	Id               *string
	Justification    *string
	ODataType        *string
	ReviewResult     *string
	ReviewedDateTime *string
	Status           *string
}

func (p ApprovalStageOperationPredicate) Matches(input ApprovalStage) bool {

	if p.AssignedToMe != nil && (input.AssignedToMe == nil || *p.AssignedToMe != *input.AssignedToMe) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Justification != nil && (input.Justification == nil || *p.Justification != *input.Justification) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReviewResult != nil && (input.ReviewResult == nil || *p.ReviewResult != *input.ReviewResult) {
		return false
	}

	if p.ReviewedDateTime != nil && (input.ReviewedDateTime == nil || *p.ReviewedDateTime != *input.ReviewedDateTime) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
