package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationStudentGender string

const (
	EducationStudentGenderfemale EducationStudentGender = "Female"
	EducationStudentGendermale   EducationStudentGender = "Male"
	EducationStudentGenderother  EducationStudentGender = "Other"
)

func PossibleValuesForEducationStudentGender() []string {
	return []string{
		string(EducationStudentGenderfemale),
		string(EducationStudentGendermale),
		string(EducationStudentGenderother),
	}
}

func parseEducationStudentGender(input string) (*EducationStudentGender, error) {
	vals := map[string]EducationStudentGender{
		"female": EducationStudentGenderfemale,
		"male":   EducationStudentGendermale,
		"other":  EducationStudentGenderother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationStudentGender(input)
	return &out, nil
}
