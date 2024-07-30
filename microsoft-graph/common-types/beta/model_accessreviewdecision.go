package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewDecision struct {
	AccessRecommendation *string       `json:"accessRecommendation,omitempty"`
	AccessReviewId       *string       `json:"accessReviewId,omitempty"`
	AppliedBy            *UserIdentity `json:"appliedBy,omitempty"`
	AppliedDateTime      *string       `json:"appliedDateTime,omitempty"`
	ApplyResult          *string       `json:"applyResult,omitempty"`
	Id                   *string       `json:"id,omitempty"`
	Justification        *string       `json:"justification,omitempty"`
	ODataType            *string       `json:"@odata.type,omitempty"`
	ReviewResult         *string       `json:"reviewResult,omitempty"`
	ReviewedBy           *UserIdentity `json:"reviewedBy,omitempty"`
	ReviewedDateTime     *string       `json:"reviewedDateTime,omitempty"`
}
