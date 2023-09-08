package user

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdFindMeetingTimeRequest struct {
	Attendees               *[]AttendeeBase     `json:"attendees,omitempty"`
	IsOrganizerOptional     *bool               `json:"isOrganizerOptional,omitempty"`
	LocationConstraint      *LocationConstraint `json:"locationConstraint,omitempty"`
	MaxCandidates           *int64              `json:"maxCandidates,omitempty"`
	MeetingDuration         *string             `json:"meetingDuration,omitempty"`
	ReturnSuggestionReasons *bool               `json:"returnSuggestionReasons,omitempty"`
	TimeConstraint          *TimeConstraint     `json:"timeConstraint,omitempty"`
}
