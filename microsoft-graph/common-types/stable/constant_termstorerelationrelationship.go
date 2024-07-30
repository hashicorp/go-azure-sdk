package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TermStoreRelationRelationship string

const (
	TermStoreRelationRelationship_Pin   TermStoreRelationRelationship = "pin"
	TermStoreRelationRelationship_Reuse TermStoreRelationRelationship = "reuse"
)

func PossibleValuesForTermStoreRelationRelationship() []string {
	return []string{
		string(TermStoreRelationRelationship_Pin),
		string(TermStoreRelationRelationship_Reuse),
	}
}

func (s *TermStoreRelationRelationship) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTermStoreRelationRelationship(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTermStoreRelationRelationship(input string) (*TermStoreRelationRelationship, error) {
	vals := map[string]TermStoreRelationRelationship{
		"pin":   TermStoreRelationRelationship_Pin,
		"reuse": TermStoreRelationRelationship_Reuse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TermStoreRelationRelationship(input)
	return &out, nil
}
