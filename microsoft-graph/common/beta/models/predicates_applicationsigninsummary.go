package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationSignInSummaryOperationPredicate struct {
	AppDisplayName        *string
	FailedSignInCount     *int64
	Id                    *string
	ODataType             *string
	SuccessfulSignInCount *int64
}

func (p ApplicationSignInSummaryOperationPredicate) Matches(input ApplicationSignInSummary) bool {

	if p.AppDisplayName != nil && (input.AppDisplayName == nil || *p.AppDisplayName != *input.AppDisplayName) {
		return false
	}

	if p.FailedSignInCount != nil && (input.FailedSignInCount == nil || *p.FailedSignInCount != *input.FailedSignInCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SuccessfulSignInCount != nil && (input.SuccessfulSignInCount == nil || *p.SuccessfulSignInCount != *input.SuccessfulSignInCount) {
		return false
	}

	return true
}
