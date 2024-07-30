package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestMeetingMessageType string

const (
	EventMessageRequestMeetingMessageType_MeetingAccepted           EventMessageRequestMeetingMessageType = "meetingAccepted"
	EventMessageRequestMeetingMessageType_MeetingCancelled          EventMessageRequestMeetingMessageType = "meetingCancelled"
	EventMessageRequestMeetingMessageType_MeetingDeclined           EventMessageRequestMeetingMessageType = "meetingDeclined"
	EventMessageRequestMeetingMessageType_MeetingRequest            EventMessageRequestMeetingMessageType = "meetingRequest"
	EventMessageRequestMeetingMessageType_MeetingTenativelyAccepted EventMessageRequestMeetingMessageType = "meetingTenativelyAccepted"
	EventMessageRequestMeetingMessageType_None                      EventMessageRequestMeetingMessageType = "none"
)

func PossibleValuesForEventMessageRequestMeetingMessageType() []string {
	return []string{
		string(EventMessageRequestMeetingMessageType_MeetingAccepted),
		string(EventMessageRequestMeetingMessageType_MeetingCancelled),
		string(EventMessageRequestMeetingMessageType_MeetingDeclined),
		string(EventMessageRequestMeetingMessageType_MeetingRequest),
		string(EventMessageRequestMeetingMessageType_MeetingTenativelyAccepted),
		string(EventMessageRequestMeetingMessageType_None),
	}
}

func (s *EventMessageRequestMeetingMessageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageRequestMeetingMessageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageRequestMeetingMessageType(input string) (*EventMessageRequestMeetingMessageType, error) {
	vals := map[string]EventMessageRequestMeetingMessageType{
		"meetingaccepted":           EventMessageRequestMeetingMessageType_MeetingAccepted,
		"meetingcancelled":          EventMessageRequestMeetingMessageType_MeetingCancelled,
		"meetingdeclined":           EventMessageRequestMeetingMessageType_MeetingDeclined,
		"meetingrequest":            EventMessageRequestMeetingMessageType_MeetingRequest,
		"meetingtenativelyaccepted": EventMessageRequestMeetingMessageType_MeetingTenativelyAccepted,
		"none":                      EventMessageRequestMeetingMessageType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestMeetingMessageType(input)
	return &out, nil
}
