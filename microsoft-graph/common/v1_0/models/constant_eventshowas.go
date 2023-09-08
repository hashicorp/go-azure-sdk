package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventShowAs string

const (
	EventShowAsbusy             EventShowAs = "Busy"
	EventShowAsfree             EventShowAs = "Free"
	EventShowAsoof              EventShowAs = "Oof"
	EventShowAstentative        EventShowAs = "Tentative"
	EventShowAsunknown          EventShowAs = "Unknown"
	EventShowAsworkingElsewhere EventShowAs = "WorkingElsewhere"
)

func PossibleValuesForEventShowAs() []string {
	return []string{
		string(EventShowAsbusy),
		string(EventShowAsfree),
		string(EventShowAsoof),
		string(EventShowAstentative),
		string(EventShowAsunknown),
		string(EventShowAsworkingElsewhere),
	}
}

func parseEventShowAs(input string) (*EventShowAs, error) {
	vals := map[string]EventShowAs{
		"busy":             EventShowAsbusy,
		"free":             EventShowAsfree,
		"oof":              EventShowAsoof,
		"tentative":        EventShowAstentative,
		"unknown":          EventShowAsunknown,
		"workingelsewhere": EventShowAsworkingElsewhere,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EventShowAs(input)
	return &out, nil
}
