package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HorizontalSectionLayout string

const (
	HorizontalSectionLayoutfullWidth           HorizontalSectionLayout = "FullWidth"
	HorizontalSectionLayoutnone                HorizontalSectionLayout = "None"
	HorizontalSectionLayoutoneColumn           HorizontalSectionLayout = "OneColumn"
	HorizontalSectionLayoutoneThirdLeftColumn  HorizontalSectionLayout = "OneThirdLeftColumn"
	HorizontalSectionLayoutoneThirdRightColumn HorizontalSectionLayout = "OneThirdRightColumn"
	HorizontalSectionLayoutthreeColumns        HorizontalSectionLayout = "ThreeColumns"
	HorizontalSectionLayouttwoColumns          HorizontalSectionLayout = "TwoColumns"
)

func PossibleValuesForHorizontalSectionLayout() []string {
	return []string{
		string(HorizontalSectionLayoutfullWidth),
		string(HorizontalSectionLayoutnone),
		string(HorizontalSectionLayoutoneColumn),
		string(HorizontalSectionLayoutoneThirdLeftColumn),
		string(HorizontalSectionLayoutoneThirdRightColumn),
		string(HorizontalSectionLayoutthreeColumns),
		string(HorizontalSectionLayouttwoColumns),
	}
}

func parseHorizontalSectionLayout(input string) (*HorizontalSectionLayout, error) {
	vals := map[string]HorizontalSectionLayout{
		"fullwidth":           HorizontalSectionLayoutfullWidth,
		"none":                HorizontalSectionLayoutnone,
		"onecolumn":           HorizontalSectionLayoutoneColumn,
		"onethirdleftcolumn":  HorizontalSectionLayoutoneThirdLeftColumn,
		"onethirdrightcolumn": HorizontalSectionLayoutoneThirdRightColumn,
		"threecolumns":        HorizontalSectionLayoutthreeColumns,
		"twocolumns":          HorizontalSectionLayouttwoColumns,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HorizontalSectionLayout(input)
	return &out, nil
}
