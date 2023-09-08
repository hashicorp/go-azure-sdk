package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageMeetingMessageType string

const (
	EventMessageMeetingMessageTypemeetingAccepted           EventMessageMeetingMessageType = "MeetingAccepted"
	EventMessageMeetingMessageTypemeetingCancelled          EventMessageMeetingMessageType = "MeetingCancelled"
	EventMessageMeetingMessageTypemeetingDeclined           EventMessageMeetingMessageType = "MeetingDeclined"
	EventMessageMeetingMessageTypemeetingRequest            EventMessageMeetingMessageType = "MeetingRequest"
	EventMessageMeetingMessageTypemeetingTenativelyAccepted EventMessageMeetingMessageType = "MeetingTenativelyAccepted"
	EventMessageMeetingMessageTypenone                      EventMessageMeetingMessageType = "None"
)

func PossibleValuesForEventMessageMeetingMessageType() []string {
	return []string{
		string(EventMessageMeetingMessageTypemeetingAccepted),
		string(EventMessageMeetingMessageTypemeetingCancelled),
		string(EventMessageMeetingMessageTypemeetingDeclined),
		string(EventMessageMeetingMessageTypemeetingRequest),
		string(EventMessageMeetingMessageTypemeetingTenativelyAccepted),
		string(EventMessageMeetingMessageTypenone),
	}
}

func parseEventMessageMeetingMessageType(input string) (*EventMessageMeetingMessageType, error) {
	vals := map[string]EventMessageMeetingMessageType{
		"meetingaccepted":           EventMessageMeetingMessageTypemeetingAccepted,
		"meetingcancelled":          EventMessageMeetingMessageTypemeetingCancelled,
		"meetingdeclined":           EventMessageMeetingMessageTypemeetingDeclined,
		"meetingrequest":            EventMessageMeetingMessageTypemeetingRequest,
		"meetingtenativelyaccepted": EventMessageMeetingMessageTypemeetingTenativelyAccepted,
		"none":                      EventMessageMeetingMessageTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageMeetingMessageType(input)
	return &out, nil
}
