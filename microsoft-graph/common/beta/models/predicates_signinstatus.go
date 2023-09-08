package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInStatusOperationPredicate struct {
	AdditionalDetails *string
	ErrorCode         *int64
	FailureReason     *string
	ODataType         *string
}

func (p SignInStatusOperationPredicate) Matches(input SignInStatus) bool {

	if p.AdditionalDetails != nil && (input.AdditionalDetails == nil || *p.AdditionalDetails != *input.AdditionalDetails) {
		return false
	}

	if p.ErrorCode != nil && (input.ErrorCode == nil || *p.ErrorCode != *input.ErrorCode) {
		return false
	}

	if p.FailureReason != nil && (input.FailureReason == nil || *p.FailureReason != *input.FailureReason) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
