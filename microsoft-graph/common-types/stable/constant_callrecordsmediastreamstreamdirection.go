package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsMediaStreamStreamDirection string

const (
	CallRecordsMediaStreamStreamDirection_CalleeToCaller CallRecordsMediaStreamStreamDirection = "calleeToCaller"
	CallRecordsMediaStreamStreamDirection_CallerToCallee CallRecordsMediaStreamStreamDirection = "callerToCallee"
)

func PossibleValuesForCallRecordsMediaStreamStreamDirection() []string {
	return []string{
		string(CallRecordsMediaStreamStreamDirection_CalleeToCaller),
		string(CallRecordsMediaStreamStreamDirection_CallerToCallee),
	}
}

func (s *CallRecordsMediaStreamStreamDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallRecordsMediaStreamStreamDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallRecordsMediaStreamStreamDirection(input string) (*CallRecordsMediaStreamStreamDirection, error) {
	vals := map[string]CallRecordsMediaStreamStreamDirection{
		"calleetocaller": CallRecordsMediaStreamStreamDirection_CalleeToCaller,
		"callertocallee": CallRecordsMediaStreamStreamDirection_CallerToCallee,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallRecordsMediaStreamStreamDirection(input)
	return &out, nil
}
