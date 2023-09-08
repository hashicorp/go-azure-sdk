package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskReportOperationPredicate struct {
	CompletedDateTime     *string
	FailedUsersCount      *int64
	Id                    *string
	LastUpdatedDateTime   *string
	ODataType             *string
	RunId                 *string
	StartedDateTime       *string
	SuccessfulUsersCount  *int64
	TotalUsersCount       *int64
	UnprocessedUsersCount *int64
}

func (p IdentityGovernanceTaskReportOperationPredicate) Matches(input IdentityGovernanceTaskReport) bool {

	if p.CompletedDateTime != nil && (input.CompletedDateTime == nil || *p.CompletedDateTime != *input.CompletedDateTime) {
		return false
	}

	if p.FailedUsersCount != nil && (input.FailedUsersCount == nil || *p.FailedUsersCount != *input.FailedUsersCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RunId != nil && (input.RunId == nil || *p.RunId != *input.RunId) {
		return false
	}

	if p.StartedDateTime != nil && (input.StartedDateTime == nil || *p.StartedDateTime != *input.StartedDateTime) {
		return false
	}

	if p.SuccessfulUsersCount != nil && (input.SuccessfulUsersCount == nil || *p.SuccessfulUsersCount != *input.SuccessfulUsersCount) {
		return false
	}

	if p.TotalUsersCount != nil && (input.TotalUsersCount == nil || *p.TotalUsersCount != *input.TotalUsersCount) {
		return false
	}

	if p.UnprocessedUsersCount != nil && (input.UnprocessedUsersCount == nil || *p.UnprocessedUsersCount != *input.UnprocessedUsersCount) {
		return false
	}

	return true
}
