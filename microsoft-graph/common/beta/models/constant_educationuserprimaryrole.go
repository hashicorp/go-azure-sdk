package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationUserPrimaryRole string

const (
	EducationUserPrimaryRolefaculty EducationUserPrimaryRole = "Faculty"
	EducationUserPrimaryRolenone    EducationUserPrimaryRole = "None"
	EducationUserPrimaryRolestudent EducationUserPrimaryRole = "Student"
	EducationUserPrimaryRoleteacher EducationUserPrimaryRole = "Teacher"
)

func PossibleValuesForEducationUserPrimaryRole() []string {
	return []string{
		string(EducationUserPrimaryRolefaculty),
		string(EducationUserPrimaryRolenone),
		string(EducationUserPrimaryRolestudent),
		string(EducationUserPrimaryRoleteacher),
	}
}

func parseEducationUserPrimaryRole(input string) (*EducationUserPrimaryRole, error) {
	vals := map[string]EducationUserPrimaryRole{
		"faculty": EducationUserPrimaryRolefaculty,
		"none":    EducationUserPrimaryRolenone,
		"student": EducationUserPrimaryRolestudent,
		"teacher": EducationUserPrimaryRoleteacher,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationUserPrimaryRole(input)
	return &out, nil
}
