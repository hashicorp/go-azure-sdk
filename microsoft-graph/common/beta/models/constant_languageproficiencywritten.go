package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyWritten string

const (
	LanguageProficiencyWrittenconversational      LanguageProficiencyWritten = "Conversational"
	LanguageProficiencyWrittenelementary          LanguageProficiencyWritten = "Elementary"
	LanguageProficiencyWrittenfullProfessional    LanguageProficiencyWritten = "FullProfessional"
	LanguageProficiencyWrittenlimitedWorking      LanguageProficiencyWritten = "LimitedWorking"
	LanguageProficiencyWrittennativeOrBilingual   LanguageProficiencyWritten = "NativeOrBilingual"
	LanguageProficiencyWrittenprofessionalWorking LanguageProficiencyWritten = "ProfessionalWorking"
)

func PossibleValuesForLanguageProficiencyWritten() []string {
	return []string{
		string(LanguageProficiencyWrittenconversational),
		string(LanguageProficiencyWrittenelementary),
		string(LanguageProficiencyWrittenfullProfessional),
		string(LanguageProficiencyWrittenlimitedWorking),
		string(LanguageProficiencyWrittennativeOrBilingual),
		string(LanguageProficiencyWrittenprofessionalWorking),
	}
}

func parseLanguageProficiencyWritten(input string) (*LanguageProficiencyWritten, error) {
	vals := map[string]LanguageProficiencyWritten{
		"conversational":      LanguageProficiencyWrittenconversational,
		"elementary":          LanguageProficiencyWrittenelementary,
		"fullprofessional":    LanguageProficiencyWrittenfullProfessional,
		"limitedworking":      LanguageProficiencyWrittenlimitedWorking,
		"nativeorbilingual":   LanguageProficiencyWrittennativeOrBilingual,
		"professionalworking": LanguageProficiencyWrittenprofessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyWritten(input)
	return &out, nil
}
