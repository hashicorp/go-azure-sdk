package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityRemediationState string

const (
	WindowsAutopilotDeviceIdentityRemediationStateautomaticRemediationRequired WindowsAutopilotDeviceIdentityRemediationState = "AutomaticRemediationRequired"
	WindowsAutopilotDeviceIdentityRemediationStatemanualRemediationRequired    WindowsAutopilotDeviceIdentityRemediationState = "ManualRemediationRequired"
	WindowsAutopilotDeviceIdentityRemediationStatenoRemediationRequired        WindowsAutopilotDeviceIdentityRemediationState = "NoRemediationRequired"
	WindowsAutopilotDeviceIdentityRemediationStateunknown                      WindowsAutopilotDeviceIdentityRemediationState = "Unknown"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityRemediationState() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityRemediationStateautomaticRemediationRequired),
		string(WindowsAutopilotDeviceIdentityRemediationStatemanualRemediationRequired),
		string(WindowsAutopilotDeviceIdentityRemediationStatenoRemediationRequired),
		string(WindowsAutopilotDeviceIdentityRemediationStateunknown),
	}
}

func parseWindowsAutopilotDeviceIdentityRemediationState(input string) (*WindowsAutopilotDeviceIdentityRemediationState, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityRemediationState{
		"automaticremediationrequired": WindowsAutopilotDeviceIdentityRemediationStateautomaticRemediationRequired,
		"manualremediationrequired":    WindowsAutopilotDeviceIdentityRemediationStatemanualRemediationRequired,
		"noremediationrequired":        WindowsAutopilotDeviceIdentityRemediationStatenoRemediationRequired,
		"unknown":                      WindowsAutopilotDeviceIdentityRemediationStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityRemediationState(input)
	return &out, nil
}
