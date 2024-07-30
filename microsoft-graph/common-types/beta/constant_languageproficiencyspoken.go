package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencySpoken string

const (
	LanguageProficiencySpoken_Conversational      LanguageProficiencySpoken = "conversational"
	LanguageProficiencySpoken_Elementary          LanguageProficiencySpoken = "elementary"
	LanguageProficiencySpoken_FullProfessional    LanguageProficiencySpoken = "fullProfessional"
	LanguageProficiencySpoken_LimitedWorking      LanguageProficiencySpoken = "limitedWorking"
	LanguageProficiencySpoken_NativeOrBilingual   LanguageProficiencySpoken = "nativeOrBilingual"
	LanguageProficiencySpoken_ProfessionalWorking LanguageProficiencySpoken = "professionalWorking"
)

func PossibleValuesForLanguageProficiencySpoken() []string {
	return []string{
		string(LanguageProficiencySpoken_Conversational),
		string(LanguageProficiencySpoken_Elementary),
		string(LanguageProficiencySpoken_FullProfessional),
		string(LanguageProficiencySpoken_LimitedWorking),
		string(LanguageProficiencySpoken_NativeOrBilingual),
		string(LanguageProficiencySpoken_ProfessionalWorking),
	}
}

func (s *LanguageProficiencySpoken) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencySpoken(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencySpoken(input string) (*LanguageProficiencySpoken, error) {
	vals := map[string]LanguageProficiencySpoken{
		"conversational":      LanguageProficiencySpoken_Conversational,
		"elementary":          LanguageProficiencySpoken_Elementary,
		"fullprofessional":    LanguageProficiencySpoken_FullProfessional,
		"limitedworking":      LanguageProficiencySpoken_LimitedWorking,
		"nativeorbilingual":   LanguageProficiencySpoken_NativeOrBilingual,
		"professionalworking": LanguageProficiencySpoken_ProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencySpoken(input)
	return &out, nil
}
