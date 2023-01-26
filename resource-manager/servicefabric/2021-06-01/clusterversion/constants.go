package clusterversion

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterVersionsEnvironment string

const (
	ClusterVersionsEnvironmentLinux   ClusterVersionsEnvironment = "Linux"
	ClusterVersionsEnvironmentWindows ClusterVersionsEnvironment = "Windows"
)

func PossibleValuesForClusterVersionsEnvironment() []string {
	return []string{
		string(ClusterVersionsEnvironmentLinux),
		string(ClusterVersionsEnvironmentWindows),
	}
}

func parseClusterVersionsEnvironment(input string) (*ClusterVersionsEnvironment, error) {
	vals := map[string]ClusterVersionsEnvironment{
		"linux":   ClusterVersionsEnvironmentLinux,
		"windows": ClusterVersionsEnvironmentWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterVersionsEnvironment(input)
	return &out, nil
}

type Environment string

const (
	EnvironmentLinux   Environment = "Linux"
	EnvironmentWindows Environment = "Windows"
)

func PossibleValuesForEnvironment() []string {
	return []string{
		string(EnvironmentLinux),
		string(EnvironmentWindows),
	}
}

func parseEnvironment(input string) (*Environment, error) {
	vals := map[string]Environment{
		"linux":   EnvironmentLinux,
		"windows": EnvironmentWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Environment(input)
	return &out, nil
}
