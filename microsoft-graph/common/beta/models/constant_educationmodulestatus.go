package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationModuleStatus string

const (
	EducationModuleStatusdraft     EducationModuleStatus = "Draft"
	EducationModuleStatuspublished EducationModuleStatus = "Published"
)

func PossibleValuesForEducationModuleStatus() []string {
	return []string{
		string(EducationModuleStatusdraft),
		string(EducationModuleStatuspublished),
	}
}

func parseEducationModuleStatus(input string) (*EducationModuleStatus, error) {
	vals := map[string]EducationModuleStatus{
		"draft":     EducationModuleStatusdraft,
		"published": EducationModuleStatuspublished,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationModuleStatus(input)
	return &out, nil
}
