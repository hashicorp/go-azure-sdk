package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelApplicableTo string

const (
	SensitivityLabelApplicableTo_Email        SensitivityLabelApplicableTo = "email"
	SensitivityLabelApplicableTo_Site         SensitivityLabelApplicableTo = "site"
	SensitivityLabelApplicableTo_Teamwork     SensitivityLabelApplicableTo = "teamwork"
	SensitivityLabelApplicableTo_UnifiedGroup SensitivityLabelApplicableTo = "unifiedGroup"
)

func PossibleValuesForSensitivityLabelApplicableTo() []string {
	return []string{
		string(SensitivityLabelApplicableTo_Email),
		string(SensitivityLabelApplicableTo_Site),
		string(SensitivityLabelApplicableTo_Teamwork),
		string(SensitivityLabelApplicableTo_UnifiedGroup),
	}
}

func (s *SensitivityLabelApplicableTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSensitivityLabelApplicableTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSensitivityLabelApplicableTo(input string) (*SensitivityLabelApplicableTo, error) {
	vals := map[string]SensitivityLabelApplicableTo{
		"email":        SensitivityLabelApplicableTo_Email,
		"site":         SensitivityLabelApplicableTo_Site,
		"teamwork":     SensitivityLabelApplicableTo_Teamwork,
		"unifiedgroup": SensitivityLabelApplicableTo_UnifiedGroup,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelApplicableTo(input)
	return &out, nil
}
