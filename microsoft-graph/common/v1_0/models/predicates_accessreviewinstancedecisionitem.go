package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItemOperationPredicate struct {
	AccessReviewId   *string
	AppliedDateTime  *string
	ApplyResult      *string
	Decision         *string
	Id               *string
	Justification    *string
	ODataType        *string
	PrincipalLink    *string
	Recommendation   *string
	ResourceLink     *string
	ReviewedDateTime *string
}

func (p AccessReviewInstanceDecisionItemOperationPredicate) Matches(input AccessReviewInstanceDecisionItem) bool {

	if p.AccessReviewId != nil && (input.AccessReviewId == nil || *p.AccessReviewId != *input.AccessReviewId) {
		return false
	}

	if p.AppliedDateTime != nil && (input.AppliedDateTime == nil || *p.AppliedDateTime != *input.AppliedDateTime) {
		return false
	}

	if p.ApplyResult != nil && (input.ApplyResult == nil || *p.ApplyResult != *input.ApplyResult) {
		return false
	}

	if p.Decision != nil && (input.Decision == nil || *p.Decision != *input.Decision) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Justification != nil && (input.Justification == nil || *p.Justification != *input.Justification) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrincipalLink != nil && (input.PrincipalLink == nil || *p.PrincipalLink != *input.PrincipalLink) {
		return false
	}

	if p.Recommendation != nil && (input.Recommendation == nil || *p.Recommendation != *input.Recommendation) {
		return false
	}

	if p.ResourceLink != nil && (input.ResourceLink == nil || *p.ResourceLink != *input.ResourceLink) {
		return false
	}

	if p.ReviewedDateTime != nil && (input.ReviewedDateTime == nil || *p.ReviewedDateTime != *input.ReviewedDateTime) {
		return false
	}

	return true
}
