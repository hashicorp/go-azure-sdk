package clusterversion

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
