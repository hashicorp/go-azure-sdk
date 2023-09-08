package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyReading string

const (
	LanguageProficiencyReadingconversational      LanguageProficiencyReading = "Conversational"
	LanguageProficiencyReadingelementary          LanguageProficiencyReading = "Elementary"
	LanguageProficiencyReadingfullProfessional    LanguageProficiencyReading = "FullProfessional"
	LanguageProficiencyReadinglimitedWorking      LanguageProficiencyReading = "LimitedWorking"
	LanguageProficiencyReadingnativeOrBilingual   LanguageProficiencyReading = "NativeOrBilingual"
	LanguageProficiencyReadingprofessionalWorking LanguageProficiencyReading = "ProfessionalWorking"
)

func PossibleValuesForLanguageProficiencyReading() []string {
	return []string{
		string(LanguageProficiencyReadingconversational),
		string(LanguageProficiencyReadingelementary),
		string(LanguageProficiencyReadingfullProfessional),
		string(LanguageProficiencyReadinglimitedWorking),
		string(LanguageProficiencyReadingnativeOrBilingual),
		string(LanguageProficiencyReadingprofessionalWorking),
	}
}

func parseLanguageProficiencyReading(input string) (*LanguageProficiencyReading, error) {
	vals := map[string]LanguageProficiencyReading{
		"conversational":      LanguageProficiencyReadingconversational,
		"elementary":          LanguageProficiencyReadingelementary,
		"fullprofessional":    LanguageProficiencyReadingfullProfessional,
		"limitedworking":      LanguageProficiencyReadinglimitedWorking,
		"nativeorbilingual":   LanguageProficiencyReadingnativeOrBilingual,
		"professionalworking": LanguageProficiencyReadingprofessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyReading(input)
	return &out, nil
}
