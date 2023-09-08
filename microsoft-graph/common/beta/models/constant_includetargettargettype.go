package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncludeTargetTargetType string

const (
	IncludeTargetTargetTypegroup IncludeTargetTargetType = "Group"
	IncludeTargetTargetTypeuser  IncludeTargetTargetType = "User"
)

func PossibleValuesForIncludeTargetTargetType() []string {
	return []string{
		string(IncludeTargetTargetTypegroup),
		string(IncludeTargetTargetTypeuser),
	}
}

func parseIncludeTargetTargetType(input string) (*IncludeTargetTargetType, error) {
	vals := map[string]IncludeTargetTargetType{
		"group": IncludeTargetTargetTypegroup,
		"user":  IncludeTargetTargetTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IncludeTargetTargetType(input)
	return &out, nil
}
