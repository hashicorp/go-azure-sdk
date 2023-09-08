package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddContentHeaderActionAlignment string

const (
	AddContentHeaderActionAlignmentcenter AddContentHeaderActionAlignment = "Center"
	AddContentHeaderActionAlignmentleft   AddContentHeaderActionAlignment = "Left"
	AddContentHeaderActionAlignmentright  AddContentHeaderActionAlignment = "Right"
)

func PossibleValuesForAddContentHeaderActionAlignment() []string {
	return []string{
		string(AddContentHeaderActionAlignmentcenter),
		string(AddContentHeaderActionAlignmentleft),
		string(AddContentHeaderActionAlignmentright),
	}
}

func parseAddContentHeaderActionAlignment(input string) (*AddContentHeaderActionAlignment, error) {
	vals := map[string]AddContentHeaderActionAlignment{
		"center": AddContentHeaderActionAlignmentcenter,
		"left":   AddContentHeaderActionAlignmentleft,
		"right":  AddContentHeaderActionAlignmentright,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AddContentHeaderActionAlignment(input)
	return &out, nil
}
