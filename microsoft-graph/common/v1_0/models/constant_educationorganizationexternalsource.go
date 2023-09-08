package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationOrganizationExternalSource string

const (
	EducationOrganizationExternalSourcemanual EducationOrganizationExternalSource = "Manual"
	EducationOrganizationExternalSourcesis    EducationOrganizationExternalSource = "Sis"
)

func PossibleValuesForEducationOrganizationExternalSource() []string {
	return []string{
		string(EducationOrganizationExternalSourcemanual),
		string(EducationOrganizationExternalSourcesis),
	}
}

func parseEducationOrganizationExternalSource(input string) (*EducationOrganizationExternalSource, error) {
	vals := map[string]EducationOrganizationExternalSource{
		"manual": EducationOrganizationExternalSourcemanual,
		"sis":    EducationOrganizationExternalSourcesis,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationOrganizationExternalSource(input)
	return &out, nil
}
