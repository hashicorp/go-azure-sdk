package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationOperationStatus string

const (
	AttackSimulationOperationStatusfailed     AttackSimulationOperationStatus = "Failed"
	AttackSimulationOperationStatusnotStarted AttackSimulationOperationStatus = "NotStarted"
	AttackSimulationOperationStatusrunning    AttackSimulationOperationStatus = "Running"
	AttackSimulationOperationStatussucceeded  AttackSimulationOperationStatus = "Succeeded"
)

func PossibleValuesForAttackSimulationOperationStatus() []string {
	return []string{
		string(AttackSimulationOperationStatusfailed),
		string(AttackSimulationOperationStatusnotStarted),
		string(AttackSimulationOperationStatusrunning),
		string(AttackSimulationOperationStatussucceeded),
	}
}

func parseAttackSimulationOperationStatus(input string) (*AttackSimulationOperationStatus, error) {
	vals := map[string]AttackSimulationOperationStatus{
		"failed":     AttackSimulationOperationStatusfailed,
		"notstarted": AttackSimulationOperationStatusnotStarted,
		"running":    AttackSimulationOperationStatusrunning,
		"succeeded":  AttackSimulationOperationStatussucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AttackSimulationOperationStatus(input)
	return &out, nil
}
