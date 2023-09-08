package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageRequestType string

const (
	EventMessageRequestTypeexception      EventMessageRequestType = "Exception"
	EventMessageRequestTypeoccurrence     EventMessageRequestType = "Occurrence"
	EventMessageRequestTypeseriesMaster   EventMessageRequestType = "SeriesMaster"
	EventMessageRequestTypesingleInstance EventMessageRequestType = "SingleInstance"
)

func PossibleValuesForEventMessageRequestType() []string {
	return []string{
		string(EventMessageRequestTypeexception),
		string(EventMessageRequestTypeoccurrence),
		string(EventMessageRequestTypeseriesMaster),
		string(EventMessageRequestTypesingleInstance),
	}
}

func parseEventMessageRequestType(input string) (*EventMessageRequestType, error) {
	vals := map[string]EventMessageRequestType{
		"exception":      EventMessageRequestTypeexception,
		"occurrence":     EventMessageRequestTypeoccurrence,
		"seriesmaster":   EventMessageRequestTypeseriesMaster,
		"singleinstance": EventMessageRequestTypesingleInstance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageRequestType(input)
	return &out, nil
}
