package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingOperationPredicate struct {
	AllowAttendeeToEnableCamera   *bool
	AllowAttendeeToEnableMic      *bool
	AllowParticipantsToChangeName *bool
	AllowTeamworkReactions        *bool
	AttendeeReport                *string
	CreationDateTime              *string
	EndDateTime                   *string
	ExternalId                    *string
	Id                            *string
	IsBroadcast                   *bool
	IsEntryExitAnnounced          *bool
	JoinWebUrl                    *string
	ODataType                     *string
	RecordAutomatically           *bool
	StartDateTime                 *string
	Subject                       *string
	VideoTeleconferenceId         *string
}

func (p OnlineMeetingOperationPredicate) Matches(input OnlineMeeting) bool {

	if p.AllowAttendeeToEnableCamera != nil && (input.AllowAttendeeToEnableCamera == nil || *p.AllowAttendeeToEnableCamera != *input.AllowAttendeeToEnableCamera) {
		return false
	}

	if p.AllowAttendeeToEnableMic != nil && (input.AllowAttendeeToEnableMic == nil || *p.AllowAttendeeToEnableMic != *input.AllowAttendeeToEnableMic) {
		return false
	}

	if p.AllowParticipantsToChangeName != nil && (input.AllowParticipantsToChangeName == nil || *p.AllowParticipantsToChangeName != *input.AllowParticipantsToChangeName) {
		return false
	}

	if p.AllowTeamworkReactions != nil && (input.AllowTeamworkReactions == nil || *p.AllowTeamworkReactions != *input.AllowTeamworkReactions) {
		return false
	}

	if p.AttendeeReport != nil && (input.AttendeeReport == nil || *p.AttendeeReport != *input.AttendeeReport) {
		return false
	}

	if p.CreationDateTime != nil && (input.CreationDateTime == nil || *p.CreationDateTime != *input.CreationDateTime) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.ExternalId != nil && (input.ExternalId == nil || *p.ExternalId != *input.ExternalId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsBroadcast != nil && (input.IsBroadcast == nil || *p.IsBroadcast != *input.IsBroadcast) {
		return false
	}

	if p.IsEntryExitAnnounced != nil && (input.IsEntryExitAnnounced == nil || *p.IsEntryExitAnnounced != *input.IsEntryExitAnnounced) {
		return false
	}

	if p.JoinWebUrl != nil && (input.JoinWebUrl == nil || *p.JoinWebUrl != *input.JoinWebUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecordAutomatically != nil && (input.RecordAutomatically == nil || *p.RecordAutomatically != *input.RecordAutomatically) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	if p.VideoTeleconferenceId != nil && (input.VideoTeleconferenceId == nil || *p.VideoTeleconferenceId != *input.VideoTeleconferenceId) {
		return false
	}

	return true
}
