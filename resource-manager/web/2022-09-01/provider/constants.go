package provider

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderOsTypeSelected string

const (
	ProviderOsTypeSelectedAll              ProviderOsTypeSelected = "All"
	ProviderOsTypeSelectedLinux            ProviderOsTypeSelected = "Linux"
	ProviderOsTypeSelectedLinuxFunctions   ProviderOsTypeSelected = "LinuxFunctions"
	ProviderOsTypeSelectedWindows          ProviderOsTypeSelected = "Windows"
	ProviderOsTypeSelectedWindowsFunctions ProviderOsTypeSelected = "WindowsFunctions"
)

func PossibleValuesForProviderOsTypeSelected() []string {
	return []string{
		string(ProviderOsTypeSelectedAll),
		string(ProviderOsTypeSelectedLinux),
		string(ProviderOsTypeSelectedLinuxFunctions),
		string(ProviderOsTypeSelectedWindows),
		string(ProviderOsTypeSelectedWindowsFunctions),
	}
}

func parseProviderOsTypeSelected(input string) (*ProviderOsTypeSelected, error) {
	vals := map[string]ProviderOsTypeSelected{
		"all":              ProviderOsTypeSelectedAll,
		"linux":            ProviderOsTypeSelectedLinux,
		"linuxfunctions":   ProviderOsTypeSelectedLinuxFunctions,
		"windows":          ProviderOsTypeSelectedWindows,
		"windowsfunctions": ProviderOsTypeSelectedWindowsFunctions,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProviderOsTypeSelected(input)
	return &out, nil
}

type ProviderStackOsType string

const (
	ProviderStackOsTypeAll     ProviderStackOsType = "All"
	ProviderStackOsTypeLinux   ProviderStackOsType = "Linux"
	ProviderStackOsTypeWindows ProviderStackOsType = "Windows"
)

func PossibleValuesForProviderStackOsType() []string {
	return []string{
		string(ProviderStackOsTypeAll),
		string(ProviderStackOsTypeLinux),
		string(ProviderStackOsTypeWindows),
	}
}

func parseProviderStackOsType(input string) (*ProviderStackOsType, error) {
	vals := map[string]ProviderStackOsType{
		"all":     ProviderStackOsTypeAll,
		"linux":   ProviderStackOsTypeLinux,
		"windows": ProviderStackOsTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProviderStackOsType(input)
	return &out, nil
}

type StackPreferredOs string

const (
	StackPreferredOsLinux   StackPreferredOs = "Linux"
	StackPreferredOsWindows StackPreferredOs = "Windows"
)

func PossibleValuesForStackPreferredOs() []string {
	return []string{
		string(StackPreferredOsLinux),
		string(StackPreferredOsWindows),
	}
}

func parseStackPreferredOs(input string) (*StackPreferredOs, error) {
	vals := map[string]StackPreferredOs{
		"linux":   StackPreferredOsLinux,
		"windows": StackPreferredOsWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StackPreferredOs(input)
	return &out, nil
}
