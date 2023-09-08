package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecisionOperationPredicate struct {
	AccessRecommendation *string
	AccessReviewId       *string
	AppliedDateTime      *string
	ApplyResult          *string
	Id                   *string
	Justification        *string
	ODataType            *string
	ReviewResult         *string
	ReviewedDateTime     *string
}

func (p AccessReviewDecisionOperationPredicate) Matches(input AccessReviewDecision) bool {

	if p.AccessRecommendation != nil && (input.AccessRecommendation == nil || *p.AccessRecommendation != *input.AccessRecommendation) {
		return false
	}

	if p.AccessReviewId != nil && (input.AccessReviewId == nil || *p.AccessReviewId != *input.AccessReviewId) {
		return false
	}

	if p.AppliedDateTime != nil && (input.AppliedDateTime == nil || *p.AppliedDateTime != *input.AppliedDateTime) {
		return false
	}

	if p.ApplyResult != nil && (input.ApplyResult == nil || *p.ApplyResult != *input.ApplyResult) {
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

	if p.ReviewResult != nil && (input.ReviewResult == nil || *p.ReviewResult != *input.ReviewResult) {
		return false
	}

	if p.ReviewedDateTime != nil && (input.ReviewedDateTime == nil || *p.ReviewedDateTime != *input.ReviewedDateTime) {
		return false
	}

	return true
}
