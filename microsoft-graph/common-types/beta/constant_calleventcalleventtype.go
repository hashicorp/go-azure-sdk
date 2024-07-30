package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallEventCallEventType string

const (
	CallEventCallEventType_CallEnded   CallEventCallEventType = "callEnded"
	CallEventCallEventType_CallStarted CallEventCallEventType = "callStarted"
)

func PossibleValuesForCallEventCallEventType() []string {
	return []string{
		string(CallEventCallEventType_CallEnded),
		string(CallEventCallEventType_CallStarted),
	}
}

func (s *CallEventCallEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallEventCallEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallEventCallEventType(input string) (*CallEventCallEventType, error) {
	vals := map[string]CallEventCallEventType{
		"callended":   CallEventCallEventType_CallEnded,
		"callstarted": CallEventCallEventType_CallStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallEventCallEventType(input)
	return &out, nil
}
