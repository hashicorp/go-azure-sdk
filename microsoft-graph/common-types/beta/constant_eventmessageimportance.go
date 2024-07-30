package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageImportance string

const (
	EventMessageImportance_High   EventMessageImportance = "high"
	EventMessageImportance_Low    EventMessageImportance = "low"
	EventMessageImportance_Normal EventMessageImportance = "normal"
)

func PossibleValuesForEventMessageImportance() []string {
	return []string{
		string(EventMessageImportance_High),
		string(EventMessageImportance_Low),
		string(EventMessageImportance_Normal),
	}
}

func (s *EventMessageImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageImportance(input string) (*EventMessageImportance, error) {
	vals := map[string]EventMessageImportance{
		"high":   EventMessageImportance_High,
		"low":    EventMessageImportance_Low,
		"normal": EventMessageImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageImportance(input)
	return &out, nil
}
