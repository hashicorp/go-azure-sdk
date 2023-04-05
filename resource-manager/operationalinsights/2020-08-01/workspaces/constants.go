package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataIngestionStatus string

const (
	DataIngestionStatusApproachingQuota      DataIngestionStatus = "ApproachingQuota"
	DataIngestionStatusForceOff              DataIngestionStatus = "ForceOff"
	DataIngestionStatusForceOn               DataIngestionStatus = "ForceOn"
	DataIngestionStatusOverQuota             DataIngestionStatus = "OverQuota"
	DataIngestionStatusRespectQuota          DataIngestionStatus = "RespectQuota"
	DataIngestionStatusSubscriptionSuspended DataIngestionStatus = "SubscriptionSuspended"
)

func PossibleValuesForDataIngestionStatus() []string {
	return []string{
		string(DataIngestionStatusApproachingQuota),
		string(DataIngestionStatusForceOff),
		string(DataIngestionStatusForceOn),
		string(DataIngestionStatusOverQuota),
		string(DataIngestionStatusRespectQuota),
		string(DataIngestionStatusSubscriptionSuspended),
	}
}

type PublicNetworkAccessType string

const (
	PublicNetworkAccessTypeDisabled PublicNetworkAccessType = "Disabled"
	PublicNetworkAccessTypeEnabled  PublicNetworkAccessType = "Enabled"
)

func PossibleValuesForPublicNetworkAccessType() []string {
	return []string{
		string(PublicNetworkAccessTypeDisabled),
		string(PublicNetworkAccessTypeEnabled),
	}
}

type PurgeState string

const (
	PurgeStateCompleted PurgeState = "completed"
	PurgeStatePending   PurgeState = "pending"
)

func PossibleValuesForPurgeState() []string {
	return []string{
		string(PurgeStateCompleted),
		string(PurgeStatePending),
	}
}

type SearchSortEnum string

const (
	SearchSortEnumAsc  SearchSortEnum = "asc"
	SearchSortEnumDesc SearchSortEnum = "desc"
)

func PossibleValuesForSearchSortEnum() []string {
	return []string{
		string(SearchSortEnumAsc),
		string(SearchSortEnumDesc),
	}
}

type WorkspaceEntityStatus string

const (
	WorkspaceEntityStatusCanceled            WorkspaceEntityStatus = "Canceled"
	WorkspaceEntityStatusCreating            WorkspaceEntityStatus = "Creating"
	WorkspaceEntityStatusDeleting            WorkspaceEntityStatus = "Deleting"
	WorkspaceEntityStatusFailed              WorkspaceEntityStatus = "Failed"
	WorkspaceEntityStatusProvisioningAccount WorkspaceEntityStatus = "ProvisioningAccount"
	WorkspaceEntityStatusSucceeded           WorkspaceEntityStatus = "Succeeded"
	WorkspaceEntityStatusUpdating            WorkspaceEntityStatus = "Updating"
)

func PossibleValuesForWorkspaceEntityStatus() []string {
	return []string{
		string(WorkspaceEntityStatusCanceled),
		string(WorkspaceEntityStatusCreating),
		string(WorkspaceEntityStatusDeleting),
		string(WorkspaceEntityStatusFailed),
		string(WorkspaceEntityStatusProvisioningAccount),
		string(WorkspaceEntityStatusSucceeded),
		string(WorkspaceEntityStatusUpdating),
	}
}

type WorkspaceSkuNameEnum string

const (
	WorkspaceSkuNameEnumCapacityReservation  WorkspaceSkuNameEnum = "CapacityReservation"
	WorkspaceSkuNameEnumFree                 WorkspaceSkuNameEnum = "Free"
	WorkspaceSkuNameEnumLACluster            WorkspaceSkuNameEnum = "LACluster"
	WorkspaceSkuNameEnumPerGBTwoZeroOneEight WorkspaceSkuNameEnum = "PerGB2018"
	WorkspaceSkuNameEnumPerNode              WorkspaceSkuNameEnum = "PerNode"
	WorkspaceSkuNameEnumPremium              WorkspaceSkuNameEnum = "Premium"
	WorkspaceSkuNameEnumStandalone           WorkspaceSkuNameEnum = "Standalone"
	WorkspaceSkuNameEnumStandard             WorkspaceSkuNameEnum = "Standard"
)

func PossibleValuesForWorkspaceSkuNameEnum() []string {
	return []string{
		string(WorkspaceSkuNameEnumCapacityReservation),
		string(WorkspaceSkuNameEnumFree),
		string(WorkspaceSkuNameEnumLACluster),
		string(WorkspaceSkuNameEnumPerGBTwoZeroOneEight),
		string(WorkspaceSkuNameEnumPerNode),
		string(WorkspaceSkuNameEnumPremium),
		string(WorkspaceSkuNameEnumStandalone),
		string(WorkspaceSkuNameEnumStandard),
	}
}
