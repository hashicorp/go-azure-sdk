package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleMemberType string

const (
	PrivilegedAccessGroupEligibilityScheduleMemberTypedirect PrivilegedAccessGroupEligibilityScheduleMemberType = "Direct"
	PrivilegedAccessGroupEligibilityScheduleMemberTypegroup  PrivilegedAccessGroupEligibilityScheduleMemberType = "Group"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleMemberTypedirect),
		string(PrivilegedAccessGroupEligibilityScheduleMemberTypegroup),
	}
}

func parsePrivilegedAccessGroupEligibilityScheduleMemberType(input string) (*PrivilegedAccessGroupEligibilityScheduleMemberType, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleMemberType{
		"direct": PrivilegedAccessGroupEligibilityScheduleMemberTypedirect,
		"group":  PrivilegedAccessGroupEligibilityScheduleMemberTypegroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleMemberType(input)
	return &out, nil
}
