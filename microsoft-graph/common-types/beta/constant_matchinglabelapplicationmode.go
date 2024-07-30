package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MatchingLabelApplicationMode string

const (
	MatchingLabelApplicationMode_Automatic   MatchingLabelApplicationMode = "automatic"
	MatchingLabelApplicationMode_Manual      MatchingLabelApplicationMode = "manual"
	MatchingLabelApplicationMode_Recommended MatchingLabelApplicationMode = "recommended"
)

func PossibleValuesForMatchingLabelApplicationMode() []string {
	return []string{
		string(MatchingLabelApplicationMode_Automatic),
		string(MatchingLabelApplicationMode_Manual),
		string(MatchingLabelApplicationMode_Recommended),
	}
}

func (s *MatchingLabelApplicationMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMatchingLabelApplicationMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMatchingLabelApplicationMode(input string) (*MatchingLabelApplicationMode, error) {
	vals := map[string]MatchingLabelApplicationMode{
		"automatic":   MatchingLabelApplicationMode_Automatic,
		"manual":      MatchingLabelApplicationMode_Manual,
		"recommended": MatchingLabelApplicationMode_Recommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MatchingLabelApplicationMode(input)
	return &out, nil
}
