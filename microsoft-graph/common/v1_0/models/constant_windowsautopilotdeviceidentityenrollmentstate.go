package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsAutopilotDeviceIdentityEnrollmentState string

const (
	WindowsAutopilotDeviceIdentityEnrollmentStateenrolled     WindowsAutopilotDeviceIdentityEnrollmentState = "Enrolled"
	WindowsAutopilotDeviceIdentityEnrollmentStatefailed       WindowsAutopilotDeviceIdentityEnrollmentState = "Failed"
	WindowsAutopilotDeviceIdentityEnrollmentStatenotContacted WindowsAutopilotDeviceIdentityEnrollmentState = "NotContacted"
	WindowsAutopilotDeviceIdentityEnrollmentStatependingReset WindowsAutopilotDeviceIdentityEnrollmentState = "PendingReset"
	WindowsAutopilotDeviceIdentityEnrollmentStateunknown      WindowsAutopilotDeviceIdentityEnrollmentState = "Unknown"
)

func PossibleValuesForWindowsAutopilotDeviceIdentityEnrollmentState() []string {
	return []string{
		string(WindowsAutopilotDeviceIdentityEnrollmentStateenrolled),
		string(WindowsAutopilotDeviceIdentityEnrollmentStatefailed),
		string(WindowsAutopilotDeviceIdentityEnrollmentStatenotContacted),
		string(WindowsAutopilotDeviceIdentityEnrollmentStatependingReset),
		string(WindowsAutopilotDeviceIdentityEnrollmentStateunknown),
	}
}

func parseWindowsAutopilotDeviceIdentityEnrollmentState(input string) (*WindowsAutopilotDeviceIdentityEnrollmentState, error) {
	vals := map[string]WindowsAutopilotDeviceIdentityEnrollmentState{
		"enrolled":     WindowsAutopilotDeviceIdentityEnrollmentStateenrolled,
		"failed":       WindowsAutopilotDeviceIdentityEnrollmentStatefailed,
		"notcontacted": WindowsAutopilotDeviceIdentityEnrollmentStatenotContacted,
		"pendingreset": WindowsAutopilotDeviceIdentityEnrollmentStatependingReset,
		"unknown":      WindowsAutopilotDeviceIdentityEnrollmentStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsAutopilotDeviceIdentityEnrollmentState(input)
	return &out, nil
}
