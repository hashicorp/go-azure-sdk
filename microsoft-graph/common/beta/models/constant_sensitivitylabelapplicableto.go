package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelApplicableTo string

const (
	SensitivityLabelApplicableToemail        SensitivityLabelApplicableTo = "Email"
	SensitivityLabelApplicableTosite         SensitivityLabelApplicableTo = "Site"
	SensitivityLabelApplicableToteamwork     SensitivityLabelApplicableTo = "Teamwork"
	SensitivityLabelApplicableTounifiedGroup SensitivityLabelApplicableTo = "UnifiedGroup"
)

func PossibleValuesForSensitivityLabelApplicableTo() []string {
	return []string{
		string(SensitivityLabelApplicableToemail),
		string(SensitivityLabelApplicableTosite),
		string(SensitivityLabelApplicableToteamwork),
		string(SensitivityLabelApplicableTounifiedGroup),
	}
}

func parseSensitivityLabelApplicableTo(input string) (*SensitivityLabelApplicableTo, error) {
	vals := map[string]SensitivityLabelApplicableTo{
		"email":        SensitivityLabelApplicableToemail,
		"site":         SensitivityLabelApplicableTosite,
		"teamwork":     SensitivityLabelApplicableToteamwork,
		"unifiedgroup": SensitivityLabelApplicableTounifiedGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelApplicableTo(input)
	return &out, nil
}
