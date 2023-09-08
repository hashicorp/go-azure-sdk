package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRunSummaryOperationPredicate struct {
	FailedRuns     *int64
	FailedTasks    *int64
	ODataType      *string
	SuccessfulRuns *int64
	TotalRuns      *int64
	TotalTasks     *int64
	TotalUsers     *int64
}

func (p IdentityGovernanceRunSummaryOperationPredicate) Matches(input IdentityGovernanceRunSummary) bool {

	if p.FailedRuns != nil && (input.FailedRuns == nil || *p.FailedRuns != *input.FailedRuns) {
		return false
	}

	if p.FailedTasks != nil && (input.FailedTasks == nil || *p.FailedTasks != *input.FailedTasks) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SuccessfulRuns != nil && (input.SuccessfulRuns == nil || *p.SuccessfulRuns != *input.SuccessfulRuns) {
		return false
	}

	if p.TotalRuns != nil && (input.TotalRuns == nil || *p.TotalRuns != *input.TotalRuns) {
		return false
	}

	if p.TotalTasks != nil && (input.TotalTasks == nil || *p.TotalTasks != *input.TotalTasks) {
		return false
	}

	if p.TotalUsers != nil && (input.TotalUsers == nil || *p.TotalUsers != *input.TotalUsers) {
		return false
	}

	return true
}
