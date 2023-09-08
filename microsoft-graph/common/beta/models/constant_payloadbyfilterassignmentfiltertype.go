package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadByFilterAssignmentFilterType string

const (
	PayloadByFilterAssignmentFilterTypeexclude PayloadByFilterAssignmentFilterType = "Exclude"
	PayloadByFilterAssignmentFilterTypeinclude PayloadByFilterAssignmentFilterType = "Include"
	PayloadByFilterAssignmentFilterTypenone    PayloadByFilterAssignmentFilterType = "None"
)

func PossibleValuesForPayloadByFilterAssignmentFilterType() []string {
	return []string{
		string(PayloadByFilterAssignmentFilterTypeexclude),
		string(PayloadByFilterAssignmentFilterTypeinclude),
		string(PayloadByFilterAssignmentFilterTypenone),
	}
}

func parsePayloadByFilterAssignmentFilterType(input string) (*PayloadByFilterAssignmentFilterType, error) {
	vals := map[string]PayloadByFilterAssignmentFilterType{
		"exclude": PayloadByFilterAssignmentFilterTypeexclude,
		"include": PayloadByFilterAssignmentFilterTypeinclude,
		"none":    PayloadByFilterAssignmentFilterTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadByFilterAssignmentFilterType(input)
	return &out, nil
}
