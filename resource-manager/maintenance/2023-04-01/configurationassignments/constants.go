package configurationassignments

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagOperators string

const (
	TagOperatorsAll TagOperators = "All"
	TagOperatorsAny TagOperators = "Any"
)

func PossibleValuesForTagOperators() []string {
	return []string{
		string(TagOperatorsAll),
		string(TagOperatorsAny),
	}
}

func parseTagOperators(input string) (*TagOperators, error) {
	vals := map[string]TagOperators{
		"all": TagOperatorsAll,
		"any": TagOperatorsAny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TagOperators(input)
	return &out, nil
}
