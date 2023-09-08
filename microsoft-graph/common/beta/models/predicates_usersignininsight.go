package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSignInInsightOperationPredicate struct {
	Id                     *string
	InsightCreatedDateTime *string
	LastSignInDateTime     *string
	ODataType              *string
}

func (p UserSignInInsightOperationPredicate) Matches(input UserSignInInsight) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InsightCreatedDateTime != nil && (input.InsightCreatedDateTime == nil || *p.InsightCreatedDateTime != *input.InsightCreatedDateTime) {
		return false
	}

	if p.LastSignInDateTime != nil && (input.LastSignInDateTime == nil || *p.LastSignInDateTime != *input.LastSignInDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
