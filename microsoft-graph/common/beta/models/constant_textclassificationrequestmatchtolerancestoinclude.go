package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextClassificationRequestMatchTolerancesToInclude string

const (
	TextClassificationRequestMatchTolerancesToIncludeexact TextClassificationRequestMatchTolerancesToInclude = "Exact"
	TextClassificationRequestMatchTolerancesToIncludenear  TextClassificationRequestMatchTolerancesToInclude = "Near"
)

func PossibleValuesForTextClassificationRequestMatchTolerancesToInclude() []string {
	return []string{
		string(TextClassificationRequestMatchTolerancesToIncludeexact),
		string(TextClassificationRequestMatchTolerancesToIncludenear),
	}
}

func parseTextClassificationRequestMatchTolerancesToInclude(input string) (*TextClassificationRequestMatchTolerancesToInclude, error) {
	vals := map[string]TextClassificationRequestMatchTolerancesToInclude{
		"exact": TextClassificationRequestMatchTolerancesToIncludeexact,
		"near":  TextClassificationRequestMatchTolerancesToIncludenear,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TextClassificationRequestMatchTolerancesToInclude(input)
	return &out, nil
}
