package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessFilterMode string

const (
	ConditionalAccessFilterModeexclude ConditionalAccessFilterMode = "Exclude"
	ConditionalAccessFilterModeinclude ConditionalAccessFilterMode = "Include"
)

func PossibleValuesForConditionalAccessFilterMode() []string {
	return []string{
		string(ConditionalAccessFilterModeexclude),
		string(ConditionalAccessFilterModeinclude),
	}
}

func parseConditionalAccessFilterMode(input string) (*ConditionalAccessFilterMode, error) {
	vals := map[string]ConditionalAccessFilterMode{
		"exclude": ConditionalAccessFilterModeexclude,
		"include": ConditionalAccessFilterModeinclude,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessFilterMode(input)
	return &out, nil
}
