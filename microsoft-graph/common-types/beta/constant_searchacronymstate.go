package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAcronymState string

const (
	SearchAcronymState_Draft     SearchAcronymState = "draft"
	SearchAcronymState_Excluded  SearchAcronymState = "excluded"
	SearchAcronymState_Published SearchAcronymState = "published"
)

func PossibleValuesForSearchAcronymState() []string {
	return []string{
		string(SearchAcronymState_Draft),
		string(SearchAcronymState_Excluded),
		string(SearchAcronymState_Published),
	}
}

func (s *SearchAcronymState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchAcronymState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchAcronymState(input string) (*SearchAcronymState, error) {
	vals := map[string]SearchAcronymState{
		"draft":     SearchAcronymState_Draft,
		"excluded":  SearchAcronymState_Excluded,
		"published": SearchAcronymState_Published,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchAcronymState(input)
	return &out, nil
}
