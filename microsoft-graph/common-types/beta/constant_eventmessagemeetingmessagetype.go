package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageMeetingMessageType string

const (
	EventMessageMeetingMessageType_MeetingAccepted            EventMessageMeetingMessageType = "meetingAccepted"
	EventMessageMeetingMessageType_MeetingCancelled           EventMessageMeetingMessageType = "meetingCancelled"
	EventMessageMeetingMessageType_MeetingDeclined            EventMessageMeetingMessageType = "meetingDeclined"
	EventMessageMeetingMessageType_MeetingRequest             EventMessageMeetingMessageType = "meetingRequest"
	EventMessageMeetingMessageType_MeetingTentativelyAccepted EventMessageMeetingMessageType = "meetingTentativelyAccepted"
	EventMessageMeetingMessageType_None                       EventMessageMeetingMessageType = "none"
)

func PossibleValuesForEventMessageMeetingMessageType() []string {
	return []string{
		string(EventMessageMeetingMessageType_MeetingAccepted),
		string(EventMessageMeetingMessageType_MeetingCancelled),
		string(EventMessageMeetingMessageType_MeetingDeclined),
		string(EventMessageMeetingMessageType_MeetingRequest),
		string(EventMessageMeetingMessageType_MeetingTentativelyAccepted),
		string(EventMessageMeetingMessageType_None),
	}
}

func (s *EventMessageMeetingMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageMeetingMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageMeetingMessageType(input string) (*EventMessageMeetingMessageType, error) {
	vals := map[string]EventMessageMeetingMessageType{
		"meetingaccepted":            EventMessageMeetingMessageType_MeetingAccepted,
		"meetingcancelled":           EventMessageMeetingMessageType_MeetingCancelled,
		"meetingdeclined":            EventMessageMeetingMessageType_MeetingDeclined,
		"meetingrequest":             EventMessageMeetingMessageType_MeetingRequest,
		"meetingtentativelyaccepted": EventMessageMeetingMessageType_MeetingTentativelyAccepted,
		"none":                       EventMessageMeetingMessageType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageMeetingMessageType(input)
	return &out, nil
}
