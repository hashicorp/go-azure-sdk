package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceExchangeAccessState string

const (
	ManagedDeviceExchangeAccessStateallowed     ManagedDeviceExchangeAccessState = "Allowed"
	ManagedDeviceExchangeAccessStateblocked     ManagedDeviceExchangeAccessState = "Blocked"
	ManagedDeviceExchangeAccessStatenone        ManagedDeviceExchangeAccessState = "None"
	ManagedDeviceExchangeAccessStatequarantined ManagedDeviceExchangeAccessState = "Quarantined"
	ManagedDeviceExchangeAccessStateunknown     ManagedDeviceExchangeAccessState = "Unknown"
)

func PossibleValuesForManagedDeviceExchangeAccessState() []string {
	return []string{
		string(ManagedDeviceExchangeAccessStateallowed),
		string(ManagedDeviceExchangeAccessStateblocked),
		string(ManagedDeviceExchangeAccessStatenone),
		string(ManagedDeviceExchangeAccessStatequarantined),
		string(ManagedDeviceExchangeAccessStateunknown),
	}
}

func parseManagedDeviceExchangeAccessState(input string) (*ManagedDeviceExchangeAccessState, error) {
	vals := map[string]ManagedDeviceExchangeAccessState{
		"allowed":     ManagedDeviceExchangeAccessStateallowed,
		"blocked":     ManagedDeviceExchangeAccessStateblocked,
		"none":        ManagedDeviceExchangeAccessStatenone,
		"quarantined": ManagedDeviceExchangeAccessStatequarantined,
		"unknown":     ManagedDeviceExchangeAccessStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceExchangeAccessState(input)
	return &out, nil
}
