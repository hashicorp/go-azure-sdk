package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewOperationPredicate struct {
	BusinessFlowTemplateId *string
	Description            *string
	DisplayName            *string
	EndDateTime            *string
	Id                     *string
	ODataType              *string
	ReviewerType           *string
	StartDateTime          *string
	Status                 *string
}

func (p AccessReviewOperationPredicate) Matches(input AccessReview) bool {

	if p.BusinessFlowTemplateId != nil && (input.BusinessFlowTemplateId == nil || *p.BusinessFlowTemplateId != *input.BusinessFlowTemplateId) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReviewerType != nil && (input.ReviewerType == nil || *p.ReviewerType != *input.ReviewerType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
