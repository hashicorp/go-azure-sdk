package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesMonitoringRuleAction string

const (
	WindowsUpdatesMonitoringRuleActionalertError      WindowsUpdatesMonitoringRuleAction = "AlertError"
	WindowsUpdatesMonitoringRuleActionofferFallback   WindowsUpdatesMonitoringRuleAction = "OfferFallback"
	WindowsUpdatesMonitoringRuleActionpauseDeployment WindowsUpdatesMonitoringRuleAction = "PauseDeployment"
)

func PossibleValuesForWindowsUpdatesMonitoringRuleAction() []string {
	return []string{
		string(WindowsUpdatesMonitoringRuleActionalertError),
		string(WindowsUpdatesMonitoringRuleActionofferFallback),
		string(WindowsUpdatesMonitoringRuleActionpauseDeployment),
	}
}

func parseWindowsUpdatesMonitoringRuleAction(input string) (*WindowsUpdatesMonitoringRuleAction, error) {
	vals := map[string]WindowsUpdatesMonitoringRuleAction{
		"alerterror":      WindowsUpdatesMonitoringRuleActionalertError,
		"offerfallback":   WindowsUpdatesMonitoringRuleActionofferFallback,
		"pausedeployment": WindowsUpdatesMonitoringRuleActionpauseDeployment,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesMonitoringRuleAction(input)
	return &out, nil
}
