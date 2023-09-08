package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceExchangeAccessState string

const (
	WindowsManagedDeviceExchangeAccessStateallowed     WindowsManagedDeviceExchangeAccessState = "Allowed"
	WindowsManagedDeviceExchangeAccessStateblocked     WindowsManagedDeviceExchangeAccessState = "Blocked"
	WindowsManagedDeviceExchangeAccessStatenone        WindowsManagedDeviceExchangeAccessState = "None"
	WindowsManagedDeviceExchangeAccessStatequarantined WindowsManagedDeviceExchangeAccessState = "Quarantined"
	WindowsManagedDeviceExchangeAccessStateunknown     WindowsManagedDeviceExchangeAccessState = "Unknown"
)

func PossibleValuesForWindowsManagedDeviceExchangeAccessState() []string {
	return []string{
		string(WindowsManagedDeviceExchangeAccessStateallowed),
		string(WindowsManagedDeviceExchangeAccessStateblocked),
		string(WindowsManagedDeviceExchangeAccessStatenone),
		string(WindowsManagedDeviceExchangeAccessStatequarantined),
		string(WindowsManagedDeviceExchangeAccessStateunknown),
	}
}

func parseWindowsManagedDeviceExchangeAccessState(input string) (*WindowsManagedDeviceExchangeAccessState, error) {
	vals := map[string]WindowsManagedDeviceExchangeAccessState{
		"allowed":     WindowsManagedDeviceExchangeAccessStateallowed,
		"blocked":     WindowsManagedDeviceExchangeAccessStateblocked,
		"none":        WindowsManagedDeviceExchangeAccessStatenone,
		"quarantined": WindowsManagedDeviceExchangeAccessStatequarantined,
		"unknown":     WindowsManagedDeviceExchangeAccessStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceExchangeAccessState(input)
	return &out, nil
}
