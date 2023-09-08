package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembershipOutlierInsightOperationPredicate struct {
	ContainerId            *string
	Id                     *string
	InsightCreatedDateTime *string
	MemberId               *string
	ODataType              *string
}

func (p MembershipOutlierInsightOperationPredicate) Matches(input MembershipOutlierInsight) bool {

	if p.ContainerId != nil && (input.ContainerId == nil || *p.ContainerId != *input.ContainerId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InsightCreatedDateTime != nil && (input.InsightCreatedDateTime == nil || *p.InsightCreatedDateTime != *input.InsightCreatedDateTime) {
		return false
	}

	if p.MemberId != nil && (input.MemberId == nil || *p.MemberId != *input.MemberId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
