package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentStatus string

const (
	EducationAssignmentStatusassigned  EducationAssignmentStatus = "Assigned"
	EducationAssignmentStatusdraft     EducationAssignmentStatus = "Draft"
	EducationAssignmentStatusinactive  EducationAssignmentStatus = "Inactive"
	EducationAssignmentStatuspublished EducationAssignmentStatus = "Published"
)

func PossibleValuesForEducationAssignmentStatus() []string {
	return []string{
		string(EducationAssignmentStatusassigned),
		string(EducationAssignmentStatusdraft),
		string(EducationAssignmentStatusinactive),
		string(EducationAssignmentStatuspublished),
	}
}

func parseEducationAssignmentStatus(input string) (*EducationAssignmentStatus, error) {
	vals := map[string]EducationAssignmentStatus{
		"assigned":  EducationAssignmentStatusassigned,
		"draft":     EducationAssignmentStatusdraft,
		"inactive":  EducationAssignmentStatusinactive,
		"published": EducationAssignmentStatuspublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentStatus(input)
	return &out, nil
}
