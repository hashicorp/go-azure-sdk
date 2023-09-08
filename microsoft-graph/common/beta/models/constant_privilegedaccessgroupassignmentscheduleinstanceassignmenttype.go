package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType string

const (
	PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentTypeactivated PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType = "Activated"
	PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentTypeassigned  PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType = "Assigned"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentTypeactivated),
		string(PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentTypeassigned),
	}
}

func parsePrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType(input string) (*PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType{
		"activated": PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentTypeactivated,
		"assigned":  PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentTypeassigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleInstanceAssignmentType(input)
	return &out, nil
}
