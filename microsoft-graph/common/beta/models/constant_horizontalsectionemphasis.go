package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HorizontalSectionEmphasis string

const (
	HorizontalSectionEmphasisneutral HorizontalSectionEmphasis = "Neutral"
	HorizontalSectionEmphasisnone    HorizontalSectionEmphasis = "None"
	HorizontalSectionEmphasissoft    HorizontalSectionEmphasis = "Soft"
	HorizontalSectionEmphasisstrong  HorizontalSectionEmphasis = "Strong"
)

func PossibleValuesForHorizontalSectionEmphasis() []string {
	return []string{
		string(HorizontalSectionEmphasisneutral),
		string(HorizontalSectionEmphasisnone),
		string(HorizontalSectionEmphasissoft),
		string(HorizontalSectionEmphasisstrong),
	}
}

func parseHorizontalSectionEmphasis(input string) (*HorizontalSectionEmphasis, error) {
	vals := map[string]HorizontalSectionEmphasis{
		"neutral": HorizontalSectionEmphasisneutral,
		"none":    HorizontalSectionEmphasisnone,
		"soft":    HorizontalSectionEmphasissoft,
		"strong":  HorizontalSectionEmphasisstrong,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HorizontalSectionEmphasis(input)
	return &out, nil
}
