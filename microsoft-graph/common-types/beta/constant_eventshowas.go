package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventShowAs string

const (
	EventShowAs_Busy             EventShowAs = "busy"
	EventShowAs_Free             EventShowAs = "free"
	EventShowAs_Oof              EventShowAs = "oof"
	EventShowAs_Tentative        EventShowAs = "tentative"
	EventShowAs_Unknown          EventShowAs = "unknown"
	EventShowAs_WorkingElsewhere EventShowAs = "workingElsewhere"
)

func PossibleValuesForEventShowAs() []string {
	return []string{
		string(EventShowAs_Busy),
		string(EventShowAs_Free),
		string(EventShowAs_Oof),
		string(EventShowAs_Tentative),
		string(EventShowAs_Unknown),
		string(EventShowAs_WorkingElsewhere),
	}
}

func (s *EventShowAs) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventShowAs(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventShowAs(input string) (*EventShowAs, error) {
	vals := map[string]EventShowAs{
		"busy":             EventShowAs_Busy,
		"free":             EventShowAs_Free,
		"oof":              EventShowAs_Oof,
		"tentative":        EventShowAs_Tentative,
		"unknown":          EventShowAs_Unknown,
		"workingelsewhere": EventShowAs_WorkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventShowAs(input)
	return &out, nil
}
