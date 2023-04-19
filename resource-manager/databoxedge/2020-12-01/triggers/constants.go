package triggers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerEventType string

const (
	TriggerEventTypeFileEvent          TriggerEventType = "FileEvent"
	TriggerEventTypePeriodicTimerEvent TriggerEventType = "PeriodicTimerEvent"
)

func PossibleValuesForTriggerEventType() []string {
	return []string{
		string(TriggerEventTypeFileEvent),
		string(TriggerEventTypePeriodicTimerEvent),
	}
}

func (s *TriggerEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggerEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggerEventType(input string) (*TriggerEventType, error) {
	vals := map[string]TriggerEventType{
		"fileevent":          TriggerEventTypeFileEvent,
		"periodictimerevent": TriggerEventTypePeriodicTimerEvent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerEventType(input)
	return &out, nil
}
