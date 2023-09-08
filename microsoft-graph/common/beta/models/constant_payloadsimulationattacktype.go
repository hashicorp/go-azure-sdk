package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadSimulationAttackType string

const (
	PayloadSimulationAttackTypecloud    PayloadSimulationAttackType = "Cloud"
	PayloadSimulationAttackTypeendpoint PayloadSimulationAttackType = "Endpoint"
	PayloadSimulationAttackTypesocial   PayloadSimulationAttackType = "Social"
	PayloadSimulationAttackTypeunknown  PayloadSimulationAttackType = "Unknown"
)

func PossibleValuesForPayloadSimulationAttackType() []string {
	return []string{
		string(PayloadSimulationAttackTypecloud),
		string(PayloadSimulationAttackTypeendpoint),
		string(PayloadSimulationAttackTypesocial),
		string(PayloadSimulationAttackTypeunknown),
	}
}

func parsePayloadSimulationAttackType(input string) (*PayloadSimulationAttackType, error) {
	vals := map[string]PayloadSimulationAttackType{
		"cloud":    PayloadSimulationAttackTypecloud,
		"endpoint": PayloadSimulationAttackTypeendpoint,
		"social":   PayloadSimulationAttackTypesocial,
		"unknown":  PayloadSimulationAttackTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadSimulationAttackType(input)
	return &out, nil
}
