package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeBaseType string

const (
	AttendeeBaseTypeoptional AttendeeBaseType = "Optional"
	AttendeeBaseTyperequired AttendeeBaseType = "Required"
	AttendeeBaseTyperesource AttendeeBaseType = "Resource"
)

func PossibleValuesForAttendeeBaseType() []string {
	return []string{
		string(AttendeeBaseTypeoptional),
		string(AttendeeBaseTyperequired),
		string(AttendeeBaseTyperesource),
	}
}

func parseAttendeeBaseType(input string) (*AttendeeBaseType, error) {
	vals := map[string]AttendeeBaseType{
		"optional": AttendeeBaseTypeoptional,
		"required": AttendeeBaseTyperequired,
		"resource": AttendeeBaseTyperesource,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttendeeBaseType(input)
	return &out, nil
}
