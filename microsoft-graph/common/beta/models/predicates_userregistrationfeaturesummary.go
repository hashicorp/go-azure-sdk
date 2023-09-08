package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureSummaryOperationPredicate struct {
	ODataType      *string
	TotalUserCount *int64
}

func (p UserRegistrationFeatureSummaryOperationPredicate) Matches(input UserRegistrationFeatureSummary) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TotalUserCount != nil && (input.TotalUserCount == nil || *p.TotalUserCount != *input.TotalUserCount) {
		return false
	}

	return true
}
