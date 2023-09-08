package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityGovernanceTaskProcessingResultOperationPredicate struct {
	CompletedDateTime *string
	CreatedDateTime   *string
	FailureReason     *string
	Id                *string
	ODataType         *string
	StartedDateTime   *string
}

func (p IdentityGovernanceTaskProcessingResultOperationPredicate) Matches(input IdentityGovernanceTaskProcessingResult) bool {

	if p.CompletedDateTime != nil && (input.CompletedDateTime == nil || *p.CompletedDateTime != *input.CompletedDateTime) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.FailureReason != nil && (input.FailureReason == nil || *p.FailureReason != *input.FailureReason) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartedDateTime != nil && (input.StartedDateTime == nil || *p.StartedDateTime != *input.StartedDateTime) {
		return false
	}

	return true
}
