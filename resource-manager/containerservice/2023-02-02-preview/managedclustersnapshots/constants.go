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
	ManagedClusterSKUNameBase ManagedClusterSKUName = "Base"
)

func PossibleValuesForManagedClusterSKUName() []string {
	return []string{
		string(ManagedClusterSKUNameBase),
	}
}

type ManagedClusterSKUTier string

const (
	ManagedClusterSKUTierFree     ManagedClusterSKUTier = "Free"
	ManagedClusterSKUTierStandard ManagedClusterSKUTier = "Standard"
)

func PossibleValuesForManagedClusterSKUTier() []string {
	return []string{
		string(ManagedClusterSKUTierFree),
		string(ManagedClusterSKUTierStandard),
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
	NetworkPolicyCilium NetworkPolicy = "cilium"
)

func PossibleValuesForNetworkPolicy() []string {
	return []string{
		string(NetworkPolicyAzure),
		string(NetworkPolicyCalico),
		string(NetworkPolicyCilium),
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
