package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationUserExternalSource string

const (
	EducationUserExternalSourcemanual EducationUserExternalSource = "Manual"
	EducationUserExternalSourcesis    EducationUserExternalSource = "Sis"
)

func PossibleValuesForEducationUserExternalSource() []string {
	return []string{
		string(EducationUserExternalSourcemanual),
		string(EducationUserExternalSourcesis),
	}
}

func parseEducationUserExternalSource(input string) (*EducationUserExternalSource, error) {
	vals := map[string]EducationUserExternalSource{
		"manual": EducationUserExternalSourcemanual,
		"sis":    EducationUserExternalSourcesis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationUserExternalSource(input)
	return &out, nil
}
