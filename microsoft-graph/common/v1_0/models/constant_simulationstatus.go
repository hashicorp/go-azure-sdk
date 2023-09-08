package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SimulationStatus string

const (
	SimulationStatuscancelled SimulationStatus = "Cancelled"
	SimulationStatusdraft     SimulationStatus = "Draft"
	SimulationStatusexcluded  SimulationStatus = "Excluded"
	SimulationStatusfailed    SimulationStatus = "Failed"
	SimulationStatusrunning   SimulationStatus = "Running"
	SimulationStatusscheduled SimulationStatus = "Scheduled"
	SimulationStatussucceeded SimulationStatus = "Succeeded"
	SimulationStatusunknown   SimulationStatus = "Unknown"
)

func PossibleValuesForSimulationStatus() []string {
	return []string{
		string(SimulationStatuscancelled),
		string(SimulationStatusdraft),
		string(SimulationStatusexcluded),
		string(SimulationStatusfailed),
		string(SimulationStatusrunning),
		string(SimulationStatusscheduled),
		string(SimulationStatussucceeded),
		string(SimulationStatusunknown),
	}
}

func parseSimulationStatus(input string) (*SimulationStatus, error) {
	vals := map[string]SimulationStatus{
		"cancelled": SimulationStatuscancelled,
		"draft":     SimulationStatusdraft,
		"excluded":  SimulationStatusexcluded,
		"failed":    SimulationStatusfailed,
		"running":   SimulationStatusrunning,
		"scheduled": SimulationStatusscheduled,
		"succeeded": SimulationStatussucceeded,
		"unknown":   SimulationStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SimulationStatus(input)
	return &out, nil
}
