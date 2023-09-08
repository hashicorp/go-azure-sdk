package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentReviewSettings struct {
	ExpirationBehavior              *AccessPackageAssignmentReviewSettingsExpirationBehavior `json:"expirationBehavior,omitempty"`
	FallbackReviewers               *[]SubjectSet                                            `json:"fallbackReviewers,omitempty"`
	IsEnabled                       *bool                                                    `json:"isEnabled,omitempty"`
	IsRecommendationEnabled         *bool                                                    `json:"isRecommendationEnabled,omitempty"`
	IsReviewerJustificationRequired *bool                                                    `json:"isReviewerJustificationRequired,omitempty"`
	IsSelfReview                    *bool                                                    `json:"isSelfReview,omitempty"`
	ODataType                       *string                                                  `json:"@odata.type,omitempty"`
	PrimaryReviewers                *[]SubjectSet                                            `json:"primaryReviewers,omitempty"`
	Schedule                        *EntitlementManagementSchedule                           `json:"schedule,omitempty"`
}
