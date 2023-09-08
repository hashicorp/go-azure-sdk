package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationClassExternalSource string

const (
	EducationClassExternalSourcemanual EducationClassExternalSource = "Manual"
	EducationClassExternalSourcesis    EducationClassExternalSource = "Sis"
)

func PossibleValuesForEducationClassExternalSource() []string {
	return []string{
		string(EducationClassExternalSourcemanual),
		string(EducationClassExternalSourcesis),
	}
}

func parseEducationClassExternalSource(input string) (*EducationClassExternalSource, error) {
	vals := map[string]EducationClassExternalSource{
		"manual": EducationClassExternalSourcemanual,
		"sis":    EducationClassExternalSourcesis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationClassExternalSource(input)
	return &out, nil
}
