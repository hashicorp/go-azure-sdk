package managedclustersnapshots

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LoadBalancerSku string

const (
	LoadBalancerSkuBasic    LoadBalancerSku = "basic"
	LoadBalancerSkuStandard LoadBalancerSku = "standard"
)

func PossibleValuesForLoadBalancerSku() []string {
	return []string{
		string(LoadBalancerSkuBasic),
		string(LoadBalancerSkuStandard),
	}
}

type ManagedClusterSKUName string

const (
	ManagedClusterSKUNameBasic ManagedClusterSKUName = "Basic"
)

func PossibleValuesForManagedClusterSKUName() []string {
	return []string{
		string(ManagedClusterSKUNameBasic),
	}
}

type ManagedClusterSKUTier string

const (
	ManagedClusterSKUTierFree ManagedClusterSKUTier = "Free"
	ManagedClusterSKUTierPaid ManagedClusterSKUTier = "Paid"
)

func PossibleValuesForManagedClusterSKUTier() []string {
	return []string{
		string(ManagedClusterSKUTierFree),
		string(ManagedClusterSKUTierPaid),
	}
}

type NetworkMode string

const (
	NetworkModeBridge      NetworkMode = "bridge"
	NetworkModeTransparent NetworkMode = "transparent"
)

func PossibleValuesForNetworkMode() []string {
	return []string{
		string(NetworkModeBridge),
		string(NetworkModeTransparent),
	}
}

type NetworkPlugin string

const (
	NetworkPluginAzure   NetworkPlugin = "azure"
	NetworkPluginKubenet NetworkPlugin = "kubenet"
	NetworkPluginNone    NetworkPlugin = "none"
)

func PossibleValuesForNetworkPlugin() []string {
	return []string{
		string(NetworkPluginAzure),
		string(NetworkPluginKubenet),
		string(NetworkPluginNone),
	}
}

type NetworkPluginMode string

const (
	NetworkPluginModeOverlay NetworkPluginMode = "Overlay"
)

func PossibleValuesForNetworkPluginMode() []string {
	return []string{
		string(NetworkPluginModeOverlay),
	}
}

type NetworkPolicy string

const (
	NetworkPolicyAzure  NetworkPolicy = "azure"
	NetworkPolicyCalico NetworkPolicy = "calico"
)

func PossibleValuesForNetworkPolicy() []string {
	return []string{
		string(NetworkPolicyAzure),
		string(NetworkPolicyCalico),
	}
}

type SnapshotType string

const (
	SnapshotTypeManagedCluster SnapshotType = "ManagedCluster"
	SnapshotTypeNodePool       SnapshotType = "NodePool"
)

func PossibleValuesForSnapshotType() []string {
	return []string{
		string(SnapshotTypeManagedCluster),
		string(SnapshotTypeNodePool),
	}
}
