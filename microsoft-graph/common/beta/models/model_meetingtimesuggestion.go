package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingTimeSuggestion struct {
	AttendeeAvailability  *[]AttendeeAvailability                     `json:"attendeeAvailability,omitempty"`
	Locations             *[]Location                                 `json:"locations,omitempty"`
	MeetingTimeSlot       *TimeSlot                                   `json:"meetingTimeSlot,omitempty"`
	ODataType             *string                                     `json:"@odata.type,omitempty"`
	Order                 *int64                                      `json:"order,omitempty"`
	OrganizerAvailability *MeetingTimeSuggestionOrganizerAvailability `json:"organizerAvailability,omitempty"`
	SuggestionReason      *string                                     `json:"suggestionReason,omitempty"`
}
