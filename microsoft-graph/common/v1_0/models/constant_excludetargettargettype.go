package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExcludeTargetTargetType string

const (
	ExcludeTargetTargetTypegroup ExcludeTargetTargetType = "Group"
	ExcludeTargetTargetTypeuser  ExcludeTargetTargetType = "User"
)

func PossibleValuesForExcludeTargetTargetType() []string {
	return []string{
		string(ExcludeTargetTargetTypegroup),
		string(ExcludeTargetTargetTypeuser),
	}
}

func parseExcludeTargetTargetType(input string) (*ExcludeTargetTargetType, error) {
	vals := map[string]ExcludeTargetTargetType{
		"group": ExcludeTargetTargetTypegroup,
		"user":  ExcludeTargetTargetTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExcludeTargetTargetType(input)
	return &out, nil
}
