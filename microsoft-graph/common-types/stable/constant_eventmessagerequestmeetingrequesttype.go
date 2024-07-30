package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestMeetingRequestType string

const (
	EventMessageRequestMeetingRequestType_FullUpdate          EventMessageRequestMeetingRequestType = "fullUpdate"
	EventMessageRequestMeetingRequestType_InformationalUpdate EventMessageRequestMeetingRequestType = "informationalUpdate"
	EventMessageRequestMeetingRequestType_NewMeetingRequest   EventMessageRequestMeetingRequestType = "newMeetingRequest"
	EventMessageRequestMeetingRequestType_None                EventMessageRequestMeetingRequestType = "none"
	EventMessageRequestMeetingRequestType_Outdated            EventMessageRequestMeetingRequestType = "outdated"
	EventMessageRequestMeetingRequestType_PrincipalWantsCopy  EventMessageRequestMeetingRequestType = "principalWantsCopy"
	EventMessageRequestMeetingRequestType_SilentUpdate        EventMessageRequestMeetingRequestType = "silentUpdate"
)

func PossibleValuesForEventMessageRequestMeetingRequestType() []string {
	return []string{
		string(EventMessageRequestMeetingRequestType_FullUpdate),
		string(EventMessageRequestMeetingRequestType_InformationalUpdate),
		string(EventMessageRequestMeetingRequestType_NewMeetingRequest),
		string(EventMessageRequestMeetingRequestType_None),
		string(EventMessageRequestMeetingRequestType_Outdated),
		string(EventMessageRequestMeetingRequestType_PrincipalWantsCopy),
		string(EventMessageRequestMeetingRequestType_SilentUpdate),
	}
}

func (s *EventMessageRequestMeetingRequestType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageRequestMeetingRequestType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageRequestMeetingRequestType(input string) (*EventMessageRequestMeetingRequestType, error) {
	vals := map[string]EventMessageRequestMeetingRequestType{
		"fullupdate":          EventMessageRequestMeetingRequestType_FullUpdate,
		"informationalupdate": EventMessageRequestMeetingRequestType_InformationalUpdate,
		"newmeetingrequest":   EventMessageRequestMeetingRequestType_NewMeetingRequest,
		"none":                EventMessageRequestMeetingRequestType_None,
		"outdated":            EventMessageRequestMeetingRequestType_Outdated,
		"principalwantscopy":  EventMessageRequestMeetingRequestType_PrincipalWantsCopy,
		"silentupdate":        EventMessageRequestMeetingRequestType_SilentUpdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestMeetingRequestType(input)
	return &out, nil
}
