package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleAssignmentType string

const (
	PrivilegedAccessGroupAssignmentScheduleAssignmentTypeactivated PrivilegedAccessGroupAssignmentScheduleAssignmentType = "Activated"
	PrivilegedAccessGroupAssignmentScheduleAssignmentTypeassigned  PrivilegedAccessGroupAssignmentScheduleAssignmentType = "Assigned"
)

func PossibleValuesForPrivilegedAccessGroupAssignmentScheduleAssignmentType() []string {
	return []string{
		string(PrivilegedAccessGroupAssignmentScheduleAssignmentTypeactivated),
		string(PrivilegedAccessGroupAssignmentScheduleAssignmentTypeassigned),
	}
}

func parsePrivilegedAccessGroupAssignmentScheduleAssignmentType(input string) (*PrivilegedAccessGroupAssignmentScheduleAssignmentType, error) {
	vals := map[string]PrivilegedAccessGroupAssignmentScheduleAssignmentType{
		"activated": PrivilegedAccessGroupAssignmentScheduleAssignmentTypeactivated,
		"assigned":  PrivilegedAccessGroupAssignmentScheduleAssignmentTypeassigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegedAccessGroupAssignmentScheduleAssignmentType(input)
	return &out, nil
}
