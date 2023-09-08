package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentReviewSettingsOperationPredicate struct {
	DurationInDays                  *int64
	IsAccessRecommendationEnabled   *bool
	IsApprovalJustificationRequired *bool
	IsEnabled                       *bool
	ODataType                       *string
	RecurrenceType                  *string
	ReviewerType                    *string
	StartDateTime                   *string
}

func (p AssignmentReviewSettingsOperationPredicate) Matches(input AssignmentReviewSettings) bool {

	if p.DurationInDays != nil && (input.DurationInDays == nil || *p.DurationInDays != *input.DurationInDays) {
		return false
	}

	if p.IsAccessRecommendationEnabled != nil && (input.IsAccessRecommendationEnabled == nil || *p.IsAccessRecommendationEnabled != *input.IsAccessRecommendationEnabled) {
		return false
	}

	if p.IsApprovalJustificationRequired != nil && (input.IsApprovalJustificationRequired == nil || *p.IsApprovalJustificationRequired != *input.IsApprovalJustificationRequired) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecurrenceType != nil && (input.RecurrenceType == nil || *p.RecurrenceType != *input.RecurrenceType) {
		return false
	}

	if p.ReviewerType != nil && (input.ReviewerType == nil || *p.ReviewerType != *input.ReviewerType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	return true
}
