package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TitleAreaTextAlignment string

const (
	TitleAreaTextAlignmentcenter TitleAreaTextAlignment = "Center"
	TitleAreaTextAlignmentleft   TitleAreaTextAlignment = "Left"
)

func PossibleValuesForTitleAreaTextAlignment() []string {
	return []string{
		string(TitleAreaTextAlignmentcenter),
		string(TitleAreaTextAlignmentleft),
	}
}

func parseTitleAreaTextAlignment(input string) (*TitleAreaTextAlignment, error) {
	vals := map[string]TitleAreaTextAlignment{
		"center": TitleAreaTextAlignmentcenter,
		"left":   TitleAreaTextAlignmentleft,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TitleAreaTextAlignment(input)
	return &out, nil
}
