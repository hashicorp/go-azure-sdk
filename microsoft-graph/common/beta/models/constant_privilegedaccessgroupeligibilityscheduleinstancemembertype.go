package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstanceMemberType string

const (
	PrivilegedAccessGroupEligibilityScheduleInstanceMemberTypedirect PrivilegedAccessGroupEligibilityScheduleInstanceMemberType = "Direct"
	PrivilegedAccessGroupEligibilityScheduleInstanceMemberTypegroup  PrivilegedAccessGroupEligibilityScheduleInstanceMemberType = "Group"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleInstanceMemberType() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleInstanceMemberTypedirect),
		string(PrivilegedAccessGroupEligibilityScheduleInstanceMemberTypegroup),
	}
}

func parsePrivilegedAccessGroupEligibilityScheduleInstanceMemberType(input string) (*PrivilegedAccessGroupEligibilityScheduleInstanceMemberType, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleInstanceMemberType{
		"direct": PrivilegedAccessGroupEligibilityScheduleInstanceMemberTypedirect,
		"group":  PrivilegedAccessGroupEligibilityScheduleInstanceMemberTypegroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleInstanceMemberType(input)
	return &out, nil
}
