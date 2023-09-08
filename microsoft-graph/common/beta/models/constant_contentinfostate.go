package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentInfoState string

const (
	ContentInfoStatemotion ContentInfoState = "Motion"
	ContentInfoStaterest   ContentInfoState = "Rest"
	ContentInfoStateuse    ContentInfoState = "Use"
)

func PossibleValuesForContentInfoState() []string {
	return []string{
		string(ContentInfoStatemotion),
		string(ContentInfoStaterest),
		string(ContentInfoStateuse),
	}
}

func parseContentInfoState(input string) (*ContentInfoState, error) {
	vals := map[string]ContentInfoState{
		"motion": ContentInfoStatemotion,
		"rest":   ContentInfoStaterest,
		"use":    ContentInfoStateuse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ContentInfoState(input)
	return &out, nil
}
