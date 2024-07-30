package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionDetectionTimingType string

const (
	RiskDetectionDetectionTimingType_NearRealtime RiskDetectionDetectionTimingType = "nearRealtime"
	RiskDetectionDetectionTimingType_NotDefined   RiskDetectionDetectionTimingType = "notDefined"
	RiskDetectionDetectionTimingType_Offline      RiskDetectionDetectionTimingType = "offline"
	RiskDetectionDetectionTimingType_Realtime     RiskDetectionDetectionTimingType = "realtime"
)

func PossibleValuesForRiskDetectionDetectionTimingType() []string {
	return []string{
		string(RiskDetectionDetectionTimingType_NearRealtime),
		string(RiskDetectionDetectionTimingType_NotDefined),
		string(RiskDetectionDetectionTimingType_Offline),
		string(RiskDetectionDetectionTimingType_Realtime),
	}
}

func (s *RiskDetectionDetectionTimingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetectionDetectionTimingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetectionDetectionTimingType(input string) (*RiskDetectionDetectionTimingType, error) {
	vals := map[string]RiskDetectionDetectionTimingType{
		"nearrealtime": RiskDetectionDetectionTimingType_NearRealtime,
		"notdefined":   RiskDetectionDetectionTimingType_NotDefined,
		"offline":      RiskDetectionDetectionTimingType_Offline,
		"realtime":     RiskDetectionDetectionTimingType_Realtime,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionDetectionTimingType(input)
	return &out, nil
}
