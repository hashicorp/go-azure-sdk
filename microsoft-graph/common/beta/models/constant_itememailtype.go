package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemEmailType string

const (
	ItemEmailTypemain     ItemEmailType = "Main"
	ItemEmailTypeother    ItemEmailType = "Other"
	ItemEmailTypepersonal ItemEmailType = "Personal"
	ItemEmailTypeunknown  ItemEmailType = "Unknown"
	ItemEmailTypework     ItemEmailType = "Work"
)

func PossibleValuesForItemEmailType() []string {
	return []string{
		string(ItemEmailTypemain),
		string(ItemEmailTypeother),
		string(ItemEmailTypepersonal),
		string(ItemEmailTypeunknown),
		string(ItemEmailTypework),
	}
}

func parseItemEmailType(input string) (*ItemEmailType, error) {
	vals := map[string]ItemEmailType{
		"main":     ItemEmailTypemain,
		"other":    ItemEmailTypeother,
		"personal": ItemEmailTypepersonal,
		"unknown":  ItemEmailTypeunknown,
		"work":     ItemEmailTypework,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemEmailType(input)
	return &out, nil
}
