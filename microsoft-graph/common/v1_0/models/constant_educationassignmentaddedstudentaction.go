package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentAddedStudentAction string

const (
	EducationAssignmentAddedStudentActionassignIfOpen EducationAssignmentAddedStudentAction = "AssignIfOpen"
	EducationAssignmentAddedStudentActionnone         EducationAssignmentAddedStudentAction = "None"
)

func PossibleValuesForEducationAssignmentAddedStudentAction() []string {
	return []string{
		string(EducationAssignmentAddedStudentActionassignIfOpen),
		string(EducationAssignmentAddedStudentActionnone),
	}
}

func parseEducationAssignmentAddedStudentAction(input string) (*EducationAssignmentAddedStudentAction, error) {
	vals := map[string]EducationAssignmentAddedStudentAction{
		"assignifopen": EducationAssignmentAddedStudentActionassignIfOpen,
		"none":         EducationAssignmentAddedStudentActionnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentAddedStudentAction(input)
	return &out, nil
}
