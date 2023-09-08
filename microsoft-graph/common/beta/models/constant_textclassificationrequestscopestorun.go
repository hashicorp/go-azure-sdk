package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextClassificationRequestScopesToRun string

const (
	TextClassificationRequestScopesToRunfullDocument    TextClassificationRequestScopesToRun = "FullDocument"
	TextClassificationRequestScopesToRunpartialDocument TextClassificationRequestScopesToRun = "PartialDocument"
)

func PossibleValuesForTextClassificationRequestScopesToRun() []string {
	return []string{
		string(TextClassificationRequestScopesToRunfullDocument),
		string(TextClassificationRequestScopesToRunpartialDocument),
	}
}

func parseTextClassificationRequestScopesToRun(input string) (*TextClassificationRequestScopesToRun, error) {
	vals := map[string]TextClassificationRequestScopesToRun{
		"fulldocument":    TextClassificationRequestScopesToRunfullDocument,
		"partialdocument": TextClassificationRequestScopesToRunpartialDocument,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TextClassificationRequestScopesToRun(input)
	return &out, nil
}
