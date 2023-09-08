package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceRunOperationPredicate struct {
	CompletedDateTime          *string
	FailedTasksCount           *int64
	FailedUsersCount           *int64
	Id                         *string
	LastUpdatedDateTime        *string
	ODataType                  *string
	ScheduledDateTime          *string
	StartedDateTime            *string
	SuccessfulUsersCount       *int64
	TotalTasksCount            *int64
	TotalUnprocessedTasksCount *int64
	TotalUsersCount            *int64
}

func (p IdentityGovernanceRunOperationPredicate) Matches(input IdentityGovernanceRun) bool {

	if p.CompletedDateTime != nil && (input.CompletedDateTime == nil || *p.CompletedDateTime != *input.CompletedDateTime) {
		return false
	}

	if p.FailedTasksCount != nil && (input.FailedTasksCount == nil || *p.FailedTasksCount != *input.FailedTasksCount) {
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

	if p.ScheduledDateTime != nil && (input.ScheduledDateTime == nil || *p.ScheduledDateTime != *input.ScheduledDateTime) {
		return false
	}

	if p.StartedDateTime != nil && (input.StartedDateTime == nil || *p.StartedDateTime != *input.StartedDateTime) {
		return false
	}

	if p.SuccessfulUsersCount != nil && (input.SuccessfulUsersCount == nil || *p.SuccessfulUsersCount != *input.SuccessfulUsersCount) {
		return false
	}

	if p.TotalTasksCount != nil && (input.TotalTasksCount == nil || *p.TotalTasksCount != *input.TotalTasksCount) {
		return false
	}

	if p.TotalUnprocessedTasksCount != nil && (input.TotalUnprocessedTasksCount == nil || *p.TotalUnprocessedTasksCount != *input.TotalUnprocessedTasksCount) {
		return false
	}

	if p.TotalUsersCount != nil && (input.TotalUsersCount == nil || *p.TotalUsersCount != *input.TotalUsersCount) {
		return false
	}

	return true
}
