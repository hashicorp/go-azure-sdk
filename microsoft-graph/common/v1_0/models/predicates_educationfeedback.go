package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationFeedbackOperationPredicate struct {
	FeedbackDateTime *string
	ODataType        *string
}

func (p EducationFeedbackOperationPredicate) Matches(input EducationFeedback) bool {

	if p.FeedbackDateTime != nil && (input.FeedbackDateTime == nil || *p.FeedbackDateTime != *input.FeedbackDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
