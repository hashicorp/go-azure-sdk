package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskReportSummaryOperationPredicate struct {
	FailedTasks      *int64
	ODataType        *string
	SuccessfulTasks  *int64
	TotalTasks       *int64
	UnprocessedTasks *int64
}

func (p IdentityGovernanceTaskReportSummaryOperationPredicate) Matches(input IdentityGovernanceTaskReportSummary) bool {

	if p.FailedTasks != nil && (input.FailedTasks == nil || *p.FailedTasks != *input.FailedTasks) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SuccessfulTasks != nil && (input.SuccessfulTasks == nil || *p.SuccessfulTasks != *input.SuccessfulTasks) {
		return false
	}

	if p.TotalTasks != nil && (input.TotalTasks == nil || *p.TotalTasks != *input.TotalTasks) {
		return false
	}

	if p.UnprocessedTasks != nil && (input.UnprocessedTasks == nil || *p.UnprocessedTasks != *input.UnprocessedTasks) {
		return false
	}

	return true
}
