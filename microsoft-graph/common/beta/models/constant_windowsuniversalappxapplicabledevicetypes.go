package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUniversalAppXApplicableDeviceTypes string

const (
	WindowsUniversalAppXApplicableDeviceTypesdesktop     WindowsUniversalAppXApplicableDeviceTypes = "Desktop"
	WindowsUniversalAppXApplicableDeviceTypesholographic WindowsUniversalAppXApplicableDeviceTypes = "Holographic"
	WindowsUniversalAppXApplicableDeviceTypesmobile      WindowsUniversalAppXApplicableDeviceTypes = "Mobile"
	WindowsUniversalAppXApplicableDeviceTypesnone        WindowsUniversalAppXApplicableDeviceTypes = "None"
	WindowsUniversalAppXApplicableDeviceTypesteam        WindowsUniversalAppXApplicableDeviceTypes = "Team"
)

func PossibleValuesForWindowsUniversalAppXApplicableDeviceTypes() []string {
	return []string{
		string(WindowsUniversalAppXApplicableDeviceTypesdesktop),
		string(WindowsUniversalAppXApplicableDeviceTypesholographic),
		string(WindowsUniversalAppXApplicableDeviceTypesmobile),
		string(WindowsUniversalAppXApplicableDeviceTypesnone),
		string(WindowsUniversalAppXApplicableDeviceTypesteam),
	}
}

func parseWindowsUniversalAppXApplicableDeviceTypes(input string) (*WindowsUniversalAppXApplicableDeviceTypes, error) {
	vals := map[string]WindowsUniversalAppXApplicableDeviceTypes{
		"desktop":     WindowsUniversalAppXApplicableDeviceTypesdesktop,
		"holographic": WindowsUniversalAppXApplicableDeviceTypesholographic,
		"mobile":      WindowsUniversalAppXApplicableDeviceTypesmobile,
		"none":        WindowsUniversalAppXApplicableDeviceTypesnone,
		"team":        WindowsUniversalAppXApplicableDeviceTypesteam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUniversalAppXApplicableDeviceTypes(input)
	return &out, nil
}
