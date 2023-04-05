package services

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostingMode string

const (
	HostingModeDefault     HostingMode = "default"
	HostingModeHighDensity HostingMode = "highDensity"
)

func PossibleValuesForHostingMode() []string {
	return []string{
		string(HostingModeDefault),
		string(HostingModeHighDensity),
	}
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkServiceConnectionStatus() []string {
	return []string{
		string(PrivateLinkServiceConnectionStatusApproved),
		string(PrivateLinkServiceConnectionStatusDisconnected),
		string(PrivateLinkServiceConnectionStatusPending),
		string(PrivateLinkServiceConnectionStatusRejected),
	}
}

type ProvisioningState string

const (
	ProvisioningStateFailed       ProvisioningState = "failed"
	ProvisioningStateProvisioning ProvisioningState = "provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
	}
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

type ResourceType string

const (
	ResourceTypeSearchServices ResourceType = "searchServices"
)

func PossibleValuesForResourceType() []string {
	return []string{
		string(ResourceTypeSearchServices),
	}
}

type SearchServiceStatus string

const (
	SearchServiceStatusDegraded     SearchServiceStatus = "degraded"
	SearchServiceStatusDeleting     SearchServiceStatus = "deleting"
	SearchServiceStatusDisabled     SearchServiceStatus = "disabled"
	SearchServiceStatusError        SearchServiceStatus = "error"
	SearchServiceStatusProvisioning SearchServiceStatus = "provisioning"
	SearchServiceStatusRunning      SearchServiceStatus = "running"
)

func PossibleValuesForSearchServiceStatus() []string {
	return []string{
		string(SearchServiceStatusDegraded),
		string(SearchServiceStatusDeleting),
		string(SearchServiceStatusDisabled),
		string(SearchServiceStatusError),
		string(SearchServiceStatusProvisioning),
		string(SearchServiceStatusRunning),
	}
}

type SharedPrivateLinkResourceProvisioningState string

const (
	SharedPrivateLinkResourceProvisioningStateDeleting   SharedPrivateLinkResourceProvisioningState = "Deleting"
	SharedPrivateLinkResourceProvisioningStateFailed     SharedPrivateLinkResourceProvisioningState = "Failed"
	SharedPrivateLinkResourceProvisioningStateIncomplete SharedPrivateLinkResourceProvisioningState = "Incomplete"
	SharedPrivateLinkResourceProvisioningStateSucceeded  SharedPrivateLinkResourceProvisioningState = "Succeeded"
	SharedPrivateLinkResourceProvisioningStateUpdating   SharedPrivateLinkResourceProvisioningState = "Updating"
)

func PossibleValuesForSharedPrivateLinkResourceProvisioningState() []string {
	return []string{
		string(SharedPrivateLinkResourceProvisioningStateDeleting),
		string(SharedPrivateLinkResourceProvisioningStateFailed),
		string(SharedPrivateLinkResourceProvisioningStateIncomplete),
		string(SharedPrivateLinkResourceProvisioningStateSucceeded),
		string(SharedPrivateLinkResourceProvisioningStateUpdating),
	}
}

type SharedPrivateLinkResourceStatus string

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = "Approved"
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = "Pending"
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = "Rejected"
)

func PossibleValuesForSharedPrivateLinkResourceStatus() []string {
	return []string{
		string(SharedPrivateLinkResourceStatusApproved),
		string(SharedPrivateLinkResourceStatusDisconnected),
		string(SharedPrivateLinkResourceStatusPending),
		string(SharedPrivateLinkResourceStatusRejected),
	}
}

type SkuName string

const (
	SkuNameBasic                SkuName = "basic"
	SkuNameFree                 SkuName = "free"
	SkuNameStandard             SkuName = "standard"
	SkuNameStandardThree        SkuName = "standard3"
	SkuNameStandardTwo          SkuName = "standard2"
	SkuNameStorageOptimizedLOne SkuName = "storage_optimized_l1"
	SkuNameStorageOptimizedLTwo SkuName = "storage_optimized_l2"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameBasic),
		string(SkuNameFree),
		string(SkuNameStandard),
		string(SkuNameStandardThree),
		string(SkuNameStandardTwo),
		string(SkuNameStorageOptimizedLOne),
		string(SkuNameStorageOptimizedLTwo),
	}
}

type UnavailableNameReason string

const (
	UnavailableNameReasonAlreadyExists UnavailableNameReason = "AlreadyExists"
	UnavailableNameReasonInvalid       UnavailableNameReason = "Invalid"
)

func PossibleValuesForUnavailableNameReason() []string {
	return []string{
		string(UnavailableNameReasonAlreadyExists),
		string(UnavailableNameReasonInvalid),
	}
}
