package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceProcessorArchitecture string

const (
	WindowsManagedDeviceProcessorArchitecturearM64   WindowsManagedDeviceProcessorArchitecture = "ArM64"
	WindowsManagedDeviceProcessorArchitecturearm     WindowsManagedDeviceProcessorArchitecture = "Arm"
	WindowsManagedDeviceProcessorArchitectureunknown WindowsManagedDeviceProcessorArchitecture = "Unknown"
	WindowsManagedDeviceProcessorArchitecturex64     WindowsManagedDeviceProcessorArchitecture = "X64"
	WindowsManagedDeviceProcessorArchitecturex86     WindowsManagedDeviceProcessorArchitecture = "X86"
)

func PossibleValuesForWindowsManagedDeviceProcessorArchitecture() []string {
	return []string{
		string(WindowsManagedDeviceProcessorArchitecturearM64),
		string(WindowsManagedDeviceProcessorArchitecturearm),
		string(WindowsManagedDeviceProcessorArchitectureunknown),
		string(WindowsManagedDeviceProcessorArchitecturex64),
		string(WindowsManagedDeviceProcessorArchitecturex86),
	}
}

func parseWindowsManagedDeviceProcessorArchitecture(input string) (*WindowsManagedDeviceProcessorArchitecture, error) {
	vals := map[string]WindowsManagedDeviceProcessorArchitecture{
		"arm64":   WindowsManagedDeviceProcessorArchitecturearM64,
		"arm":     WindowsManagedDeviceProcessorArchitecturearm,
		"unknown": WindowsManagedDeviceProcessorArchitectureunknown,
		"x64":     WindowsManagedDeviceProcessorArchitecturex64,
		"x86":     WindowsManagedDeviceProcessorArchitecturex86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceProcessorArchitecture(input)
	return &out, nil
}
