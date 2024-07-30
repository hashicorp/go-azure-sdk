package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HorizontalSectionEmphasis string

const (
	HorizontalSectionEmphasis_Neutral HorizontalSectionEmphasis = "neutral"
	HorizontalSectionEmphasis_None    HorizontalSectionEmphasis = "none"
	HorizontalSectionEmphasis_Soft    HorizontalSectionEmphasis = "soft"
	HorizontalSectionEmphasis_Strong  HorizontalSectionEmphasis = "strong"
)

func PossibleValuesForHorizontalSectionEmphasis() []string {
	return []string{
		string(HorizontalSectionEmphasis_Neutral),
		string(HorizontalSectionEmphasis_None),
		string(HorizontalSectionEmphasis_Soft),
		string(HorizontalSectionEmphasis_Strong),
	}
}

func (s *HorizontalSectionEmphasis) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHorizontalSectionEmphasis(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHorizontalSectionEmphasis(input string) (*HorizontalSectionEmphasis, error) {
	vals := map[string]HorizontalSectionEmphasis{
		"neutral": HorizontalSectionEmphasis_Neutral,
		"none":    HorizontalSectionEmphasis_None,
		"soft":    HorizontalSectionEmphasis_Soft,
		"strong":  HorizontalSectionEmphasis_Strong,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HorizontalSectionEmphasis(input)
	return &out, nil
}
