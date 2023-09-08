package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchBookmarkState string

const (
	SearchBookmarkStatedraft     SearchBookmarkState = "Draft"
	SearchBookmarkStateexcluded  SearchBookmarkState = "Excluded"
	SearchBookmarkStatepublished SearchBookmarkState = "Published"
)

func PossibleValuesForSearchBookmarkState() []string {
	return []string{
		string(SearchBookmarkStatedraft),
		string(SearchBookmarkStateexcluded),
		string(SearchBookmarkStatepublished),
	}
}

func parseSearchBookmarkState(input string) (*SearchBookmarkState, error) {
	vals := map[string]SearchBookmarkState{
		"draft":     SearchBookmarkStatedraft,
		"excluded":  SearchBookmarkStateexcluded,
		"published": SearchBookmarkStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchBookmarkState(input)
	return &out, nil
}
