package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReview struct {
	BusinessFlowTemplateId *string                 `json:"businessFlowTemplateId,omitempty"`
	CreatedBy              *UserIdentity           `json:"createdBy,omitempty"`
	Decisions              *[]AccessReviewDecision `json:"decisions,omitempty"`
	Description            *string                 `json:"description,omitempty"`
	DisplayName            *string                 `json:"displayName,omitempty"`
	EndDateTime            *string                 `json:"endDateTime,omitempty"`
	Id                     *string                 `json:"id,omitempty"`
	Instances              *[]AccessReview         `json:"instances,omitempty"`
	MyDecisions            *[]AccessReviewDecision `json:"myDecisions,omitempty"`
	ODataType              *string                 `json:"@odata.type,omitempty"`
	ReviewedEntity         *Identity               `json:"reviewedEntity,omitempty"`
	ReviewerType           *string                 `json:"reviewerType,omitempty"`
	Reviewers              *[]AccessReviewReviewer `json:"reviewers,omitempty"`
	Settings               *AccessReviewSettings   `json:"settings,omitempty"`
	StartDateTime          *string                 `json:"startDateTime,omitempty"`
	Status                 *string                 `json:"status,omitempty"`
}
