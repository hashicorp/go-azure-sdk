package managedclusterversion

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedClusterVersionEnvironment string

const (
	ManagedClusterVersionEnvironmentWindows ManagedClusterVersionEnvironment = "Windows"
)

func PossibleValuesForManagedClusterVersionEnvironment() []string {
	return []string{
		string(ManagedClusterVersionEnvironmentWindows),
	}
}

func parseManagedClusterVersionEnvironment(input string) (*ManagedClusterVersionEnvironment, error) {
	vals := map[string]ManagedClusterVersionEnvironment{
		"windows": ManagedClusterVersionEnvironmentWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedClusterVersionEnvironment(input)
	return &out, nil
}

type OsType string

const (
	OsTypeWindows OsType = "Windows"
)

func PossibleValuesForOsType() []string {
	return []string{
		string(OsTypeWindows),
	}
}

func parseOsType(input string) (*OsType, error) {
	vals := map[string]OsType{
		"windows": OsTypeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OsType(input)
	return &out, nil
}
