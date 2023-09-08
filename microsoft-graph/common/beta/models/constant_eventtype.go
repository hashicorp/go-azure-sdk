package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventType string

const (
	EventTypeexception      EventType = "Exception"
	EventTypeoccurrence     EventType = "Occurrence"
	EventTypeseriesMaster   EventType = "SeriesMaster"
	EventTypesingleInstance EventType = "SingleInstance"
)

func PossibleValuesForEventType() []string {
	return []string{
		string(EventTypeexception),
		string(EventTypeoccurrence),
		string(EventTypeseriesMaster),
		string(EventTypesingleInstance),
	}
}

func parseEventType(input string) (*EventType, error) {
	vals := map[string]EventType{
		"exception":      EventTypeexception,
		"occurrence":     EventTypeoccurrence,
		"seriesmaster":   EventTypeseriesMaster,
		"singleinstance": EventTypesingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventType(input)
	return &out, nil
}
