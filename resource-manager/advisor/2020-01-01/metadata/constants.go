package metadata

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
