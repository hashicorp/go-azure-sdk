package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyProficiency string

const (
	LanguageProficiencyProficiency_Conversational      LanguageProficiencyProficiency = "conversational"
	LanguageProficiencyProficiency_Elementary          LanguageProficiencyProficiency = "elementary"
	LanguageProficiencyProficiency_FullProfessional    LanguageProficiencyProficiency = "fullProfessional"
	LanguageProficiencyProficiency_LimitedWorking      LanguageProficiencyProficiency = "limitedWorking"
	LanguageProficiencyProficiency_NativeOrBilingual   LanguageProficiencyProficiency = "nativeOrBilingual"
	LanguageProficiencyProficiency_ProfessionalWorking LanguageProficiencyProficiency = "professionalWorking"
)

func PossibleValuesForLanguageProficiencyProficiency() []string {
	return []string{
		string(LanguageProficiencyProficiency_Conversational),
		string(LanguageProficiencyProficiency_Elementary),
		string(LanguageProficiencyProficiency_FullProfessional),
		string(LanguageProficiencyProficiency_LimitedWorking),
		string(LanguageProficiencyProficiency_NativeOrBilingual),
		string(LanguageProficiencyProficiency_ProfessionalWorking),
	}
}

func (s *LanguageProficiencyProficiency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyProficiency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyProficiency(input string) (*LanguageProficiencyProficiency, error) {
	vals := map[string]LanguageProficiencyProficiency{
		"conversational":      LanguageProficiencyProficiency_Conversational,
		"elementary":          LanguageProficiencyProficiency_Elementary,
		"fullprofessional":    LanguageProficiencyProficiency_FullProfessional,
		"limitedworking":      LanguageProficiencyProficiency_LimitedWorking,
		"nativeorbilingual":   LanguageProficiencyProficiency_NativeOrBilingual,
		"professionalworking": LanguageProficiencyProficiency_ProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyProficiency(input)
	return &out, nil
}
