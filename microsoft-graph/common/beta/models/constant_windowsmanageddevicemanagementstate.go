package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceManagementState string

const (
	WindowsManagedDeviceManagementStatedeletePending  WindowsManagedDeviceManagementState = "DeletePending"
	WindowsManagedDeviceManagementStatediscovered     WindowsManagedDeviceManagementState = "Discovered"
	WindowsManagedDeviceManagementStatemanaged        WindowsManagedDeviceManagementState = "Managed"
	WindowsManagedDeviceManagementStateretireCanceled WindowsManagedDeviceManagementState = "RetireCanceled"
	WindowsManagedDeviceManagementStateretireFailed   WindowsManagedDeviceManagementState = "RetireFailed"
	WindowsManagedDeviceManagementStateretireIssued   WindowsManagedDeviceManagementState = "RetireIssued"
	WindowsManagedDeviceManagementStateretirePending  WindowsManagedDeviceManagementState = "RetirePending"
	WindowsManagedDeviceManagementStateunhealthy      WindowsManagedDeviceManagementState = "Unhealthy"
	WindowsManagedDeviceManagementStatewipeCanceled   WindowsManagedDeviceManagementState = "WipeCanceled"
	WindowsManagedDeviceManagementStatewipeFailed     WindowsManagedDeviceManagementState = "WipeFailed"
	WindowsManagedDeviceManagementStatewipeIssued     WindowsManagedDeviceManagementState = "WipeIssued"
	WindowsManagedDeviceManagementStatewipePending    WindowsManagedDeviceManagementState = "WipePending"
)

func PossibleValuesForWindowsManagedDeviceManagementState() []string {
	return []string{
		string(WindowsManagedDeviceManagementStatedeletePending),
		string(WindowsManagedDeviceManagementStatediscovered),
		string(WindowsManagedDeviceManagementStatemanaged),
		string(WindowsManagedDeviceManagementStateretireCanceled),
		string(WindowsManagedDeviceManagementStateretireFailed),
		string(WindowsManagedDeviceManagementStateretireIssued),
		string(WindowsManagedDeviceManagementStateretirePending),
		string(WindowsManagedDeviceManagementStateunhealthy),
		string(WindowsManagedDeviceManagementStatewipeCanceled),
		string(WindowsManagedDeviceManagementStatewipeFailed),
		string(WindowsManagedDeviceManagementStatewipeIssued),
		string(WindowsManagedDeviceManagementStatewipePending),
	}
}

func parseWindowsManagedDeviceManagementState(input string) (*WindowsManagedDeviceManagementState, error) {
	vals := map[string]WindowsManagedDeviceManagementState{
		"deletepending":  WindowsManagedDeviceManagementStatedeletePending,
		"discovered":     WindowsManagedDeviceManagementStatediscovered,
		"managed":        WindowsManagedDeviceManagementStatemanaged,
		"retirecanceled": WindowsManagedDeviceManagementStateretireCanceled,
		"retirefailed":   WindowsManagedDeviceManagementStateretireFailed,
		"retireissued":   WindowsManagedDeviceManagementStateretireIssued,
		"retirepending":  WindowsManagedDeviceManagementStateretirePending,
		"unhealthy":      WindowsManagedDeviceManagementStateunhealthy,
		"wipecanceled":   WindowsManagedDeviceManagementStatewipeCanceled,
		"wipefailed":     WindowsManagedDeviceManagementStatewipeFailed,
		"wipeissued":     WindowsManagedDeviceManagementStatewipeIssued,
		"wipepending":    WindowsManagedDeviceManagementStatewipePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceManagementState(input)
	return &out, nil
}
