package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseMeetingMessageType string

const (
	EventMessageResponseMeetingMessageTypemeetingAccepted            EventMessageResponseMeetingMessageType = "MeetingAccepted"
	EventMessageResponseMeetingMessageTypemeetingCancelled           EventMessageResponseMeetingMessageType = "MeetingCancelled"
	EventMessageResponseMeetingMessageTypemeetingDeclined            EventMessageResponseMeetingMessageType = "MeetingDeclined"
	EventMessageResponseMeetingMessageTypemeetingRequest             EventMessageResponseMeetingMessageType = "MeetingRequest"
	EventMessageResponseMeetingMessageTypemeetingTentativelyAccepted EventMessageResponseMeetingMessageType = "MeetingTentativelyAccepted"
	EventMessageResponseMeetingMessageTypenone                       EventMessageResponseMeetingMessageType = "None"
)

func PossibleValuesForEventMessageResponseMeetingMessageType() []string {
	return []string{
		string(EventMessageResponseMeetingMessageTypemeetingAccepted),
		string(EventMessageResponseMeetingMessageTypemeetingCancelled),
		string(EventMessageResponseMeetingMessageTypemeetingDeclined),
		string(EventMessageResponseMeetingMessageTypemeetingRequest),
		string(EventMessageResponseMeetingMessageTypemeetingTentativelyAccepted),
		string(EventMessageResponseMeetingMessageTypenone),
	}
}

func parseEventMessageResponseMeetingMessageType(input string) (*EventMessageResponseMeetingMessageType, error) {
	vals := map[string]EventMessageResponseMeetingMessageType{
		"meetingaccepted":            EventMessageResponseMeetingMessageTypemeetingAccepted,
		"meetingcancelled":           EventMessageResponseMeetingMessageTypemeetingCancelled,
		"meetingdeclined":            EventMessageResponseMeetingMessageTypemeetingDeclined,
		"meetingrequest":             EventMessageResponseMeetingMessageTypemeetingRequest,
		"meetingtentativelyaccepted": EventMessageResponseMeetingMessageTypemeetingTentativelyAccepted,
		"none":                       EventMessageResponseMeetingMessageTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseMeetingMessageType(input)
	return &out, nil
}
