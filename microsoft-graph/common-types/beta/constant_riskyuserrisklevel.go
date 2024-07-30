package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserRiskLevel string

const (
	RiskyUserRiskLevel_Hidden RiskyUserRiskLevel = "hidden"
	RiskyUserRiskLevel_High   RiskyUserRiskLevel = "high"
	RiskyUserRiskLevel_Low    RiskyUserRiskLevel = "low"
	RiskyUserRiskLevel_Medium RiskyUserRiskLevel = "medium"
	RiskyUserRiskLevel_None   RiskyUserRiskLevel = "none"
)

func PossibleValuesForRiskyUserRiskLevel() []string {
	return []string{
		string(RiskyUserRiskLevel_Hidden),
		string(RiskyUserRiskLevel_High),
		string(RiskyUserRiskLevel_Low),
		string(RiskyUserRiskLevel_Medium),
		string(RiskyUserRiskLevel_None),
	}
}

func (s *RiskyUserRiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyUserRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyUserRiskLevel(input string) (*RiskyUserRiskLevel, error) {
	vals := map[string]RiskyUserRiskLevel{
		"hidden": RiskyUserRiskLevel_Hidden,
		"high":   RiskyUserRiskLevel_High,
		"low":    RiskyUserRiskLevel_Low,
		"medium": RiskyUserRiskLevel_Medium,
		"none":   RiskyUserRiskLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserRiskLevel(input)
	return &out, nil
}
