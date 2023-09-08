package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkillProficiencyProficiency string

const (
	SkillProficiencyProficiencyadvancedProfessional SkillProficiencyProficiency = "AdvancedProfessional"
	SkillProficiencyProficiencyelementary           SkillProficiencyProficiency = "Elementary"
	SkillProficiencyProficiencyexpert               SkillProficiencyProficiency = "Expert"
	SkillProficiencyProficiencygeneralProfessional  SkillProficiencyProficiency = "GeneralProfessional"
	SkillProficiencyProficiencylimitedWorking       SkillProficiencyProficiency = "LimitedWorking"
)

func PossibleValuesForSkillProficiencyProficiency() []string {
	return []string{
		string(SkillProficiencyProficiencyadvancedProfessional),
		string(SkillProficiencyProficiencyelementary),
		string(SkillProficiencyProficiencyexpert),
		string(SkillProficiencyProficiencygeneralProfessional),
		string(SkillProficiencyProficiencylimitedWorking),
	}
}

func parseSkillProficiencyProficiency(input string) (*SkillProficiencyProficiency, error) {
	vals := map[string]SkillProficiencyProficiency{
		"advancedprofessional": SkillProficiencyProficiencyadvancedProfessional,
		"elementary":           SkillProficiencyProficiencyelementary,
		"expert":               SkillProficiencyProficiencyexpert,
		"generalprofessional":  SkillProficiencyProficiencygeneralProfessional,
		"limitedworking":       SkillProficiencyProficiencylimitedWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkillProficiencyProficiency(input)
	return &out, nil
}
