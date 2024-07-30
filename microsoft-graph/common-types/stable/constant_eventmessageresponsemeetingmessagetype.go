package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseMeetingMessageType string

const (
	EventMessageResponseMeetingMessageType_MeetingAccepted           EventMessageResponseMeetingMessageType = "meetingAccepted"
	EventMessageResponseMeetingMessageType_MeetingCancelled          EventMessageResponseMeetingMessageType = "meetingCancelled"
	EventMessageResponseMeetingMessageType_MeetingDeclined           EventMessageResponseMeetingMessageType = "meetingDeclined"
	EventMessageResponseMeetingMessageType_MeetingRequest            EventMessageResponseMeetingMessageType = "meetingRequest"
	EventMessageResponseMeetingMessageType_MeetingTenativelyAccepted EventMessageResponseMeetingMessageType = "meetingTenativelyAccepted"
	EventMessageResponseMeetingMessageType_None                      EventMessageResponseMeetingMessageType = "none"
)

func PossibleValuesForEventMessageResponseMeetingMessageType() []string {
	return []string{
		string(EventMessageResponseMeetingMessageType_MeetingAccepted),
		string(EventMessageResponseMeetingMessageType_MeetingCancelled),
		string(EventMessageResponseMeetingMessageType_MeetingDeclined),
		string(EventMessageResponseMeetingMessageType_MeetingRequest),
		string(EventMessageResponseMeetingMessageType_MeetingTenativelyAccepted),
		string(EventMessageResponseMeetingMessageType_None),
	}
}

func (s *EventMessageResponseMeetingMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageResponseMeetingMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageResponseMeetingMessageType(input string) (*EventMessageResponseMeetingMessageType, error) {
	vals := map[string]EventMessageResponseMeetingMessageType{
		"meetingaccepted":           EventMessageResponseMeetingMessageType_MeetingAccepted,
		"meetingcancelled":          EventMessageResponseMeetingMessageType_MeetingCancelled,
		"meetingdeclined":           EventMessageResponseMeetingMessageType_MeetingDeclined,
		"meetingrequest":            EventMessageResponseMeetingMessageType_MeetingRequest,
		"meetingtenativelyaccepted": EventMessageResponseMeetingMessageType_MeetingTenativelyAccepted,
		"none":                      EventMessageResponseMeetingMessageType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseMeetingMessageType(input)
	return &out, nil
}
