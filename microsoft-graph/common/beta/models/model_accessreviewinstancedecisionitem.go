package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItem struct {
	AccessReviewId              *string                                   `json:"accessReviewId,omitempty"`
	AppliedBy                   *UserIdentity                             `json:"appliedBy,omitempty"`
	AppliedDateTime             *string                                   `json:"appliedDateTime,omitempty"`
	ApplyResult                 *string                                   `json:"applyResult,omitempty"`
	Decision                    *string                                   `json:"decision,omitempty"`
	Id                          *string                                   `json:"id,omitempty"`
	Insights                    *[]GovernanceInsight                      `json:"insights,omitempty"`
	Instance                    *AccessReviewInstance                     `json:"instance,omitempty"`
	Justification               *string                                   `json:"justification,omitempty"`
	ODataType                   *string                                   `json:"@odata.type,omitempty"`
	Principal                   *Identity                                 `json:"principal,omitempty"`
	PrincipalLink               *string                                   `json:"principalLink,omitempty"`
	PrincipalResourceMembership *DecisionItemPrincipalResourceMembership  `json:"principalResourceMembership,omitempty"`
	Recommendation              *string                                   `json:"recommendation,omitempty"`
	Resource                    *AccessReviewInstanceDecisionItemResource `json:"resource,omitempty"`
	ResourceLink                *string                                   `json:"resourceLink,omitempty"`
	ReviewedBy                  *UserIdentity                             `json:"reviewedBy,omitempty"`
	ReviewedDateTime            *string                                   `json:"reviewedDateTime,omitempty"`
	Target                      *AccessReviewInstanceDecisionItemTarget   `json:"target,omitempty"`
}
