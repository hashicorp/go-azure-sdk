package redisenterprise

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessKeyType string

const (
	AccessKeyTypePrimary   AccessKeyType = "Primary"
	AccessKeyTypeSecondary AccessKeyType = "Secondary"
)

func PossibleValuesForAccessKeyType() []string {
	return []string{
		string(AccessKeyTypePrimary),
		string(AccessKeyTypeSecondary),
	}
}

type AofFrequency string

const (
	AofFrequencyAlways AofFrequency = "always"
	AofFrequencyOnes   AofFrequency = "1s"
)

func PossibleValuesForAofFrequency() []string {
	return []string{
		string(AofFrequencyAlways),
		string(AofFrequencyOnes),
	}
}

type ClusteringPolicy string

const (
	ClusteringPolicyEnterpriseCluster ClusteringPolicy = "EnterpriseCluster"
	ClusteringPolicyOSSCluster        ClusteringPolicy = "OSSCluster"
)

func PossibleValuesForClusteringPolicy() []string {
	return []string{
		string(ClusteringPolicyEnterpriseCluster),
		string(ClusteringPolicyOSSCluster),
	}
}

type EvictionPolicy string

const (
	EvictionPolicyAllKeysLFU     EvictionPolicy = "AllKeysLFU"
	EvictionPolicyAllKeysLRU     EvictionPolicy = "AllKeysLRU"
	EvictionPolicyAllKeysRandom  EvictionPolicy = "AllKeysRandom"
	EvictionPolicyNoEviction     EvictionPolicy = "NoEviction"
	EvictionPolicyVolatileLFU    EvictionPolicy = "VolatileLFU"
	EvictionPolicyVolatileLRU    EvictionPolicy = "VolatileLRU"
	EvictionPolicyVolatileRandom EvictionPolicy = "VolatileRandom"
	EvictionPolicyVolatileTTL    EvictionPolicy = "VolatileTTL"
)

func PossibleValuesForEvictionPolicy() []string {
	return []string{
		string(EvictionPolicyAllKeysLFU),
		string(EvictionPolicyAllKeysLRU),
		string(EvictionPolicyAllKeysRandom),
		string(EvictionPolicyNoEviction),
		string(EvictionPolicyVolatileLFU),
		string(EvictionPolicyVolatileLRU),
		string(EvictionPolicyVolatileRandom),
		string(EvictionPolicyVolatileTTL),
	}
}

type LinkState string

const (
	LinkStateLinkFailed   LinkState = "LinkFailed"
	LinkStateLinked       LinkState = "Linked"
	LinkStateLinking      LinkState = "Linking"
	LinkStateUnlinkFailed LinkState = "UnlinkFailed"
	LinkStateUnlinking    LinkState = "Unlinking"
)

func PossibleValuesForLinkState() []string {
	return []string{
		string(LinkStateLinkFailed),
		string(LinkStateLinked),
		string(LinkStateLinking),
		string(LinkStateUnlinkFailed),
		string(LinkStateUnlinking),
	}
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
	}
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}

type Protocol string

const (
	ProtocolEncrypted Protocol = "Encrypted"
	ProtocolPlaintext Protocol = "Plaintext"
)

func PossibleValuesForProtocol() []string {
	return []string{
		string(ProtocolEncrypted),
		string(ProtocolPlaintext),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type RdbFrequency string

const (
	RdbFrequencyOneTwoh RdbFrequency = "12h"
	RdbFrequencyOneh    RdbFrequency = "1h"
	RdbFrequencySixh    RdbFrequency = "6h"
)

func PossibleValuesForRdbFrequency() []string {
	return []string{
		string(RdbFrequencyOneTwoh),
		string(RdbFrequencyOneh),
		string(RdbFrequencySixh),
	}
}

type ResourceState string

const (
	ResourceStateCreateFailed  ResourceState = "CreateFailed"
	ResourceStateCreating      ResourceState = "Creating"
	ResourceStateDeleteFailed  ResourceState = "DeleteFailed"
	ResourceStateDeleting      ResourceState = "Deleting"
	ResourceStateDisableFailed ResourceState = "DisableFailed"
	ResourceStateDisabled      ResourceState = "Disabled"
	ResourceStateDisabling     ResourceState = "Disabling"
	ResourceStateEnableFailed  ResourceState = "EnableFailed"
	ResourceStateEnabling      ResourceState = "Enabling"
	ResourceStateRunning       ResourceState = "Running"
	ResourceStateUpdateFailed  ResourceState = "UpdateFailed"
	ResourceStateUpdating      ResourceState = "Updating"
)

func PossibleValuesForResourceState() []string {
	return []string{
		string(ResourceStateCreateFailed),
		string(ResourceStateCreating),
		string(ResourceStateDeleteFailed),
		string(ResourceStateDeleting),
		string(ResourceStateDisableFailed),
		string(ResourceStateDisabled),
		string(ResourceStateDisabling),
		string(ResourceStateEnableFailed),
		string(ResourceStateEnabling),
		string(ResourceStateRunning),
		string(ResourceStateUpdateFailed),
		string(ResourceStateUpdating),
	}
}

type SkuName string

const (
	SkuNameEnterpriseEFiveZero            SkuName = "Enterprise_E50"
	SkuNameEnterpriseEOneHundred          SkuName = "Enterprise_E100"
	SkuNameEnterpriseEOneZero             SkuName = "Enterprise_E10"
	SkuNameEnterpriseETwoZero             SkuName = "Enterprise_E20"
	SkuNameEnterpriseFlashFOneFiveHundred SkuName = "EnterpriseFlash_F1500"
	SkuNameEnterpriseFlashFSevenHundred   SkuName = "EnterpriseFlash_F700"
	SkuNameEnterpriseFlashFThreeHundred   SkuName = "EnterpriseFlash_F300"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameEnterpriseEFiveZero),
		string(SkuNameEnterpriseEOneHundred),
		string(SkuNameEnterpriseEOneZero),
		string(SkuNameEnterpriseETwoZero),
		string(SkuNameEnterpriseFlashFOneFiveHundred),
		string(SkuNameEnterpriseFlashFSevenHundred),
		string(SkuNameEnterpriseFlashFThreeHundred),
	}
}

type TlsVersion string

const (
	TlsVersionOnePointOne  TlsVersion = "1.1"
	TlsVersionOnePointTwo  TlsVersion = "1.2"
	TlsVersionOnePointZero TlsVersion = "1.0"
)

func PossibleValuesForTlsVersion() []string {
	return []string{
		string(TlsVersionOnePointOne),
		string(TlsVersionOnePointTwo),
		string(TlsVersionOnePointZero),
	}
}
