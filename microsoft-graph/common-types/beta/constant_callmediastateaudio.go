package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallMediaStateAudio string

const (
	CallMediaStateAudio_Active   CallMediaStateAudio = "active"
	CallMediaStateAudio_Inactive CallMediaStateAudio = "inactive"
)

func PossibleValuesForCallMediaStateAudio() []string {
	return []string{
		string(CallMediaStateAudio_Active),
		string(CallMediaStateAudio_Inactive),
	}
}

func (s *CallMediaStateAudio) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallMediaStateAudio(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallMediaStateAudio(input string) (*CallMediaStateAudio, error) {
	vals := map[string]CallMediaStateAudio{
		"active":   CallMediaStateAudio_Active,
		"inactive": CallMediaStateAudio_Inactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallMediaStateAudio(input)
	return &out, nil
}
