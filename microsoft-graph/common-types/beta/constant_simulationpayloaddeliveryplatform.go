package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationPayloadDeliveryPlatform string

const (
	SimulationPayloadDeliveryPlatform_Email   SimulationPayloadDeliveryPlatform = "email"
	SimulationPayloadDeliveryPlatform_Sms     SimulationPayloadDeliveryPlatform = "sms"
	SimulationPayloadDeliveryPlatform_Teams   SimulationPayloadDeliveryPlatform = "teams"
	SimulationPayloadDeliveryPlatform_Unknown SimulationPayloadDeliveryPlatform = "unknown"
)

func PossibleValuesForSimulationPayloadDeliveryPlatform() []string {
	return []string{
		string(SimulationPayloadDeliveryPlatform_Email),
		string(SimulationPayloadDeliveryPlatform_Sms),
		string(SimulationPayloadDeliveryPlatform_Teams),
		string(SimulationPayloadDeliveryPlatform_Unknown),
	}
}

func (s *SimulationPayloadDeliveryPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSimulationPayloadDeliveryPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSimulationPayloadDeliveryPlatform(input string) (*SimulationPayloadDeliveryPlatform, error) {
	vals := map[string]SimulationPayloadDeliveryPlatform{
		"email":   SimulationPayloadDeliveryPlatform_Email,
		"sms":     SimulationPayloadDeliveryPlatform_Sms,
		"teams":   SimulationPayloadDeliveryPlatform_Teams,
		"unknown": SimulationPayloadDeliveryPlatform_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationPayloadDeliveryPlatform(input)
	return &out, nil
}
