package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAddContentHeaderActionAlignment string

const (
	SecurityAddContentHeaderActionAlignmentcenter SecurityAddContentHeaderActionAlignment = "Center"
	SecurityAddContentHeaderActionAlignmentleft   SecurityAddContentHeaderActionAlignment = "Left"
	SecurityAddContentHeaderActionAlignmentright  SecurityAddContentHeaderActionAlignment = "Right"
)

func PossibleValuesForSecurityAddContentHeaderActionAlignment() []string {
	return []string{
		string(SecurityAddContentHeaderActionAlignmentcenter),
		string(SecurityAddContentHeaderActionAlignmentleft),
		string(SecurityAddContentHeaderActionAlignmentright),
	}
}

func parseSecurityAddContentHeaderActionAlignment(input string) (*SecurityAddContentHeaderActionAlignment, error) {
	vals := map[string]SecurityAddContentHeaderActionAlignment{
		"center": SecurityAddContentHeaderActionAlignmentcenter,
		"left":   SecurityAddContentHeaderActionAlignmentleft,
		"right":  SecurityAddContentHeaderActionAlignmentright,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAddContentHeaderActionAlignment(input)
	return &out, nil
}
