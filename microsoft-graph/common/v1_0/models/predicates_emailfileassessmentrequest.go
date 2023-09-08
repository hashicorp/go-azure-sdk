package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailFileAssessmentRequestOperationPredicate struct {
	ContentData     *string
	CreatedDateTime *string
	Id              *string
	ODataType       *string
	RecipientEmail  *string
}

func (p EmailFileAssessmentRequestOperationPredicate) Matches(input EmailFileAssessmentRequest) bool {

	if p.ContentData != nil && (input.ContentData == nil || *p.ContentData != *input.ContentData) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecipientEmail != nil && (input.RecipientEmail == nil || *p.RecipientEmail != *input.RecipientEmail) {
		return false
	}

	return true
}
