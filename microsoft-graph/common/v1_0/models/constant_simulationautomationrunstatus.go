package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAutomationRunStatus string

const (
	SimulationAutomationRunStatusfailed    SimulationAutomationRunStatus = "Failed"
	SimulationAutomationRunStatusrunning   SimulationAutomationRunStatus = "Running"
	SimulationAutomationRunStatusskipped   SimulationAutomationRunStatus = "Skipped"
	SimulationAutomationRunStatussucceeded SimulationAutomationRunStatus = "Succeeded"
	SimulationAutomationRunStatusunknown   SimulationAutomationRunStatus = "Unknown"
)

func PossibleValuesForSimulationAutomationRunStatus() []string {
	return []string{
		string(SimulationAutomationRunStatusfailed),
		string(SimulationAutomationRunStatusrunning),
		string(SimulationAutomationRunStatusskipped),
		string(SimulationAutomationRunStatussucceeded),
		string(SimulationAutomationRunStatusunknown),
	}
}

func parseSimulationAutomationRunStatus(input string) (*SimulationAutomationRunStatus, error) {
	vals := map[string]SimulationAutomationRunStatus{
		"failed":    SimulationAutomationRunStatusfailed,
		"running":   SimulationAutomationRunStatusrunning,
		"skipped":   SimulationAutomationRunStatusskipped,
		"succeeded": SimulationAutomationRunStatussucceeded,
		"unknown":   SimulationAutomationRunStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationAutomationRunStatus(input)
	return &out, nil
}
