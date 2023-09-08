package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseType string

const (
	EventMessageResponseTypeexception      EventMessageResponseType = "Exception"
	EventMessageResponseTypeoccurrence     EventMessageResponseType = "Occurrence"
	EventMessageResponseTypeseriesMaster   EventMessageResponseType = "SeriesMaster"
	EventMessageResponseTypesingleInstance EventMessageResponseType = "SingleInstance"
)

func PossibleValuesForEventMessageResponseType() []string {
	return []string{
		string(EventMessageResponseTypeexception),
		string(EventMessageResponseTypeoccurrence),
		string(EventMessageResponseTypeseriesMaster),
		string(EventMessageResponseTypesingleInstance),
	}
}

func parseEventMessageResponseType(input string) (*EventMessageResponseType, error) {
	vals := map[string]EventMessageResponseType{
		"exception":      EventMessageResponseTypeexception,
		"occurrence":     EventMessageResponseTypeoccurrence,
		"seriesmaster":   EventMessageResponseTypeseriesMaster,
		"singleinstance": EventMessageResponseTypesingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseType(input)
	return &out, nil
}
