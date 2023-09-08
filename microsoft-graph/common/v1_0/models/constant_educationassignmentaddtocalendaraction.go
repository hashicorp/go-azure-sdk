package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentAddToCalendarAction string

const (
	EducationAssignmentAddToCalendarActionnone                  EducationAssignmentAddToCalendarAction = "None"
	EducationAssignmentAddToCalendarActionstudentsAndPublisher  EducationAssignmentAddToCalendarAction = "StudentsAndPublisher"
	EducationAssignmentAddToCalendarActionstudentsAndTeamOwners EducationAssignmentAddToCalendarAction = "StudentsAndTeamOwners"
	EducationAssignmentAddToCalendarActionstudentsOnly          EducationAssignmentAddToCalendarAction = "StudentsOnly"
)

func PossibleValuesForEducationAssignmentAddToCalendarAction() []string {
	return []string{
		string(EducationAssignmentAddToCalendarActionnone),
		string(EducationAssignmentAddToCalendarActionstudentsAndPublisher),
		string(EducationAssignmentAddToCalendarActionstudentsAndTeamOwners),
		string(EducationAssignmentAddToCalendarActionstudentsOnly),
	}
}

func parseEducationAssignmentAddToCalendarAction(input string) (*EducationAssignmentAddToCalendarAction, error) {
	vals := map[string]EducationAssignmentAddToCalendarAction{
		"none":                  EducationAssignmentAddToCalendarActionnone,
		"studentsandpublisher":  EducationAssignmentAddToCalendarActionstudentsAndPublisher,
		"studentsandteamowners": EducationAssignmentAddToCalendarActionstudentsAndTeamOwners,
		"studentsonly":          EducationAssignmentAddToCalendarActionstudentsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentAddToCalendarAction(input)
	return &out, nil
}
