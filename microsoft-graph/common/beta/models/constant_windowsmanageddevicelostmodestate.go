package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceLostModeState string

const (
	WindowsManagedDeviceLostModeStatedisabled WindowsManagedDeviceLostModeState = "Disabled"
	WindowsManagedDeviceLostModeStateenabled  WindowsManagedDeviceLostModeState = "Enabled"
)

func PossibleValuesForWindowsManagedDeviceLostModeState() []string {
	return []string{
		string(WindowsManagedDeviceLostModeStatedisabled),
		string(WindowsManagedDeviceLostModeStateenabled),
	}
}

func parseWindowsManagedDeviceLostModeState(input string) (*WindowsManagedDeviceLostModeState, error) {
	vals := map[string]WindowsManagedDeviceLostModeState{
		"disabled": WindowsManagedDeviceLostModeStatedisabled,
		"enabled":  WindowsManagedDeviceLostModeStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceLostModeState(input)
	return &out, nil
}
