package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceManagementState string

const (
	ManagedDeviceManagementStatedeletePending  ManagedDeviceManagementState = "DeletePending"
	ManagedDeviceManagementStatediscovered     ManagedDeviceManagementState = "Discovered"
	ManagedDeviceManagementStatemanaged        ManagedDeviceManagementState = "Managed"
	ManagedDeviceManagementStateretireCanceled ManagedDeviceManagementState = "RetireCanceled"
	ManagedDeviceManagementStateretireFailed   ManagedDeviceManagementState = "RetireFailed"
	ManagedDeviceManagementStateretireIssued   ManagedDeviceManagementState = "RetireIssued"
	ManagedDeviceManagementStateretirePending  ManagedDeviceManagementState = "RetirePending"
	ManagedDeviceManagementStateunhealthy      ManagedDeviceManagementState = "Unhealthy"
	ManagedDeviceManagementStatewipeCanceled   ManagedDeviceManagementState = "WipeCanceled"
	ManagedDeviceManagementStatewipeFailed     ManagedDeviceManagementState = "WipeFailed"
	ManagedDeviceManagementStatewipeIssued     ManagedDeviceManagementState = "WipeIssued"
	ManagedDeviceManagementStatewipePending    ManagedDeviceManagementState = "WipePending"
)

func PossibleValuesForManagedDeviceManagementState() []string {
	return []string{
		string(ManagedDeviceManagementStatedeletePending),
		string(ManagedDeviceManagementStatediscovered),
		string(ManagedDeviceManagementStatemanaged),
		string(ManagedDeviceManagementStateretireCanceled),
		string(ManagedDeviceManagementStateretireFailed),
		string(ManagedDeviceManagementStateretireIssued),
		string(ManagedDeviceManagementStateretirePending),
		string(ManagedDeviceManagementStateunhealthy),
		string(ManagedDeviceManagementStatewipeCanceled),
		string(ManagedDeviceManagementStatewipeFailed),
		string(ManagedDeviceManagementStatewipeIssued),
		string(ManagedDeviceManagementStatewipePending),
	}
}

func parseManagedDeviceManagementState(input string) (*ManagedDeviceManagementState, error) {
	vals := map[string]ManagedDeviceManagementState{
		"deletepending":  ManagedDeviceManagementStatedeletePending,
		"discovered":     ManagedDeviceManagementStatediscovered,
		"managed":        ManagedDeviceManagementStatemanaged,
		"retirecanceled": ManagedDeviceManagementStateretireCanceled,
		"retirefailed":   ManagedDeviceManagementStateretireFailed,
		"retireissued":   ManagedDeviceManagementStateretireIssued,
		"retirepending":  ManagedDeviceManagementStateretirePending,
		"unhealthy":      ManagedDeviceManagementStateunhealthy,
		"wipecanceled":   ManagedDeviceManagementStatewipeCanceled,
		"wipefailed":     ManagedDeviceManagementStatewipeFailed,
		"wipeissued":     ManagedDeviceManagementStatewipeIssued,
		"wipepending":    ManagedDeviceManagementStatewipePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceManagementState(input)
	return &out, nil
}
