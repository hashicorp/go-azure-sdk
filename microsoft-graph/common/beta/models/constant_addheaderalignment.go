package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddHeaderAlignment string

const (
	AddHeaderAlignmentcenter AddHeaderAlignment = "Center"
	AddHeaderAlignmentleft   AddHeaderAlignment = "Left"
	AddHeaderAlignmentright  AddHeaderAlignment = "Right"
)

func PossibleValuesForAddHeaderAlignment() []string {
	return []string{
		string(AddHeaderAlignmentcenter),
		string(AddHeaderAlignmentleft),
		string(AddHeaderAlignmentright),
	}
}

func parseAddHeaderAlignment(input string) (*AddHeaderAlignment, error) {
	vals := map[string]AddHeaderAlignment{
		"center": AddHeaderAlignmentcenter,
		"left":   AddHeaderAlignmentleft,
		"right":  AddHeaderAlignmentright,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddHeaderAlignment(input)
	return &out, nil
}
