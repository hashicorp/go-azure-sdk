package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchBookmarkState string

const (
	SearchBookmarkState_Draft     SearchBookmarkState = "draft"
	SearchBookmarkState_Excluded  SearchBookmarkState = "excluded"
	SearchBookmarkState_Published SearchBookmarkState = "published"
)

func PossibleValuesForSearchBookmarkState() []string {
	return []string{
		string(SearchBookmarkState_Draft),
		string(SearchBookmarkState_Excluded),
		string(SearchBookmarkState_Published),
	}
}

func (s *SearchBookmarkState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchBookmarkState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchBookmarkState(input string) (*SearchBookmarkState, error) {
	vals := map[string]SearchBookmarkState{
		"draft":     SearchBookmarkState_Draft,
		"excluded":  SearchBookmarkState_Excluded,
		"published": SearchBookmarkState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchBookmarkState(input)
	return &out, nil
}
