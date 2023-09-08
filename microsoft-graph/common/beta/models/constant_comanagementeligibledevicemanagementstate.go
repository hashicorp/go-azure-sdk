package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComanagementEligibleDeviceManagementState string

const (
	ComanagementEligibleDeviceManagementStatedeletePending  ComanagementEligibleDeviceManagementState = "DeletePending"
	ComanagementEligibleDeviceManagementStatediscovered     ComanagementEligibleDeviceManagementState = "Discovered"
	ComanagementEligibleDeviceManagementStatemanaged        ComanagementEligibleDeviceManagementState = "Managed"
	ComanagementEligibleDeviceManagementStateretireCanceled ComanagementEligibleDeviceManagementState = "RetireCanceled"
	ComanagementEligibleDeviceManagementStateretireFailed   ComanagementEligibleDeviceManagementState = "RetireFailed"
	ComanagementEligibleDeviceManagementStateretireIssued   ComanagementEligibleDeviceManagementState = "RetireIssued"
	ComanagementEligibleDeviceManagementStateretirePending  ComanagementEligibleDeviceManagementState = "RetirePending"
	ComanagementEligibleDeviceManagementStateunhealthy      ComanagementEligibleDeviceManagementState = "Unhealthy"
	ComanagementEligibleDeviceManagementStatewipeCanceled   ComanagementEligibleDeviceManagementState = "WipeCanceled"
	ComanagementEligibleDeviceManagementStatewipeFailed     ComanagementEligibleDeviceManagementState = "WipeFailed"
	ComanagementEligibleDeviceManagementStatewipeIssued     ComanagementEligibleDeviceManagementState = "WipeIssued"
	ComanagementEligibleDeviceManagementStatewipePending    ComanagementEligibleDeviceManagementState = "WipePending"
)

func PossibleValuesForComanagementEligibleDeviceManagementState() []string {
	return []string{
		string(ComanagementEligibleDeviceManagementStatedeletePending),
		string(ComanagementEligibleDeviceManagementStatediscovered),
		string(ComanagementEligibleDeviceManagementStatemanaged),
		string(ComanagementEligibleDeviceManagementStateretireCanceled),
		string(ComanagementEligibleDeviceManagementStateretireFailed),
		string(ComanagementEligibleDeviceManagementStateretireIssued),
		string(ComanagementEligibleDeviceManagementStateretirePending),
		string(ComanagementEligibleDeviceManagementStateunhealthy),
		string(ComanagementEligibleDeviceManagementStatewipeCanceled),
		string(ComanagementEligibleDeviceManagementStatewipeFailed),
		string(ComanagementEligibleDeviceManagementStatewipeIssued),
		string(ComanagementEligibleDeviceManagementStatewipePending),
	}
}

func parseComanagementEligibleDeviceManagementState(input string) (*ComanagementEligibleDeviceManagementState, error) {
	vals := map[string]ComanagementEligibleDeviceManagementState{
		"deletepending":  ComanagementEligibleDeviceManagementStatedeletePending,
		"discovered":     ComanagementEligibleDeviceManagementStatediscovered,
		"managed":        ComanagementEligibleDeviceManagementStatemanaged,
		"retirecanceled": ComanagementEligibleDeviceManagementStateretireCanceled,
		"retirefailed":   ComanagementEligibleDeviceManagementStateretireFailed,
		"retireissued":   ComanagementEligibleDeviceManagementStateretireIssued,
		"retirepending":  ComanagementEligibleDeviceManagementStateretirePending,
		"unhealthy":      ComanagementEligibleDeviceManagementStateunhealthy,
		"wipecanceled":   ComanagementEligibleDeviceManagementStatewipeCanceled,
		"wipefailed":     ComanagementEligibleDeviceManagementStatewipeFailed,
		"wipeissued":     ComanagementEligibleDeviceManagementStatewipeIssued,
		"wipepending":    ComanagementEligibleDeviceManagementStatewipePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ComanagementEligibleDeviceManagementState(input)
	return &out, nil
}
