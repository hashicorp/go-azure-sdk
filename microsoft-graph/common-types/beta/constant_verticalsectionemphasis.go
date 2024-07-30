package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerticalSectionEmphasis string

const (
	VerticalSectionEmphasis_Neutral VerticalSectionEmphasis = "neutral"
	VerticalSectionEmphasis_None    VerticalSectionEmphasis = "none"
	VerticalSectionEmphasis_Soft    VerticalSectionEmphasis = "soft"
	VerticalSectionEmphasis_Strong  VerticalSectionEmphasis = "strong"
)

func PossibleValuesForVerticalSectionEmphasis() []string {
	return []string{
		string(VerticalSectionEmphasis_Neutral),
		string(VerticalSectionEmphasis_None),
		string(VerticalSectionEmphasis_Soft),
		string(VerticalSectionEmphasis_Strong),
	}
}

func (s *VerticalSectionEmphasis) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVerticalSectionEmphasis(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVerticalSectionEmphasis(input string) (*VerticalSectionEmphasis, error) {
	vals := map[string]VerticalSectionEmphasis{
		"neutral": VerticalSectionEmphasis_Neutral,
		"none":    VerticalSectionEmphasis_None,
		"soft":    VerticalSectionEmphasis_Soft,
		"strong":  VerticalSectionEmphasis_Strong,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VerticalSectionEmphasis(input)
	return &out, nil
}
