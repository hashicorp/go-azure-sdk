package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyReading string

const (
	LanguageProficiencyReading_Conversational      LanguageProficiencyReading = "conversational"
	LanguageProficiencyReading_Elementary          LanguageProficiencyReading = "elementary"
	LanguageProficiencyReading_FullProfessional    LanguageProficiencyReading = "fullProfessional"
	LanguageProficiencyReading_LimitedWorking      LanguageProficiencyReading = "limitedWorking"
	LanguageProficiencyReading_NativeOrBilingual   LanguageProficiencyReading = "nativeOrBilingual"
	LanguageProficiencyReading_ProfessionalWorking LanguageProficiencyReading = "professionalWorking"
)

func PossibleValuesForLanguageProficiencyReading() []string {
	return []string{
		string(LanguageProficiencyReading_Conversational),
		string(LanguageProficiencyReading_Elementary),
		string(LanguageProficiencyReading_FullProfessional),
		string(LanguageProficiencyReading_LimitedWorking),
		string(LanguageProficiencyReading_NativeOrBilingual),
		string(LanguageProficiencyReading_ProfessionalWorking),
	}
}

func (s *LanguageProficiencyReading) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyReading(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyReading(input string) (*LanguageProficiencyReading, error) {
	vals := map[string]LanguageProficiencyReading{
		"conversational":      LanguageProficiencyReading_Conversational,
		"elementary":          LanguageProficiencyReading_Elementary,
		"fullprofessional":    LanguageProficiencyReading_FullProfessional,
		"limitedworking":      LanguageProficiencyReading_LimitedWorking,
		"nativeorbilingual":   LanguageProficiencyReading_NativeOrBilingual,
		"professionalworking": LanguageProficiencyReading_ProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyReading(input)
	return &out, nil
}
