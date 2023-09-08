package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstance struct {
	ContactedReviewers *[]AccessReviewReviewer             `json:"contactedReviewers,omitempty"`
	Decisions          *[]AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`
	Definition         *AccessReviewScheduleDefinition     `json:"definition,omitempty"`
	EndDateTime        *string                             `json:"endDateTime,omitempty"`
	Errors             *[]AccessReviewError                `json:"errors,omitempty"`
	FallbackReviewers  *[]AccessReviewReviewerScope        `json:"fallbackReviewers,omitempty"`
	Id                 *string                             `json:"id,omitempty"`
	ODataType          *string                             `json:"@odata.type,omitempty"`
	Reviewers          *[]AccessReviewReviewerScope        `json:"reviewers,omitempty"`
	Scope              *AccessReviewScope                  `json:"scope,omitempty"`
	Stages             *[]AccessReviewStage                `json:"stages,omitempty"`
	StartDateTime      *string                             `json:"startDateTime,omitempty"`
	Status             *string                             `json:"status,omitempty"`
}
