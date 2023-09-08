package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeConstraintActivityDomain string

const (
	TimeConstraintActivityDomainpersonal     TimeConstraintActivityDomain = "Personal"
	TimeConstraintActivityDomainunknown      TimeConstraintActivityDomain = "Unknown"
	TimeConstraintActivityDomainunrestricted TimeConstraintActivityDomain = "Unrestricted"
	TimeConstraintActivityDomainwork         TimeConstraintActivityDomain = "Work"
)

func PossibleValuesForTimeConstraintActivityDomain() []string {
	return []string{
		string(TimeConstraintActivityDomainpersonal),
		string(TimeConstraintActivityDomainunknown),
		string(TimeConstraintActivityDomainunrestricted),
		string(TimeConstraintActivityDomainwork),
	}
}

func parseTimeConstraintActivityDomain(input string) (*TimeConstraintActivityDomain, error) {
	vals := map[string]TimeConstraintActivityDomain{
		"personal":     TimeConstraintActivityDomainpersonal,
		"unknown":      TimeConstraintActivityDomainunknown,
		"unrestricted": TimeConstraintActivityDomainunrestricted,
		"work":         TimeConstraintActivityDomainwork,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TimeConstraintActivityDomain(input)
	return &out, nil
}
