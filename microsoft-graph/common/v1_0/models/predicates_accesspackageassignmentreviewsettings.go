package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentReviewSettingsOperationPredicate struct {
	IsEnabled                       *bool
	IsRecommendationEnabled         *bool
	IsReviewerJustificationRequired *bool
	IsSelfReview                    *bool
	ODataType                       *string
}

func (p AccessPackageAssignmentReviewSettingsOperationPredicate) Matches(input AccessPackageAssignmentReviewSettings) bool {

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.IsRecommendationEnabled != nil && (input.IsRecommendationEnabled == nil || *p.IsRecommendationEnabled != *input.IsRecommendationEnabled) {
		return false
	}

	if p.IsReviewerJustificationRequired != nil && (input.IsReviewerJustificationRequired == nil || *p.IsReviewerJustificationRequired != *input.IsReviewerJustificationRequired) {
		return false
	}

	if p.IsSelfReview != nil && (input.IsSelfReview == nil || *p.IsSelfReview != *input.IsSelfReview) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
