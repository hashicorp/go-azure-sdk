package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponseResponseType string

const (
	EventMessageResponseResponseTypeaccepted            EventMessageResponseResponseType = "Accepted"
	EventMessageResponseResponseTypedeclined            EventMessageResponseResponseType = "Declined"
	EventMessageResponseResponseTypenone                EventMessageResponseResponseType = "None"
	EventMessageResponseResponseTypenotResponded        EventMessageResponseResponseType = "NotResponded"
	EventMessageResponseResponseTypeorganizer           EventMessageResponseResponseType = "Organizer"
	EventMessageResponseResponseTypetentativelyAccepted EventMessageResponseResponseType = "TentativelyAccepted"
)

func PossibleValuesForEventMessageResponseResponseType() []string {
	return []string{
		string(EventMessageResponseResponseTypeaccepted),
		string(EventMessageResponseResponseTypedeclined),
		string(EventMessageResponseResponseTypenone),
		string(EventMessageResponseResponseTypenotResponded),
		string(EventMessageResponseResponseTypeorganizer),
		string(EventMessageResponseResponseTypetentativelyAccepted),
	}
}

func parseEventMessageResponseResponseType(input string) (*EventMessageResponseResponseType, error) {
	vals := map[string]EventMessageResponseResponseType{
		"accepted":            EventMessageResponseResponseTypeaccepted,
		"declined":            EventMessageResponseResponseTypedeclined,
		"none":                EventMessageResponseResponseTypenone,
		"notresponded":        EventMessageResponseResponseTypenotResponded,
		"organizer":           EventMessageResponseResponseTypeorganizer,
		"tentativelyaccepted": EventMessageResponseResponseTypetentativelyAccepted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventMessageResponseResponseType(input)
	return &out, nil
}
