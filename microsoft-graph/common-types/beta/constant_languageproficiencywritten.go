package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyWritten string

const (
	LanguageProficiencyWritten_Conversational      LanguageProficiencyWritten = "conversational"
	LanguageProficiencyWritten_Elementary          LanguageProficiencyWritten = "elementary"
	LanguageProficiencyWritten_FullProfessional    LanguageProficiencyWritten = "fullProfessional"
	LanguageProficiencyWritten_LimitedWorking      LanguageProficiencyWritten = "limitedWorking"
	LanguageProficiencyWritten_NativeOrBilingual   LanguageProficiencyWritten = "nativeOrBilingual"
	LanguageProficiencyWritten_ProfessionalWorking LanguageProficiencyWritten = "professionalWorking"
)

func PossibleValuesForLanguageProficiencyWritten() []string {
	return []string{
		string(LanguageProficiencyWritten_Conversational),
		string(LanguageProficiencyWritten_Elementary),
		string(LanguageProficiencyWritten_FullProfessional),
		string(LanguageProficiencyWritten_LimitedWorking),
		string(LanguageProficiencyWritten_NativeOrBilingual),
		string(LanguageProficiencyWritten_ProfessionalWorking),
	}
}

func (s *LanguageProficiencyWritten) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLanguageProficiencyWritten(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLanguageProficiencyWritten(input string) (*LanguageProficiencyWritten, error) {
	vals := map[string]LanguageProficiencyWritten{
		"conversational":      LanguageProficiencyWritten_Conversational,
		"elementary":          LanguageProficiencyWritten_Elementary,
		"fullprofessional":    LanguageProficiencyWritten_FullProfessional,
		"limitedworking":      LanguageProficiencyWritten_LimitedWorking,
		"nativeorbilingual":   LanguageProficiencyWritten_NativeOrBilingual,
		"professionalworking": LanguageProficiencyWritten_ProfessionalWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyWritten(input)
	return &out, nil
}
