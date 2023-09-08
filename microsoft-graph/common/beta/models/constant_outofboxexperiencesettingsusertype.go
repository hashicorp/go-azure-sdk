package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfBoxExperienceSettingsUserType string

const (
	OutOfBoxExperienceSettingsUserTypeadministrator OutOfBoxExperienceSettingsUserType = "Administrator"
	OutOfBoxExperienceSettingsUserTypestandard      OutOfBoxExperienceSettingsUserType = "Standard"
)

func PossibleValuesForOutOfBoxExperienceSettingsUserType() []string {
	return []string{
		string(OutOfBoxExperienceSettingsUserTypeadministrator),
		string(OutOfBoxExperienceSettingsUserTypestandard),
	}
}

func parseOutOfBoxExperienceSettingsUserType(input string) (*OutOfBoxExperienceSettingsUserType, error) {
	vals := map[string]OutOfBoxExperienceSettingsUserType{
		"administrator": OutOfBoxExperienceSettingsUserTypeadministrator,
		"standard":      OutOfBoxExperienceSettingsUserTypestandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutOfBoxExperienceSettingsUserType(input)
	return &out, nil
}
