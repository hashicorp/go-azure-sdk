package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyServicePrincipalHistoryItemRiskLevel string

const (
	RiskyServicePrincipalHistoryItemRiskLevel_Hidden RiskyServicePrincipalHistoryItemRiskLevel = "hidden"
	RiskyServicePrincipalHistoryItemRiskLevel_High   RiskyServicePrincipalHistoryItemRiskLevel = "high"
	RiskyServicePrincipalHistoryItemRiskLevel_Low    RiskyServicePrincipalHistoryItemRiskLevel = "low"
	RiskyServicePrincipalHistoryItemRiskLevel_Medium RiskyServicePrincipalHistoryItemRiskLevel = "medium"
	RiskyServicePrincipalHistoryItemRiskLevel_None   RiskyServicePrincipalHistoryItemRiskLevel = "none"
)

func PossibleValuesForRiskyServicePrincipalHistoryItemRiskLevel() []string {
	return []string{
		string(RiskyServicePrincipalHistoryItemRiskLevel_Hidden),
		string(RiskyServicePrincipalHistoryItemRiskLevel_High),
		string(RiskyServicePrincipalHistoryItemRiskLevel_Low),
		string(RiskyServicePrincipalHistoryItemRiskLevel_Medium),
		string(RiskyServicePrincipalHistoryItemRiskLevel_None),
	}
}

func (s *RiskyServicePrincipalHistoryItemRiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyServicePrincipalHistoryItemRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyServicePrincipalHistoryItemRiskLevel(input string) (*RiskyServicePrincipalHistoryItemRiskLevel, error) {
	vals := map[string]RiskyServicePrincipalHistoryItemRiskLevel{
		"hidden": RiskyServicePrincipalHistoryItemRiskLevel_Hidden,
		"high":   RiskyServicePrincipalHistoryItemRiskLevel_High,
		"low":    RiskyServicePrincipalHistoryItemRiskLevel_Low,
		"medium": RiskyServicePrincipalHistoryItemRiskLevel_Medium,
		"none":   RiskyServicePrincipalHistoryItemRiskLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyServicePrincipalHistoryItemRiskLevel(input)
	return &out, nil
}
