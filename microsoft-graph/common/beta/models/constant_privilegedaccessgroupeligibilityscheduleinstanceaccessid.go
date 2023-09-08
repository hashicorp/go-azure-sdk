package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstanceAccessId string

const (
	PrivilegedAccessGroupEligibilityScheduleInstanceAccessIdmember PrivilegedAccessGroupEligibilityScheduleInstanceAccessId = "Member"
	PrivilegedAccessGroupEligibilityScheduleInstanceAccessIdowner  PrivilegedAccessGroupEligibilityScheduleInstanceAccessId = "Owner"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleInstanceAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleInstanceAccessIdmember),
		string(PrivilegedAccessGroupEligibilityScheduleInstanceAccessIdowner),
	}
}

func parsePrivilegedAccessGroupEligibilityScheduleInstanceAccessId(input string) (*PrivilegedAccessGroupEligibilityScheduleInstanceAccessId, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleInstanceAccessId{
		"member": PrivilegedAccessGroupEligibilityScheduleInstanceAccessIdmember,
		"owner":  PrivilegedAccessGroupEligibilityScheduleInstanceAccessIdowner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleInstanceAccessId(input)
	return &out, nil
}
