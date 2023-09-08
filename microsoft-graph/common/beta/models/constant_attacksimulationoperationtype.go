package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationOperationType string

const (
	AttackSimulationOperationTypecreateSimualation AttackSimulationOperationType = "CreateSimualation"
	AttackSimulationOperationTypeupdateSimulation  AttackSimulationOperationType = "UpdateSimulation"
)

func PossibleValuesForAttackSimulationOperationType() []string {
	return []string{
		string(AttackSimulationOperationTypecreateSimualation),
		string(AttackSimulationOperationTypeupdateSimulation),
	}
}

func parseAttackSimulationOperationType(input string) (*AttackSimulationOperationType, error) {
	vals := map[string]AttackSimulationOperationType{
		"createsimualation": AttackSimulationOperationTypecreateSimualation,
		"updatesimulation":  AttackSimulationOperationTypeupdateSimulation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttackSimulationOperationType(input)
	return &out, nil
}
