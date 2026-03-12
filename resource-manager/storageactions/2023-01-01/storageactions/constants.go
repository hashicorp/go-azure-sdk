package storageactions

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchedBlockName string

const (
	MatchedBlockNameElse MatchedBlockName = "Else"
	MatchedBlockNameIf   MatchedBlockName = "If"
	MatchedBlockNameNone MatchedBlockName = "None"
)

func PossibleValuesForMatchedBlockName() []string {
	return []string{
		string(MatchedBlockNameElse),
		string(MatchedBlockNameIf),
		string(MatchedBlockNameNone),
	}
}

func (s *MatchedBlockName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMatchedBlockName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMatchedBlockName(input string) (*MatchedBlockName, error) {
	vals := map[string]MatchedBlockName{
		"else": MatchedBlockNameElse,
		"if":   MatchedBlockNameIf,
		"none": MatchedBlockNameNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MatchedBlockName(input)
	return &out, nil
}
