package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceComplianceState string

const (
	WindowsManagedDeviceComplianceStatecompliant     WindowsManagedDeviceComplianceState = "Compliant"
	WindowsManagedDeviceComplianceStateconfigManager WindowsManagedDeviceComplianceState = "ConfigManager"
	WindowsManagedDeviceComplianceStateconflict      WindowsManagedDeviceComplianceState = "Conflict"
	WindowsManagedDeviceComplianceStateerror         WindowsManagedDeviceComplianceState = "Error"
	WindowsManagedDeviceComplianceStateinGracePeriod WindowsManagedDeviceComplianceState = "InGracePeriod"
	WindowsManagedDeviceComplianceStatenoncompliant  WindowsManagedDeviceComplianceState = "Noncompliant"
	WindowsManagedDeviceComplianceStateunknown       WindowsManagedDeviceComplianceState = "Unknown"
)

func PossibleValuesForWindowsManagedDeviceComplianceState() []string {
	return []string{
		string(WindowsManagedDeviceComplianceStatecompliant),
		string(WindowsManagedDeviceComplianceStateconfigManager),
		string(WindowsManagedDeviceComplianceStateconflict),
		string(WindowsManagedDeviceComplianceStateerror),
		string(WindowsManagedDeviceComplianceStateinGracePeriod),
		string(WindowsManagedDeviceComplianceStatenoncompliant),
		string(WindowsManagedDeviceComplianceStateunknown),
	}
}

func parseWindowsManagedDeviceComplianceState(input string) (*WindowsManagedDeviceComplianceState, error) {
	vals := map[string]WindowsManagedDeviceComplianceState{
		"compliant":     WindowsManagedDeviceComplianceStatecompliant,
		"configmanager": WindowsManagedDeviceComplianceStateconfigManager,
		"conflict":      WindowsManagedDeviceComplianceStateconflict,
		"error":         WindowsManagedDeviceComplianceStateerror,
		"ingraceperiod": WindowsManagedDeviceComplianceStateinGracePeriod,
		"noncompliant":  WindowsManagedDeviceComplianceStatenoncompliant,
		"unknown":       WindowsManagedDeviceComplianceStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceComplianceState(input)
	return &out, nil
}
