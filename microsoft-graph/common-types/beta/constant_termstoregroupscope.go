package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreGroupScope string

const (
	TermStoreGroupScope_Global         TermStoreGroupScope = "global"
	TermStoreGroupScope_SiteCollection TermStoreGroupScope = "siteCollection"
	TermStoreGroupScope_System         TermStoreGroupScope = "system"
)

func PossibleValuesForTermStoreGroupScope() []string {
	return []string{
		string(TermStoreGroupScope_Global),
		string(TermStoreGroupScope_SiteCollection),
		string(TermStoreGroupScope_System),
	}
}

func (s *TermStoreGroupScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTermStoreGroupScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTermStoreGroupScope(input string) (*TermStoreGroupScope, error) {
	vals := map[string]TermStoreGroupScope{
		"global":         TermStoreGroupScope_Global,
		"sitecollection": TermStoreGroupScope_SiteCollection,
		"system":         TermStoreGroupScope_System,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TermStoreGroupScope(input)
	return &out, nil
}
