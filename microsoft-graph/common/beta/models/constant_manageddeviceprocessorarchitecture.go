package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceProcessorArchitecture string

const (
	ManagedDeviceProcessorArchitecturearM64   ManagedDeviceProcessorArchitecture = "ArM64"
	ManagedDeviceProcessorArchitecturearm     ManagedDeviceProcessorArchitecture = "Arm"
	ManagedDeviceProcessorArchitectureunknown ManagedDeviceProcessorArchitecture = "Unknown"
	ManagedDeviceProcessorArchitecturex64     ManagedDeviceProcessorArchitecture = "X64"
	ManagedDeviceProcessorArchitecturex86     ManagedDeviceProcessorArchitecture = "X86"
)

func PossibleValuesForManagedDeviceProcessorArchitecture() []string {
	return []string{
		string(ManagedDeviceProcessorArchitecturearM64),
		string(ManagedDeviceProcessorArchitecturearm),
		string(ManagedDeviceProcessorArchitectureunknown),
		string(ManagedDeviceProcessorArchitecturex64),
		string(ManagedDeviceProcessorArchitecturex86),
	}
}

func parseManagedDeviceProcessorArchitecture(input string) (*ManagedDeviceProcessorArchitecture, error) {
	vals := map[string]ManagedDeviceProcessorArchitecture{
		"arm64":   ManagedDeviceProcessorArchitecturearM64,
		"arm":     ManagedDeviceProcessorArchitecturearm,
		"unknown": ManagedDeviceProcessorArchitectureunknown,
		"x64":     ManagedDeviceProcessorArchitecturex64,
		"x86":     ManagedDeviceProcessorArchitecturex86,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceProcessorArchitecture(input)
	return &out, nil
}
