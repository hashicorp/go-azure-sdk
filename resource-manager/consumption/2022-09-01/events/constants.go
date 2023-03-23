package events

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventType string

const (
	EventTypeCreditExpired        EventType = "CreditExpired"
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
		string(EventTypeCreditExpired),
		string(EventTypeNewCredit),
		string(EventTypePendingAdjustments),
		string(EventTypePendingCharges),
		string(EventTypePendingExpiredCredit),
		string(EventTypePendingNewCredit),
		string(EventTypeSettledCharges),
		string(EventTypeUnKnown),
	}
}

func parseEventType(input string) (*EventType, error) {
	vals := map[string]EventType{
		"creditexpired":        EventTypeCreditExpired,
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
