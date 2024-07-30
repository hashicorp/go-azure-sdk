package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkillProficiencyProficiency string

const (
	SkillProficiencyProficiency_AdvancedProfessional SkillProficiencyProficiency = "advancedProfessional"
	SkillProficiencyProficiency_Elementary           SkillProficiencyProficiency = "elementary"
	SkillProficiencyProficiency_Expert               SkillProficiencyProficiency = "expert"
	SkillProficiencyProficiency_GeneralProfessional  SkillProficiencyProficiency = "generalProfessional"
	SkillProficiencyProficiency_LimitedWorking       SkillProficiencyProficiency = "limitedWorking"
)

func PossibleValuesForSkillProficiencyProficiency() []string {
	return []string{
		string(SkillProficiencyProficiency_AdvancedProfessional),
		string(SkillProficiencyProficiency_Elementary),
		string(SkillProficiencyProficiency_Expert),
		string(SkillProficiencyProficiency_GeneralProfessional),
		string(SkillProficiencyProficiency_LimitedWorking),
	}
}

func (s *SkillProficiencyProficiency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSkillProficiencyProficiency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSkillProficiencyProficiency(input string) (*SkillProficiencyProficiency, error) {
	vals := map[string]SkillProficiencyProficiency{
		"advancedprofessional": SkillProficiencyProficiency_AdvancedProfessional,
		"elementary":           SkillProficiencyProficiency_Elementary,
		"expert":               SkillProficiencyProficiency_Expert,
		"generalprofessional":  SkillProficiencyProficiency_GeneralProfessional,
		"limitedworking":       SkillProficiencyProficiency_LimitedWorking,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkillProficiencyProficiency(input)
	return &out, nil
}
