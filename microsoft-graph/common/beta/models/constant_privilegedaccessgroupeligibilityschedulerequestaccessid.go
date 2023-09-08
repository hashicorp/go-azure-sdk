package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleRequestAccessId string

const (
	PrivilegedAccessGroupEligibilityScheduleRequestAccessIdmember PrivilegedAccessGroupEligibilityScheduleRequestAccessId = "Member"
	PrivilegedAccessGroupEligibilityScheduleRequestAccessIdowner  PrivilegedAccessGroupEligibilityScheduleRequestAccessId = "Owner"
)

func PossibleValuesForPrivilegedAccessGroupEligibilityScheduleRequestAccessId() []string {
	return []string{
		string(PrivilegedAccessGroupEligibilityScheduleRequestAccessIdmember),
		string(PrivilegedAccessGroupEligibilityScheduleRequestAccessIdowner),
	}
}

func parsePrivilegedAccessGroupEligibilityScheduleRequestAccessId(input string) (*PrivilegedAccessGroupEligibilityScheduleRequestAccessId, error) {
	vals := map[string]PrivilegedAccessGroupEligibilityScheduleRequestAccessId{
		"member": PrivilegedAccessGroupEligibilityScheduleRequestAccessIdmember,
		"owner":  PrivilegedAccessGroupEligibilityScheduleRequestAccessIdowner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupEligibilityScheduleRequestAccessId(input)
	return &out, nil
}
