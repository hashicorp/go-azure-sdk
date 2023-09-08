package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationAutomationStatus string

const (
	SimulationAutomationStatuscompleted  SimulationAutomationStatus = "Completed"
	SimulationAutomationStatusdraft      SimulationAutomationStatus = "Draft"
	SimulationAutomationStatusnotRunning SimulationAutomationStatus = "NotRunning"
	SimulationAutomationStatusrunning    SimulationAutomationStatus = "Running"
	SimulationAutomationStatusunknown    SimulationAutomationStatus = "Unknown"
)

func PossibleValuesForSimulationAutomationStatus() []string {
	return []string{
		string(SimulationAutomationStatuscompleted),
		string(SimulationAutomationStatusdraft),
		string(SimulationAutomationStatusnotRunning),
		string(SimulationAutomationStatusrunning),
		string(SimulationAutomationStatusunknown),
	}
}

func parseSimulationAutomationStatus(input string) (*SimulationAutomationStatus, error) {
	vals := map[string]SimulationAutomationStatus{
		"completed":  SimulationAutomationStatuscompleted,
		"draft":      SimulationAutomationStatusdraft,
		"notrunning": SimulationAutomationStatusnotRunning,
		"running":    SimulationAutomationStatusrunning,
		"unknown":    SimulationAutomationStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationAutomationStatus(input)
	return &out, nil
}
