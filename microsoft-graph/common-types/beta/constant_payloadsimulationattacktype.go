package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadSimulationAttackType string

const (
	PayloadSimulationAttackType_Cloud    PayloadSimulationAttackType = "cloud"
	PayloadSimulationAttackType_Endpoint PayloadSimulationAttackType = "endpoint"
	PayloadSimulationAttackType_Social   PayloadSimulationAttackType = "social"
	PayloadSimulationAttackType_Unknown  PayloadSimulationAttackType = "unknown"
)

func PossibleValuesForPayloadSimulationAttackType() []string {
	return []string{
		string(PayloadSimulationAttackType_Cloud),
		string(PayloadSimulationAttackType_Endpoint),
		string(PayloadSimulationAttackType_Social),
		string(PayloadSimulationAttackType_Unknown),
	}
}

func (s *PayloadSimulationAttackType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadSimulationAttackType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadSimulationAttackType(input string) (*PayloadSimulationAttackType, error) {
	vals := map[string]PayloadSimulationAttackType{
		"cloud":    PayloadSimulationAttackType_Cloud,
		"endpoint": PayloadSimulationAttackType_Endpoint,
		"social":   PayloadSimulationAttackType_Social,
		"unknown":  PayloadSimulationAttackType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadSimulationAttackType(input)
	return &out, nil
}
