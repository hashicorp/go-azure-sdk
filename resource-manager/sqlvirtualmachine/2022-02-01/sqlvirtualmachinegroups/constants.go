package sqlvirtualmachinegroups

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterConfiguration string

const (
	ClusterConfigurationDomainful ClusterConfiguration = "Domainful"
)

func PossibleValuesForClusterConfiguration() []string {
	return []string{
		string(ClusterConfigurationDomainful),
	}
}

func parseClusterConfiguration(input string) (*ClusterConfiguration, error) {
	vals := map[string]ClusterConfiguration{
		"domainful": ClusterConfigurationDomainful,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterConfiguration(input)
	return &out, nil
}

type ClusterManagerType string

const (
	ClusterManagerTypeWSFC ClusterManagerType = "WSFC"
)

func PossibleValuesForClusterManagerType() []string {
	return []string{
		string(ClusterManagerTypeWSFC),
	}
}

func parseClusterManagerType(input string) (*ClusterManagerType, error) {
	vals := map[string]ClusterManagerType{
		"wsfc": ClusterManagerTypeWSFC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterManagerType(input)
	return &out, nil
}

type ClusterSubnetType string

const (
	ClusterSubnetTypeMultiSubnet  ClusterSubnetType = "MultiSubnet"
	ClusterSubnetTypeSingleSubnet ClusterSubnetType = "SingleSubnet"
)

func PossibleValuesForClusterSubnetType() []string {
	return []string{
		string(ClusterSubnetTypeMultiSubnet),
		string(ClusterSubnetTypeSingleSubnet),
	}
}

func parseClusterSubnetType(input string) (*ClusterSubnetType, error) {
	vals := map[string]ClusterSubnetType{
		"multisubnet":  ClusterSubnetTypeMultiSubnet,
		"singlesubnet": ClusterSubnetTypeSingleSubnet,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ClusterSubnetType(input)
	return &out, nil
}

type ScaleType string

const (
	ScaleTypeHA ScaleType = "HA"
)

func PossibleValuesForScaleType() []string {
	return []string{
		string(ScaleTypeHA),
	}
}

func parseScaleType(input string) (*ScaleType, error) {
	vals := map[string]ScaleType{
		"ha": ScaleTypeHA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScaleType(input)
	return &out, nil
}

type SqlVmGroupImageSku string

const (
	SqlVmGroupImageSkuDeveloper  SqlVmGroupImageSku = "Developer"
	SqlVmGroupImageSkuEnterprise SqlVmGroupImageSku = "Enterprise"
)

func PossibleValuesForSqlVmGroupImageSku() []string {
	return []string{
		string(SqlVmGroupImageSkuDeveloper),
		string(SqlVmGroupImageSkuEnterprise),
	}
}

func parseSqlVmGroupImageSku(input string) (*SqlVmGroupImageSku, error) {
	vals := map[string]SqlVmGroupImageSku{
		"developer":  SqlVmGroupImageSkuDeveloper,
		"enterprise": SqlVmGroupImageSkuEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlVmGroupImageSku(input)
	return &out, nil
}
