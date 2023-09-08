package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageType string

const (
	EventMessageTypeexception      EventMessageType = "Exception"
	EventMessageTypeoccurrence     EventMessageType = "Occurrence"
	EventMessageTypeseriesMaster   EventMessageType = "SeriesMaster"
	EventMessageTypesingleInstance EventMessageType = "SingleInstance"
)

func PossibleValuesForEventMessageType() []string {
	return []string{
		string(EventMessageTypeexception),
		string(EventMessageTypeoccurrence),
		string(EventMessageTypeseriesMaster),
		string(EventMessageTypesingleInstance),
	}
}

func parseEventMessageType(input string) (*EventMessageType, error) {
	vals := map[string]EventMessageType{
		"exception":      EventMessageTypeexception,
		"occurrence":     EventMessageTypeoccurrence,
		"seriesmaster":   EventMessageTypeseriesMaster,
		"singleinstance": EventMessageTypesingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageType(input)
	return &out, nil
}
