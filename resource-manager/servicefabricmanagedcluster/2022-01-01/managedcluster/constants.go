package managedcluster

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Access string

const (
	AccessAllow Access = "allow"
	AccessDeny  Access = "deny"
)

func PossibleValuesForAccess() []string {
	return []string{
		string(AccessAllow),
		string(AccessDeny),
	}
}

type AddonFeatures string

const (
	AddonFeaturesBackupRestoreService   AddonFeatures = "BackupRestoreService"
	AddonFeaturesDnsService             AddonFeatures = "DnsService"
	AddonFeaturesResourceMonitorService AddonFeatures = "ResourceMonitorService"
)

func PossibleValuesForAddonFeatures() []string {
	return []string{
		string(AddonFeaturesBackupRestoreService),
		string(AddonFeaturesDnsService),
		string(AddonFeaturesResourceMonitorService),
	}
}

type ClusterState string

const (
	ClusterStateBaselineUpgrade ClusterState = "BaselineUpgrade"
	ClusterStateDeploying       ClusterState = "Deploying"
	ClusterStateReady           ClusterState = "Ready"
	ClusterStateUpgradeFailed   ClusterState = "UpgradeFailed"
	ClusterStateUpgrading       ClusterState = "Upgrading"
	ClusterStateWaitingForNodes ClusterState = "WaitingForNodes"
)

func PossibleValuesForClusterState() []string {
	return []string{
		string(ClusterStateBaselineUpgrade),
		string(ClusterStateDeploying),
		string(ClusterStateReady),
		string(ClusterStateUpgradeFailed),
		string(ClusterStateUpgrading),
		string(ClusterStateWaitingForNodes),
	}
}

type ClusterUpgradeCadence string

const (
	ClusterUpgradeCadenceWaveOne  ClusterUpgradeCadence = "Wave1"
	ClusterUpgradeCadenceWaveTwo  ClusterUpgradeCadence = "Wave2"
	ClusterUpgradeCadenceWaveZero ClusterUpgradeCadence = "Wave0"
)

func PossibleValuesForClusterUpgradeCadence() []string {
	return []string{
		string(ClusterUpgradeCadenceWaveOne),
		string(ClusterUpgradeCadenceWaveTwo),
		string(ClusterUpgradeCadenceWaveZero),
	}
}

type ClusterUpgradeMode string

const (
	ClusterUpgradeModeAutomatic ClusterUpgradeMode = "Automatic"
	ClusterUpgradeModeManual    ClusterUpgradeMode = "Manual"
)

func PossibleValuesForClusterUpgradeMode() []string {
	return []string{
		string(ClusterUpgradeModeAutomatic),
		string(ClusterUpgradeModeManual),
	}
}

type Direction string

const (
	DirectionInbound  Direction = "inbound"
	DirectionOutbound Direction = "outbound"
)

func PossibleValuesForDirection() []string {
	return []string{
		string(DirectionInbound),
		string(DirectionOutbound),
	}
}

type ManagedResourceProvisioningState string

const (
	ManagedResourceProvisioningStateCanceled  ManagedResourceProvisioningState = "Canceled"
	ManagedResourceProvisioningStateCreated   ManagedResourceProvisioningState = "Created"
	ManagedResourceProvisioningStateCreating  ManagedResourceProvisioningState = "Creating"
	ManagedResourceProvisioningStateDeleted   ManagedResourceProvisioningState = "Deleted"
	ManagedResourceProvisioningStateDeleting  ManagedResourceProvisioningState = "Deleting"
	ManagedResourceProvisioningStateFailed    ManagedResourceProvisioningState = "Failed"
	ManagedResourceProvisioningStateNone      ManagedResourceProvisioningState = "None"
	ManagedResourceProvisioningStateOther     ManagedResourceProvisioningState = "Other"
	ManagedResourceProvisioningStateSucceeded ManagedResourceProvisioningState = "Succeeded"
	ManagedResourceProvisioningStateUpdating  ManagedResourceProvisioningState = "Updating"
)

func PossibleValuesForManagedResourceProvisioningState() []string {
	return []string{
		string(ManagedResourceProvisioningStateCanceled),
		string(ManagedResourceProvisioningStateCreated),
		string(ManagedResourceProvisioningStateCreating),
		string(ManagedResourceProvisioningStateDeleted),
		string(ManagedResourceProvisioningStateDeleting),
		string(ManagedResourceProvisioningStateFailed),
		string(ManagedResourceProvisioningStateNone),
		string(ManagedResourceProvisioningStateOther),
		string(ManagedResourceProvisioningStateSucceeded),
		string(ManagedResourceProvisioningStateUpdating),
	}
}

type NsgProtocol string

const (
	NsgProtocolAh    NsgProtocol = "ah"
	NsgProtocolEsp   NsgProtocol = "esp"
	NsgProtocolHTTP  NsgProtocol = "http"
	NsgProtocolHTTPS NsgProtocol = "https"
	NsgProtocolIcmp  NsgProtocol = "icmp"
	NsgProtocolTcp   NsgProtocol = "tcp"
	NsgProtocolUdp   NsgProtocol = "udp"
)

func PossibleValuesForNsgProtocol() []string {
	return []string{
		string(NsgProtocolAh),
		string(NsgProtocolEsp),
		string(NsgProtocolHTTP),
		string(NsgProtocolHTTPS),
		string(NsgProtocolIcmp),
		string(NsgProtocolTcp),
		string(NsgProtocolUdp),
	}
}

type PrivateEndpointNetworkPolicies string

const (
	PrivateEndpointNetworkPoliciesDisabled PrivateEndpointNetworkPolicies = "disabled"
	PrivateEndpointNetworkPoliciesEnabled  PrivateEndpointNetworkPolicies = "enabled"
)

func PossibleValuesForPrivateEndpointNetworkPolicies() []string {
	return []string{
		string(PrivateEndpointNetworkPoliciesDisabled),
		string(PrivateEndpointNetworkPoliciesEnabled),
	}
}

type PrivateLinkServiceNetworkPolicies string

const (
	PrivateLinkServiceNetworkPoliciesDisabled PrivateLinkServiceNetworkPolicies = "disabled"
	PrivateLinkServiceNetworkPoliciesEnabled  PrivateLinkServiceNetworkPolicies = "enabled"
)

func PossibleValuesForPrivateLinkServiceNetworkPolicies() []string {
	return []string{
		string(PrivateLinkServiceNetworkPoliciesDisabled),
		string(PrivateLinkServiceNetworkPoliciesEnabled),
	}
}

type ProbeProtocol string

const (
	ProbeProtocolHTTP  ProbeProtocol = "http"
	ProbeProtocolHTTPS ProbeProtocol = "https"
	ProbeProtocolTcp   ProbeProtocol = "tcp"
)

func PossibleValuesForProbeProtocol() []string {
	return []string{
		string(ProbeProtocolHTTP),
		string(ProbeProtocolHTTPS),
		string(ProbeProtocolTcp),
	}
}

type Protocol string

const (
	ProtocolTcp Protocol = "tcp"
	ProtocolUdp Protocol = "udp"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolTcp),
		string(ProtocolUdp),
	}
}

type SkuName string

const (
	SkuNameBasic    SkuName = "Basic"
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameBasic),
		string(SkuNameStandard),
	}
}
