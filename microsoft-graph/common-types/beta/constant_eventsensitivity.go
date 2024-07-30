package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventSensitivity string

const (
	EventSensitivity_Confidential EventSensitivity = "confidential"
	EventSensitivity_Normal       EventSensitivity = "normal"
	EventSensitivity_Personal     EventSensitivity = "personal"
	EventSensitivity_Private      EventSensitivity = "private"
)

func PossibleValuesForEventSensitivity() []string {
	return []string{
		string(EventSensitivity_Confidential),
		string(EventSensitivity_Normal),
		string(EventSensitivity_Personal),
		string(EventSensitivity_Private),
	}
}

func (s *EventSensitivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventSensitivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventSensitivity(input string) (*EventSensitivity, error) {
	vals := map[string]EventSensitivity{
		"confidential": EventSensitivity_Confidential,
		"normal":       EventSensitivity_Normal,
		"personal":     EventSensitivity_Personal,
		"private":      EventSensitivity_Private,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventSensitivity(input)
	return &out, nil
}
