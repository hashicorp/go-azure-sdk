package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentDefaultsAddedStudentAction string

const (
	EducationAssignmentDefaultsAddedStudentActionassignIfOpen EducationAssignmentDefaultsAddedStudentAction = "AssignIfOpen"
	EducationAssignmentDefaultsAddedStudentActionnone         EducationAssignmentDefaultsAddedStudentAction = "None"
)

func PossibleValuesForEducationAssignmentDefaultsAddedStudentAction() []string {
	return []string{
		string(EducationAssignmentDefaultsAddedStudentActionassignIfOpen),
		string(EducationAssignmentDefaultsAddedStudentActionnone),
	}
}

func parseEducationAssignmentDefaultsAddedStudentAction(input string) (*EducationAssignmentDefaultsAddedStudentAction, error) {
	vals := map[string]EducationAssignmentDefaultsAddedStudentAction{
		"assignifopen": EducationAssignmentDefaultsAddedStudentActionassignIfOpen,
		"none":         EducationAssignmentDefaultsAddedStudentActionnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentDefaultsAddedStudentAction(input)
	return &out, nil
}
