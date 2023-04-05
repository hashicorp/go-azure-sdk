package sqlvirtualmachinegroups

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

type ClusterManagerType string

const (
	ClusterManagerTypeWSFC ClusterManagerType = "WSFC"
)

func PossibleValuesForClusterManagerType() []string {
	return []string{
		string(ClusterManagerTypeWSFC),
	}
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

type ScaleType string

const (
	ScaleTypeHA ScaleType = "HA"
)

func PossibleValuesForScaleType() []string {
	return []string{
		string(ScaleTypeHA),
	}
}

type SqlVMGroupImageSku string

const (
	SqlVMGroupImageSkuDeveloper  SqlVMGroupImageSku = "Developer"
	SqlVMGroupImageSkuEnterprise SqlVMGroupImageSku = "Enterprise"
)

func PossibleValuesForSqlVMGroupImageSku() []string {
	return []string{
		string(SqlVMGroupImageSkuDeveloper),
		string(SqlVMGroupImageSkuEnterprise),
	}
}
