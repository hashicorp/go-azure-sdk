package workspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityReservationLevel int64

const (
	CapacityReservationLevelFiveHundred  CapacityReservationLevel = 500
	CapacityReservationLevelFiveThousand CapacityReservationLevel = 5000
	CapacityReservationLevelFourHundred  CapacityReservationLevel = 400
	CapacityReservationLevelOneHundred   CapacityReservationLevel = 100
	CapacityReservationLevelOneThousand  CapacityReservationLevel = 1000
	CapacityReservationLevelThreeHundred CapacityReservationLevel = 300
	CapacityReservationLevelTwoHundred   CapacityReservationLevel = 200
	CapacityReservationLevelTwoThousand  CapacityReservationLevel = 2000
)

func PossibleValuesForCapacityReservationLevel() []int64 {
	return []int64{
		int64(CapacityReservationLevelFiveHundred),
		int64(CapacityReservationLevelFiveThousand),
		int64(CapacityReservationLevelFourHundred),
		int64(CapacityReservationLevelOneHundred),
		int64(CapacityReservationLevelOneThousand),
		int64(CapacityReservationLevelThreeHundred),
		int64(CapacityReservationLevelTwoHundred),
		int64(CapacityReservationLevelTwoThousand),
	}
}

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
