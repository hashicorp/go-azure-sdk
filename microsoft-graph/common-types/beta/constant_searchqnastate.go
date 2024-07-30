package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchQnaState string

const (
	SearchQnaState_Draft     SearchQnaState = "draft"
	SearchQnaState_Excluded  SearchQnaState = "excluded"
	SearchQnaState_Published SearchQnaState = "published"
)

func PossibleValuesForSearchQnaState() []string {
	return []string{
		string(SearchQnaState_Draft),
		string(SearchQnaState_Excluded),
		string(SearchQnaState_Published),
	}
}

func (s *SearchQnaState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchQnaState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchQnaState(input string) (*SearchQnaState, error) {
	vals := map[string]SearchQnaState{
		"draft":     SearchQnaState_Draft,
		"excluded":  SearchQnaState_Excluded,
		"published": SearchQnaState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchQnaState(input)
	return &out, nil
}
