package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunDetailsOperationPredicate struct {
	FailureReason   *string
	LastRunDateTime *string
	ODataType       *string
}

func (p SecurityRunDetailsOperationPredicate) Matches(input SecurityRunDetails) bool {

	if p.FailureReason != nil && (input.FailureReason == nil || *p.FailureReason != *input.FailureReason) {
		return false
	}

	if p.LastRunDateTime != nil && (input.LastRunDateTime == nil || *p.LastRunDateTime != *input.LastRunDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
