package events

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventType string

const (
	EventTypeNewCredit            EventType = "NewCredit"
	EventTypePendingAdjustments   EventType = "PendingAdjustments"
	EventTypePendingCharges       EventType = "PendingCharges"
	EventTypePendingExpiredCredit EventType = "PendingExpiredCredit"
	EventTypePendingNewCredit     EventType = "PendingNewCredit"
	EventTypeSettledCharges       EventType = "SettledCharges"
	EventTypeUnKnown              EventType = "UnKnown"
)

func PossibleValuesForEventType() []string {
	return []string{
		string(EventTypeNewCredit),
		string(EventTypePendingAdjustments),
		string(EventTypePendingCharges),
		string(EventTypePendingExpiredCredit),
		string(EventTypePendingNewCredit),
		string(EventTypeSettledCharges),
		string(EventTypeUnKnown),
	}
}

func (s *EventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEventType(input string) (*EventType, error) {
	vals := map[string]EventType{
		"newcredit":            EventTypeNewCredit,
		"pendingadjustments":   EventTypePendingAdjustments,
		"pendingcharges":       EventTypePendingCharges,
		"pendingexpiredcredit": EventTypePendingExpiredCredit,
		"pendingnewcredit":     EventTypePendingNewCredit,
		"settledcharges":       EventTypeSettledCharges,
		"unknown":              EventTypeUnKnown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventType(input)
	return &out, nil
}
