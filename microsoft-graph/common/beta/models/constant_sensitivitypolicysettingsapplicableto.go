package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityPolicySettingsApplicableTo string

const (
	SensitivityPolicySettingsApplicableToemail        SensitivityPolicySettingsApplicableTo = "Email"
	SensitivityPolicySettingsApplicableTosite         SensitivityPolicySettingsApplicableTo = "Site"
	SensitivityPolicySettingsApplicableToteamwork     SensitivityPolicySettingsApplicableTo = "Teamwork"
	SensitivityPolicySettingsApplicableTounifiedGroup SensitivityPolicySettingsApplicableTo = "UnifiedGroup"
)

func PossibleValuesForSensitivityPolicySettingsApplicableTo() []string {
	return []string{
		string(SensitivityPolicySettingsApplicableToemail),
		string(SensitivityPolicySettingsApplicableTosite),
		string(SensitivityPolicySettingsApplicableToteamwork),
		string(SensitivityPolicySettingsApplicableTounifiedGroup),
	}
}

func parseSensitivityPolicySettingsApplicableTo(input string) (*SensitivityPolicySettingsApplicableTo, error) {
	vals := map[string]SensitivityPolicySettingsApplicableTo{
		"email":        SensitivityPolicySettingsApplicableToemail,
		"site":         SensitivityPolicySettingsApplicableTosite,
		"teamwork":     SensitivityPolicySettingsApplicableToteamwork,
		"unifiedgroup": SensitivityPolicySettingsApplicableTounifiedGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityPolicySettingsApplicableTo(input)
	return &out, nil
}
