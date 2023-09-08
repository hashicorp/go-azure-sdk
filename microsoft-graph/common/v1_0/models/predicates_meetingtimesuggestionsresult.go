package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingTimeSuggestionsResultOperationPredicate struct {
	EmptySuggestionsReason *string
	ODataType              *string
}

func (p MeetingTimeSuggestionsResultOperationPredicate) Matches(input MeetingTimeSuggestionsResult) bool {

	if p.EmptySuggestionsReason != nil && (input.EmptySuggestionsReason == nil || *p.EmptySuggestionsReason != *input.EmptySuggestionsReason) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
