package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentDefaultsAddToCalendarAction string

const (
	EducationAssignmentDefaultsAddToCalendarActionnone                  EducationAssignmentDefaultsAddToCalendarAction = "None"
	EducationAssignmentDefaultsAddToCalendarActionstudentsAndPublisher  EducationAssignmentDefaultsAddToCalendarAction = "StudentsAndPublisher"
	EducationAssignmentDefaultsAddToCalendarActionstudentsAndTeamOwners EducationAssignmentDefaultsAddToCalendarAction = "StudentsAndTeamOwners"
	EducationAssignmentDefaultsAddToCalendarActionstudentsOnly          EducationAssignmentDefaultsAddToCalendarAction = "StudentsOnly"
)

func PossibleValuesForEducationAssignmentDefaultsAddToCalendarAction() []string {
	return []string{
		string(EducationAssignmentDefaultsAddToCalendarActionnone),
		string(EducationAssignmentDefaultsAddToCalendarActionstudentsAndPublisher),
		string(EducationAssignmentDefaultsAddToCalendarActionstudentsAndTeamOwners),
		string(EducationAssignmentDefaultsAddToCalendarActionstudentsOnly),
	}
}

func parseEducationAssignmentDefaultsAddToCalendarAction(input string) (*EducationAssignmentDefaultsAddToCalendarAction, error) {
	vals := map[string]EducationAssignmentDefaultsAddToCalendarAction{
		"none":                  EducationAssignmentDefaultsAddToCalendarActionnone,
		"studentsandpublisher":  EducationAssignmentDefaultsAddToCalendarActionstudentsAndPublisher,
		"studentsandteamowners": EducationAssignmentDefaultsAddToCalendarActionstudentsAndTeamOwners,
		"studentsonly":          EducationAssignmentDefaultsAddToCalendarActionstudentsOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationAssignmentDefaultsAddToCalendarAction(input)
	return &out, nil
}
