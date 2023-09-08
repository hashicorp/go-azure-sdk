package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchRequestEntityTypes string

const (
	SearchRequestEntityTypesacronym      SearchRequestEntityTypes = "Acronym"
	SearchRequestEntityTypesbookmark     SearchRequestEntityTypes = "Bookmark"
	SearchRequestEntityTypeschatMessage  SearchRequestEntityTypes = "ChatMessage"
	SearchRequestEntityTypesdrive        SearchRequestEntityTypes = "Drive"
	SearchRequestEntityTypesdriveItem    SearchRequestEntityTypes = "DriveItem"
	SearchRequestEntityTypesevent        SearchRequestEntityTypes = "Event"
	SearchRequestEntityTypesexternalItem SearchRequestEntityTypes = "ExternalItem"
	SearchRequestEntityTypeslist         SearchRequestEntityTypes = "List"
	SearchRequestEntityTypeslistItem     SearchRequestEntityTypes = "ListItem"
	SearchRequestEntityTypesmessage      SearchRequestEntityTypes = "Message"
	SearchRequestEntityTypesperson       SearchRequestEntityTypes = "Person"
	SearchRequestEntityTypessite         SearchRequestEntityTypes = "Site"
)

func PossibleValuesForSearchRequestEntityTypes() []string {
	return []string{
		string(SearchRequestEntityTypesacronym),
		string(SearchRequestEntityTypesbookmark),
		string(SearchRequestEntityTypeschatMessage),
		string(SearchRequestEntityTypesdrive),
		string(SearchRequestEntityTypesdriveItem),
		string(SearchRequestEntityTypesevent),
		string(SearchRequestEntityTypesexternalItem),
		string(SearchRequestEntityTypeslist),
		string(SearchRequestEntityTypeslistItem),
		string(SearchRequestEntityTypesmessage),
		string(SearchRequestEntityTypesperson),
		string(SearchRequestEntityTypessite),
	}
}

func parseSearchRequestEntityTypes(input string) (*SearchRequestEntityTypes, error) {
	vals := map[string]SearchRequestEntityTypes{
		"acronym":      SearchRequestEntityTypesacronym,
		"bookmark":     SearchRequestEntityTypesbookmark,
		"chatmessage":  SearchRequestEntityTypeschatMessage,
		"drive":        SearchRequestEntityTypesdrive,
		"driveitem":    SearchRequestEntityTypesdriveItem,
		"event":        SearchRequestEntityTypesevent,
		"externalitem": SearchRequestEntityTypesexternalItem,
		"list":         SearchRequestEntityTypeslist,
		"listitem":     SearchRequestEntityTypeslistItem,
		"message":      SearchRequestEntityTypesmessage,
		"person":       SearchRequestEntityTypesperson,
		"site":         SearchRequestEntityTypessite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchRequestEntityTypes(input)
	return &out, nil
}
