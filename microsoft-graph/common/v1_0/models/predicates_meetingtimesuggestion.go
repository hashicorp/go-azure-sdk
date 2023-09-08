package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingTimeSuggestionOperationPredicate struct {
	ODataType        *string
	Order            *int64
	SuggestionReason *string
}

func (p MeetingTimeSuggestionOperationPredicate) Matches(input MeetingTimeSuggestion) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Order != nil && (input.Order == nil || *p.Order != *input.Order) {
		return false
	}

	if p.SuggestionReason != nil && (input.SuggestionReason == nil || *p.SuggestionReason != *input.SuggestionReason) {
		return false
	}

	return true
}
