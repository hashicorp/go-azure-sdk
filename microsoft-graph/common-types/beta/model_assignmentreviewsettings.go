package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentReviewSettings struct {
	AccessReviewTimeoutBehavior     *AssignmentReviewSettingsAccessReviewTimeoutBehavior `json:"accessReviewTimeoutBehavior,omitempty"`
	DurationInDays                  *int64                                               `json:"durationInDays,omitempty"`
	IsAccessRecommendationEnabled   *bool                                                `json:"isAccessRecommendationEnabled,omitempty"`
	IsApprovalJustificationRequired *bool                                                `json:"isApprovalJustificationRequired,omitempty"`
	IsEnabled                       *bool                                                `json:"isEnabled,omitempty"`
	ODataType                       *string                                              `json:"@odata.type,omitempty"`
	RecurrenceType                  *string                                              `json:"recurrenceType,omitempty"`
	ReviewerType                    *string                                              `json:"reviewerType,omitempty"`
	Reviewers                       *[]UserSet                                           `json:"reviewers,omitempty"`
	StartDateTime                   *string                                              `json:"startDateTime,omitempty"`
}
