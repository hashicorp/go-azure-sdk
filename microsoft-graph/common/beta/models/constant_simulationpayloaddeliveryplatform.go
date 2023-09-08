package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationPayloadDeliveryPlatform string

const (
	SimulationPayloadDeliveryPlatformemail   SimulationPayloadDeliveryPlatform = "Email"
	SimulationPayloadDeliveryPlatformsms     SimulationPayloadDeliveryPlatform = "Sms"
	SimulationPayloadDeliveryPlatformteams   SimulationPayloadDeliveryPlatform = "Teams"
	SimulationPayloadDeliveryPlatformunknown SimulationPayloadDeliveryPlatform = "Unknown"
)

func PossibleValuesForSimulationPayloadDeliveryPlatform() []string {
	return []string{
		string(SimulationPayloadDeliveryPlatformemail),
		string(SimulationPayloadDeliveryPlatformsms),
		string(SimulationPayloadDeliveryPlatformteams),
		string(SimulationPayloadDeliveryPlatformunknown),
	}
}

func parseSimulationPayloadDeliveryPlatform(input string) (*SimulationPayloadDeliveryPlatform, error) {
	vals := map[string]SimulationPayloadDeliveryPlatform{
		"email":   SimulationPayloadDeliveryPlatformemail,
		"sms":     SimulationPayloadDeliveryPlatformsms,
		"teams":   SimulationPayloadDeliveryPlatformteams,
		"unknown": SimulationPayloadDeliveryPlatformunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationPayloadDeliveryPlatform(input)
	return &out, nil
}
