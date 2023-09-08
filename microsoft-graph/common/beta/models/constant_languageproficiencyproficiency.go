package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyProficiency string

const (
	LanguageProficiencyProficiencyconversational      LanguageProficiencyProficiency = "Conversational"
	LanguageProficiencyProficiencyelementary          LanguageProficiencyProficiency = "Elementary"
	LanguageProficiencyProficiencyfullProfessional    LanguageProficiencyProficiency = "FullProfessional"
	LanguageProficiencyProficiencylimitedWorking      LanguageProficiencyProficiency = "LimitedWorking"
	LanguageProficiencyProficiencynativeOrBilingual   LanguageProficiencyProficiency = "NativeOrBilingual"
	LanguageProficiencyProficiencyprofessionalWorking LanguageProficiencyProficiency = "ProfessionalWorking"
)

func PossibleValuesForLanguageProficiencyProficiency() []string {
	return []string{
		string(LanguageProficiencyProficiencyconversational),
		string(LanguageProficiencyProficiencyelementary),
		string(LanguageProficiencyProficiencyfullProfessional),
		string(LanguageProficiencyProficiencylimitedWorking),
		string(LanguageProficiencyProficiencynativeOrBilingual),
		string(LanguageProficiencyProficiencyprofessionalWorking),
	}
}

func parseLanguageProficiencyProficiency(input string) (*LanguageProficiencyProficiency, error) {
	vals := map[string]LanguageProficiencyProficiency{
		"conversational":      LanguageProficiencyProficiencyconversational,
		"elementary":          LanguageProficiencyProficiencyelementary,
		"fullprofessional":    LanguageProficiencyProficiencyfullProfessional,
		"limitedworking":      LanguageProficiencyProficiencylimitedWorking,
		"nativeorbilingual":   LanguageProficiencyProficiencynativeOrBilingual,
		"professionalworking": LanguageProficiencyProficiencyprofessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyProficiency(input)
	return &out, nil
}
