package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HorizontalSectionLayout string

const (
	HorizontalSectionLayout_FullWidth           HorizontalSectionLayout = "fullWidth"
	HorizontalSectionLayout_None                HorizontalSectionLayout = "none"
	HorizontalSectionLayout_OneColumn           HorizontalSectionLayout = "oneColumn"
	HorizontalSectionLayout_OneThirdLeftColumn  HorizontalSectionLayout = "oneThirdLeftColumn"
	HorizontalSectionLayout_OneThirdRightColumn HorizontalSectionLayout = "oneThirdRightColumn"
	HorizontalSectionLayout_ThreeColumns        HorizontalSectionLayout = "threeColumns"
	HorizontalSectionLayout_TwoColumns          HorizontalSectionLayout = "twoColumns"
)

func PossibleValuesForHorizontalSectionLayout() []string {
	return []string{
		string(HorizontalSectionLayout_FullWidth),
		string(HorizontalSectionLayout_None),
		string(HorizontalSectionLayout_OneColumn),
		string(HorizontalSectionLayout_OneThirdLeftColumn),
		string(HorizontalSectionLayout_OneThirdRightColumn),
		string(HorizontalSectionLayout_ThreeColumns),
		string(HorizontalSectionLayout_TwoColumns),
	}
}

func (s *HorizontalSectionLayout) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseHorizontalSectionLayout(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseHorizontalSectionLayout(input string) (*HorizontalSectionLayout, error) {
	vals := map[string]HorizontalSectionLayout{
		"fullwidth":           HorizontalSectionLayout_FullWidth,
		"none":                HorizontalSectionLayout_None,
		"onecolumn":           HorizontalSectionLayout_OneColumn,
		"onethirdleftcolumn":  HorizontalSectionLayout_OneThirdLeftColumn,
		"onethirdrightcolumn": HorizontalSectionLayout_OneThirdRightColumn,
		"threecolumns":        HorizontalSectionLayout_ThreeColumns,
		"twocolumns":          HorizontalSectionLayout_TwoColumns,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HorizontalSectionLayout(input)
	return &out, nil
}
