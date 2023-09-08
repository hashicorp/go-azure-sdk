package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TypedEmailAddressType string

const (
	TypedEmailAddressTypemain     TypedEmailAddressType = "Main"
	TypedEmailAddressTypeother    TypedEmailAddressType = "Other"
	TypedEmailAddressTypepersonal TypedEmailAddressType = "Personal"
	TypedEmailAddressTypeunknown  TypedEmailAddressType = "Unknown"
	TypedEmailAddressTypework     TypedEmailAddressType = "Work"
)

func PossibleValuesForTypedEmailAddressType() []string {
	return []string{
		string(TypedEmailAddressTypemain),
		string(TypedEmailAddressTypeother),
		string(TypedEmailAddressTypepersonal),
		string(TypedEmailAddressTypeunknown),
		string(TypedEmailAddressTypework),
	}
}

func parseTypedEmailAddressType(input string) (*TypedEmailAddressType, error) {
	vals := map[string]TypedEmailAddressType{
		"main":     TypedEmailAddressTypemain,
		"other":    TypedEmailAddressTypeother,
		"personal": TypedEmailAddressTypepersonal,
		"unknown":  TypedEmailAddressTypeunknown,
		"work":     TypedEmailAddressTypework,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TypedEmailAddressType(input)
	return &out, nil
}
