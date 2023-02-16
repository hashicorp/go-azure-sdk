package metadata

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Scenario string

const (
	ScenarioAlerts Scenario = "Alerts"
)

func PossibleValuesForScenario() []string {
	return []string{
		string(ScenarioAlerts),
	}
}

func parseScenario(input string) (*Scenario, error) {
	vals := map[string]Scenario{
		"alerts": ScenarioAlerts,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Scenario(input)
	return &out, nil
}
