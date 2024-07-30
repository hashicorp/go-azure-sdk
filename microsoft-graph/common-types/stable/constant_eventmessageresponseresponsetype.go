package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseResponseType string

const (
	EventMessageResponseResponseType_Accepted            EventMessageResponseResponseType = "accepted"
	EventMessageResponseResponseType_Declined            EventMessageResponseResponseType = "declined"
	EventMessageResponseResponseType_None                EventMessageResponseResponseType = "none"
	EventMessageResponseResponseType_NotResponded        EventMessageResponseResponseType = "notResponded"
	EventMessageResponseResponseType_Organizer           EventMessageResponseResponseType = "organizer"
	EventMessageResponseResponseType_TentativelyAccepted EventMessageResponseResponseType = "tentativelyAccepted"
)

func PossibleValuesForEventMessageResponseResponseType() []string {
	return []string{
		string(EventMessageResponseResponseType_Accepted),
		string(EventMessageResponseResponseType_Declined),
		string(EventMessageResponseResponseType_None),
		string(EventMessageResponseResponseType_NotResponded),
		string(EventMessageResponseResponseType_Organizer),
		string(EventMessageResponseResponseType_TentativelyAccepted),
	}
}

func (s *EventMessageResponseResponseType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageResponseResponseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageResponseResponseType(input string) (*EventMessageResponseResponseType, error) {
	vals := map[string]EventMessageResponseResponseType{
		"accepted":            EventMessageResponseResponseType_Accepted,
		"declined":            EventMessageResponseResponseType_Declined,
		"none":                EventMessageResponseResponseType_None,
		"notresponded":        EventMessageResponseResponseType_NotResponded,
		"organizer":           EventMessageResponseResponseType_Organizer,
		"tentativelyaccepted": EventMessageResponseResponseType_TentativelyAccepted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseResponseType(input)
	return &out, nil
}
