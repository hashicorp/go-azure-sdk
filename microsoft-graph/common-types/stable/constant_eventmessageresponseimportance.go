package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseImportance string

const (
	EventMessageResponseImportance_High   EventMessageResponseImportance = "high"
	EventMessageResponseImportance_Low    EventMessageResponseImportance = "low"
	EventMessageResponseImportance_Normal EventMessageResponseImportance = "normal"
)

func PossibleValuesForEventMessageResponseImportance() []string {
	return []string{
		string(EventMessageResponseImportance_High),
		string(EventMessageResponseImportance_Low),
		string(EventMessageResponseImportance_Normal),
	}
}

func (s *EventMessageResponseImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageResponseImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageResponseImportance(input string) (*EventMessageResponseImportance, error) {
	vals := map[string]EventMessageResponseImportance{
		"high":   EventMessageResponseImportance_High,
		"low":    EventMessageResponseImportance_Low,
		"normal": EventMessageResponseImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseImportance(input)
	return &out, nil
}
