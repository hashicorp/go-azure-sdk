package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAddContentFooterActionAlignment string

const (
	SecurityAddContentFooterActionAlignmentcenter SecurityAddContentFooterActionAlignment = "Center"
	SecurityAddContentFooterActionAlignmentleft   SecurityAddContentFooterActionAlignment = "Left"
	SecurityAddContentFooterActionAlignmentright  SecurityAddContentFooterActionAlignment = "Right"
)

func PossibleValuesForSecurityAddContentFooterActionAlignment() []string {
	return []string{
		string(SecurityAddContentFooterActionAlignmentcenter),
		string(SecurityAddContentFooterActionAlignmentleft),
		string(SecurityAddContentFooterActionAlignmentright),
	}
}

func parseSecurityAddContentFooterActionAlignment(input string) (*SecurityAddContentFooterActionAlignment, error) {
	vals := map[string]SecurityAddContentFooterActionAlignment{
		"center": SecurityAddContentFooterActionAlignmentcenter,
		"left":   SecurityAddContentFooterActionAlignmentleft,
		"right":  SecurityAddContentFooterActionAlignmentright,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAddContentFooterActionAlignment(input)
	return &out, nil
}
