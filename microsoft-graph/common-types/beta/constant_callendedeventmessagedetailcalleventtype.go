package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallEndedEventMessageDetailCallEventType string

const (
	CallEndedEventMessageDetailCallEventType_Call        CallEndedEventMessageDetailCallEventType = "call"
	CallEndedEventMessageDetailCallEventType_Meeting     CallEndedEventMessageDetailCallEventType = "meeting"
	CallEndedEventMessageDetailCallEventType_ScreenShare CallEndedEventMessageDetailCallEventType = "screenShare"
)

func PossibleValuesForCallEndedEventMessageDetailCallEventType() []string {
	return []string{
		string(CallEndedEventMessageDetailCallEventType_Call),
		string(CallEndedEventMessageDetailCallEventType_Meeting),
		string(CallEndedEventMessageDetailCallEventType_ScreenShare),
	}
}

func (s *CallEndedEventMessageDetailCallEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallEndedEventMessageDetailCallEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallEndedEventMessageDetailCallEventType(input string) (*CallEndedEventMessageDetailCallEventType, error) {
	vals := map[string]CallEndedEventMessageDetailCallEventType{
		"call":        CallEndedEventMessageDetailCallEventType_Call,
		"meeting":     CallEndedEventMessageDetailCallEventType_Meeting,
		"screenshare": CallEndedEventMessageDetailCallEventType_ScreenShare,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallEndedEventMessageDetailCallEventType(input)
	return &out, nil
}
