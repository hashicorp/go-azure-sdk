package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendeeType string

const (
	AttendeeTypeoptional AttendeeType = "Optional"
	AttendeeTyperequired AttendeeType = "Required"
	AttendeeTyperesource AttendeeType = "Resource"
)

func PossibleValuesForAttendeeType() []string {
	return []string{
		string(AttendeeTypeoptional),
		string(AttendeeTyperequired),
		string(AttendeeTyperesource),
	}
}

func parseAttendeeType(input string) (*AttendeeType, error) {
	vals := map[string]AttendeeType{
		"optional": AttendeeTypeoptional,
		"required": AttendeeTyperequired,
		"resource": AttendeeTyperesource,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttendeeType(input)
	return &out, nil
}
