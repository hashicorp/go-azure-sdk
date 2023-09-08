package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSchoolExternalSource string

const (
	EducationSchoolExternalSourcelms    EducationSchoolExternalSource = "Lms"
	EducationSchoolExternalSourcemanual EducationSchoolExternalSource = "Manual"
	EducationSchoolExternalSourcesis    EducationSchoolExternalSource = "Sis"
)

func PossibleValuesForEducationSchoolExternalSource() []string {
	return []string{
		string(EducationSchoolExternalSourcelms),
		string(EducationSchoolExternalSourcemanual),
		string(EducationSchoolExternalSourcesis),
	}
}

func parseEducationSchoolExternalSource(input string) (*EducationSchoolExternalSource, error) {
	vals := map[string]EducationSchoolExternalSource{
		"lms":    EducationSchoolExternalSourcelms,
		"manual": EducationSchoolExternalSourcemanual,
		"sis":    EducationSchoolExternalSourcesis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationSchoolExternalSource(input)
	return &out, nil
}
