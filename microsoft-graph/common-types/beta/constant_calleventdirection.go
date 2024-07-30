package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallEventDirection string

const (
	CallEventDirection_Incoming CallEventDirection = "incoming"
	CallEventDirection_Outgoing CallEventDirection = "outgoing"
)

func PossibleValuesForCallEventDirection() []string {
	return []string{
		string(CallEventDirection_Incoming),
		string(CallEventDirection_Outgoing),
	}
}

func (s *CallEventDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallEventDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallEventDirection(input string) (*CallEventDirection, error) {
	vals := map[string]CallEventDirection{
		"incoming": CallEventDirection_Incoming,
		"outgoing": CallEventDirection_Outgoing,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallEventDirection(input)
	return &out, nil
}
