package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityPolicySettingsApplicableTo string

const (
	SensitivityPolicySettingsApplicableTo_Email        SensitivityPolicySettingsApplicableTo = "email"
	SensitivityPolicySettingsApplicableTo_Site         SensitivityPolicySettingsApplicableTo = "site"
	SensitivityPolicySettingsApplicableTo_Teamwork     SensitivityPolicySettingsApplicableTo = "teamwork"
	SensitivityPolicySettingsApplicableTo_UnifiedGroup SensitivityPolicySettingsApplicableTo = "unifiedGroup"
)

func PossibleValuesForSensitivityPolicySettingsApplicableTo() []string {
	return []string{
		string(SensitivityPolicySettingsApplicableTo_Email),
		string(SensitivityPolicySettingsApplicableTo_Site),
		string(SensitivityPolicySettingsApplicableTo_Teamwork),
		string(SensitivityPolicySettingsApplicableTo_UnifiedGroup),
	}
}

func (s *SensitivityPolicySettingsApplicableTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityPolicySettingsApplicableTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityPolicySettingsApplicableTo(input string) (*SensitivityPolicySettingsApplicableTo, error) {
	vals := map[string]SensitivityPolicySettingsApplicableTo{
		"email":        SensitivityPolicySettingsApplicableTo_Email,
		"site":         SensitivityPolicySettingsApplicableTo_Site,
		"teamwork":     SensitivityPolicySettingsApplicableTo_Teamwork,
		"unifiedgroup": SensitivityPolicySettingsApplicableTo_UnifiedGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityPolicySettingsApplicableTo(input)
	return &out, nil
}
