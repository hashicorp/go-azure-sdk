package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallDirection string

const (
	CallDirection_Incoming CallDirection = "incoming"
	CallDirection_Outgoing CallDirection = "outgoing"
)

func PossibleValuesForCallDirection() []string {
	return []string{
		string(CallDirection_Incoming),
		string(CallDirection_Outgoing),
	}
}

func (s *CallDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallDirection(input string) (*CallDirection, error) {
	vals := map[string]CallDirection{
		"incoming": CallDirection_Incoming,
		"outgoing": CallDirection_Outgoing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallDirection(input)
	return &out, nil
}
