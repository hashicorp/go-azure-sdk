package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceUserSummaryOperationPredicate struct {
	FailedTasks     *int64
	FailedUsers     *int64
	ODataType       *string
	SuccessfulUsers *int64
	TotalTasks      *int64
	TotalUsers      *int64
}

func (p IdentityGovernanceUserSummaryOperationPredicate) Matches(input IdentityGovernanceUserSummary) bool {

	if p.FailedTasks != nil && (input.FailedTasks == nil || *p.FailedTasks != *input.FailedTasks) {
		return false
	}

	if p.FailedUsers != nil && (input.FailedUsers == nil || *p.FailedUsers != *input.FailedUsers) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SuccessfulUsers != nil && (input.SuccessfulUsers == nil || *p.SuccessfulUsers != *input.SuccessfulUsers) {
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
