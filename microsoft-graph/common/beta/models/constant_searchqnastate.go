package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchQnaState string

const (
	SearchQnaStatedraft     SearchQnaState = "Draft"
	SearchQnaStateexcluded  SearchQnaState = "Excluded"
	SearchQnaStatepublished SearchQnaState = "Published"
)

func PossibleValuesForSearchQnaState() []string {
	return []string{
		string(SearchQnaStatedraft),
		string(SearchQnaStateexcluded),
		string(SearchQnaStatepublished),
	}
}

func parseSearchQnaState(input string) (*SearchQnaState, error) {
	vals := map[string]SearchQnaState{
		"draft":     SearchQnaStatedraft,
		"excluded":  SearchQnaStateexcluded,
		"published": SearchQnaStatepublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchQnaState(input)
	return &out, nil
}
