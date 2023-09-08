package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesMonitoringRuleSignal string

const (
	WindowsUpdatesMonitoringRuleSignalineligible WindowsUpdatesMonitoringRuleSignal = "Ineligible"
	WindowsUpdatesMonitoringRuleSignalrollback   WindowsUpdatesMonitoringRuleSignal = "Rollback"
)

func PossibleValuesForWindowsUpdatesMonitoringRuleSignal() []string {
	return []string{
		string(WindowsUpdatesMonitoringRuleSignalineligible),
		string(WindowsUpdatesMonitoringRuleSignalrollback),
	}
}

func parseWindowsUpdatesMonitoringRuleSignal(input string) (*WindowsUpdatesMonitoringRuleSignal, error) {
	vals := map[string]WindowsUpdatesMonitoringRuleSignal{
		"ineligible": WindowsUpdatesMonitoringRuleSignalineligible,
		"rollback":   WindowsUpdatesMonitoringRuleSignalrollback,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdatesMonitoringRuleSignal(input)
	return &out, nil
}
