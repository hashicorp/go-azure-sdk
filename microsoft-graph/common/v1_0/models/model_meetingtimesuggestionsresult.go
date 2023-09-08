package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingTimeSuggestionsResult struct {
	EmptySuggestionsReason *string                  `json:"emptySuggestionsReason,omitempty"`
	MeetingTimeSuggestions *[]MeetingTimeSuggestion `json:"meetingTimeSuggestions,omitempty"`
	ODataType              *string                  `json:"@odata.type,omitempty"`
}
