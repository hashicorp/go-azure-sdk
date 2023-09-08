package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CredentialUsageSummaryOperationPredicate struct {
	FailureActivityCount    *int64
	Id                      *string
	ODataType               *string
	SuccessfulActivityCount *int64
}

func (p CredentialUsageSummaryOperationPredicate) Matches(input CredentialUsageSummary) bool {

	if p.FailureActivityCount != nil && (input.FailureActivityCount == nil || *p.FailureActivityCount != *input.FailureActivityCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SuccessfulActivityCount != nil && (input.SuccessfulActivityCount == nil || *p.SuccessfulActivityCount != *input.SuccessfulActivityCount) {
		return false
	}

	return true
}
