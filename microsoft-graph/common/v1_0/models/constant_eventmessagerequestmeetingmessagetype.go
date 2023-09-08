package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestMeetingMessageType string

const (
	EventMessageRequestMeetingMessageTypemeetingAccepted           EventMessageRequestMeetingMessageType = "MeetingAccepted"
	EventMessageRequestMeetingMessageTypemeetingCancelled          EventMessageRequestMeetingMessageType = "MeetingCancelled"
	EventMessageRequestMeetingMessageTypemeetingDeclined           EventMessageRequestMeetingMessageType = "MeetingDeclined"
	EventMessageRequestMeetingMessageTypemeetingRequest            EventMessageRequestMeetingMessageType = "MeetingRequest"
	EventMessageRequestMeetingMessageTypemeetingTenativelyAccepted EventMessageRequestMeetingMessageType = "MeetingTenativelyAccepted"
	EventMessageRequestMeetingMessageTypenone                      EventMessageRequestMeetingMessageType = "None"
)

func PossibleValuesForEventMessageRequestMeetingMessageType() []string {
	return []string{
		string(EventMessageRequestMeetingMessageTypemeetingAccepted),
		string(EventMessageRequestMeetingMessageTypemeetingCancelled),
		string(EventMessageRequestMeetingMessageTypemeetingDeclined),
		string(EventMessageRequestMeetingMessageTypemeetingRequest),
		string(EventMessageRequestMeetingMessageTypemeetingTenativelyAccepted),
		string(EventMessageRequestMeetingMessageTypenone),
	}
}

func parseEventMessageRequestMeetingMessageType(input string) (*EventMessageRequestMeetingMessageType, error) {
	vals := map[string]EventMessageRequestMeetingMessageType{
		"meetingaccepted":           EventMessageRequestMeetingMessageTypemeetingAccepted,
		"meetingcancelled":          EventMessageRequestMeetingMessageTypemeetingCancelled,
		"meetingdeclined":           EventMessageRequestMeetingMessageTypemeetingDeclined,
		"meetingrequest":            EventMessageRequestMeetingMessageTypemeetingRequest,
		"meetingtenativelyaccepted": EventMessageRequestMeetingMessageTypemeetingTenativelyAccepted,
		"none":                      EventMessageRequestMeetingMessageTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestMeetingMessageType(input)
	return &out, nil
}
