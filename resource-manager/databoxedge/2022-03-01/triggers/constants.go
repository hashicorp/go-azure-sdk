package triggers

import "strings"

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
