package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventImportance string

const (
	EventImportance_High   EventImportance = "high"
	EventImportance_Low    EventImportance = "low"
	EventImportance_Normal EventImportance = "normal"
)

func PossibleValuesForEventImportance() []string {
	return []string{
		string(EventImportance_High),
		string(EventImportance_Low),
		string(EventImportance_Normal),
	}
}

func (s *EventImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventImportance(input string) (*EventImportance, error) {
	vals := map[string]EventImportance{
		"high":   EventImportance_High,
		"low":    EventImportance_Low,
		"normal": EventImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventImportance(input)
	return &out, nil
}
