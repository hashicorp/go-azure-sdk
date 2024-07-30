package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchRequestEntityTypes string

const (
	SearchRequestEntityTypes_Acronym      SearchRequestEntityTypes = "acronym"
	SearchRequestEntityTypes_Bookmark     SearchRequestEntityTypes = "bookmark"
	SearchRequestEntityTypes_ChatMessage  SearchRequestEntityTypes = "chatMessage"
	SearchRequestEntityTypes_Drive        SearchRequestEntityTypes = "drive"
	SearchRequestEntityTypes_DriveItem    SearchRequestEntityTypes = "driveItem"
	SearchRequestEntityTypes_Event        SearchRequestEntityTypes = "event"
	SearchRequestEntityTypes_ExternalItem SearchRequestEntityTypes = "externalItem"
	SearchRequestEntityTypes_List         SearchRequestEntityTypes = "list"
	SearchRequestEntityTypes_ListItem     SearchRequestEntityTypes = "listItem"
	SearchRequestEntityTypes_Message      SearchRequestEntityTypes = "message"
	SearchRequestEntityTypes_Person       SearchRequestEntityTypes = "person"
	SearchRequestEntityTypes_Qna          SearchRequestEntityTypes = "qna"
	SearchRequestEntityTypes_Site         SearchRequestEntityTypes = "site"
)

func PossibleValuesForSearchRequestEntityTypes() []string {
	return []string{
		string(SearchRequestEntityTypes_Acronym),
		string(SearchRequestEntityTypes_Bookmark),
		string(SearchRequestEntityTypes_ChatMessage),
		string(SearchRequestEntityTypes_Drive),
		string(SearchRequestEntityTypes_DriveItem),
		string(SearchRequestEntityTypes_Event),
		string(SearchRequestEntityTypes_ExternalItem),
		string(SearchRequestEntityTypes_List),
		string(SearchRequestEntityTypes_ListItem),
		string(SearchRequestEntityTypes_Message),
		string(SearchRequestEntityTypes_Person),
		string(SearchRequestEntityTypes_Qna),
		string(SearchRequestEntityTypes_Site),
	}
}

func (s *SearchRequestEntityTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchRequestEntityTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchRequestEntityTypes(input string) (*SearchRequestEntityTypes, error) {
	vals := map[string]SearchRequestEntityTypes{
		"acronym":      SearchRequestEntityTypes_Acronym,
		"bookmark":     SearchRequestEntityTypes_Bookmark,
		"chatmessage":  SearchRequestEntityTypes_ChatMessage,
		"drive":        SearchRequestEntityTypes_Drive,
		"driveitem":    SearchRequestEntityTypes_DriveItem,
		"event":        SearchRequestEntityTypes_Event,
		"externalitem": SearchRequestEntityTypes_ExternalItem,
		"list":         SearchRequestEntityTypes_List,
		"listitem":     SearchRequestEntityTypes_ListItem,
		"message":      SearchRequestEntityTypes_Message,
		"person":       SearchRequestEntityTypes_Person,
		"qna":          SearchRequestEntityTypes_Qna,
		"site":         SearchRequestEntityTypes_Site,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchRequestEntityTypes(input)
	return &out, nil
}
