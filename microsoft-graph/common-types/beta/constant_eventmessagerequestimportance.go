package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestImportance string

const (
	EventMessageRequestImportance_High   EventMessageRequestImportance = "high"
	EventMessageRequestImportance_Low    EventMessageRequestImportance = "low"
	EventMessageRequestImportance_Normal EventMessageRequestImportance = "normal"
)

func PossibleValuesForEventMessageRequestImportance() []string {
	return []string{
		string(EventMessageRequestImportance_High),
		string(EventMessageRequestImportance_Low),
		string(EventMessageRequestImportance_Normal),
	}
}

func (s *EventMessageRequestImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventMessageRequestImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventMessageRequestImportance(input string) (*EventMessageRequestImportance, error) {
	vals := map[string]EventMessageRequestImportance{
		"high":   EventMessageRequestImportance_High,
		"low":    EventMessageRequestImportance_Low,
		"normal": EventMessageRequestImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestImportance(input)
	return &out, nil
}
