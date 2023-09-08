package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddFooterAlignment string

const (
	AddFooterAlignmentcenter AddFooterAlignment = "Center"
	AddFooterAlignmentleft   AddFooterAlignment = "Left"
	AddFooterAlignmentright  AddFooterAlignment = "Right"
)

func PossibleValuesForAddFooterAlignment() []string {
	return []string{
		string(AddFooterAlignmentcenter),
		string(AddFooterAlignmentleft),
		string(AddFooterAlignmentright),
	}
}

func parseAddFooterAlignment(input string) (*AddFooterAlignment, error) {
	vals := map[string]AddFooterAlignment{
		"center": AddFooterAlignmentcenter,
		"left":   AddFooterAlignmentleft,
		"right":  AddFooterAlignmentright,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddFooterAlignment(input)
	return &out, nil
}
