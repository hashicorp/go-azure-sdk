package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceComplianceState string

const (
	ManagedDeviceComplianceStatecompliant     ManagedDeviceComplianceState = "Compliant"
	ManagedDeviceComplianceStateconfigManager ManagedDeviceComplianceState = "ConfigManager"
	ManagedDeviceComplianceStateconflict      ManagedDeviceComplianceState = "Conflict"
	ManagedDeviceComplianceStateerror         ManagedDeviceComplianceState = "Error"
	ManagedDeviceComplianceStateinGracePeriod ManagedDeviceComplianceState = "InGracePeriod"
	ManagedDeviceComplianceStatenoncompliant  ManagedDeviceComplianceState = "Noncompliant"
	ManagedDeviceComplianceStateunknown       ManagedDeviceComplianceState = "Unknown"
)

func PossibleValuesForManagedDeviceComplianceState() []string {
	return []string{
		string(ManagedDeviceComplianceStatecompliant),
		string(ManagedDeviceComplianceStateconfigManager),
		string(ManagedDeviceComplianceStateconflict),
		string(ManagedDeviceComplianceStateerror),
		string(ManagedDeviceComplianceStateinGracePeriod),
		string(ManagedDeviceComplianceStatenoncompliant),
		string(ManagedDeviceComplianceStateunknown),
	}
}

func parseManagedDeviceComplianceState(input string) (*ManagedDeviceComplianceState, error) {
	vals := map[string]ManagedDeviceComplianceState{
		"compliant":     ManagedDeviceComplianceStatecompliant,
		"configmanager": ManagedDeviceComplianceStateconfigManager,
		"conflict":      ManagedDeviceComplianceStateconflict,
		"error":         ManagedDeviceComplianceStateerror,
		"ingraceperiod": ManagedDeviceComplianceStateinGracePeriod,
		"noncompliant":  ManagedDeviceComplianceStatenoncompliant,
		"unknown":       ManagedDeviceComplianceStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceComplianceState(input)
	return &out, nil
}
