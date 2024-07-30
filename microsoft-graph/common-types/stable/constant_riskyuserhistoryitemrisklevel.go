package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskyUserHistoryItemRiskLevel string

const (
	RiskyUserHistoryItemRiskLevel_Hidden RiskyUserHistoryItemRiskLevel = "hidden"
	RiskyUserHistoryItemRiskLevel_High   RiskyUserHistoryItemRiskLevel = "high"
	RiskyUserHistoryItemRiskLevel_Low    RiskyUserHistoryItemRiskLevel = "low"
	RiskyUserHistoryItemRiskLevel_Medium RiskyUserHistoryItemRiskLevel = "medium"
	RiskyUserHistoryItemRiskLevel_None   RiskyUserHistoryItemRiskLevel = "none"
)

func PossibleValuesForRiskyUserHistoryItemRiskLevel() []string {
	return []string{
		string(RiskyUserHistoryItemRiskLevel_Hidden),
		string(RiskyUserHistoryItemRiskLevel_High),
		string(RiskyUserHistoryItemRiskLevel_Low),
		string(RiskyUserHistoryItemRiskLevel_Medium),
		string(RiskyUserHistoryItemRiskLevel_None),
	}
}

func (s *RiskyUserHistoryItemRiskLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskyUserHistoryItemRiskLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskyUserHistoryItemRiskLevel(input string) (*RiskyUserHistoryItemRiskLevel, error) {
	vals := map[string]RiskyUserHistoryItemRiskLevel{
		"hidden": RiskyUserHistoryItemRiskLevel_Hidden,
		"high":   RiskyUserHistoryItemRiskLevel_High,
		"low":    RiskyUserHistoryItemRiskLevel_Low,
		"medium": RiskyUserHistoryItemRiskLevel_Medium,
		"none":   RiskyUserHistoryItemRiskLevel_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskyUserHistoryItemRiskLevel(input)
	return &out, nil
}
