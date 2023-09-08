package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAttackType string

const (
	SimulationAttackTypecloud    SimulationAttackType = "Cloud"
	SimulationAttackTypeendpoint SimulationAttackType = "Endpoint"
	SimulationAttackTypesocial   SimulationAttackType = "Social"
	SimulationAttackTypeunknown  SimulationAttackType = "Unknown"
)

func PossibleValuesForSimulationAttackType() []string {
	return []string{
		string(SimulationAttackTypecloud),
		string(SimulationAttackTypeendpoint),
		string(SimulationAttackTypesocial),
		string(SimulationAttackTypeunknown),
	}
}

func parseSimulationAttackType(input string) (*SimulationAttackType, error) {
	vals := map[string]SimulationAttackType{
		"cloud":    SimulationAttackTypecloud,
		"endpoint": SimulationAttackTypeendpoint,
		"social":   SimulationAttackTypesocial,
		"unknown":  SimulationAttackTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationAttackType(input)
	return &out, nil
}
