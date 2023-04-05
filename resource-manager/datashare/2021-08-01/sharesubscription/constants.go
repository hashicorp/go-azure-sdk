package sharesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSetType string

const (
	DataSetTypeAdlsGenOneFile               DataSetType = "AdlsGen1File"
	DataSetTypeAdlsGenOneFolder             DataSetType = "AdlsGen1Folder"
	DataSetTypeAdlsGenTwoFile               DataSetType = "AdlsGen2File"
	DataSetTypeAdlsGenTwoFileSystem         DataSetType = "AdlsGen2FileSystem"
	DataSetTypeAdlsGenTwoFolder             DataSetType = "AdlsGen2Folder"
	DataSetTypeBlob                         DataSetType = "Blob"
	DataSetTypeBlobFolder                   DataSetType = "BlobFolder"
	DataSetTypeContainer                    DataSetType = "Container"
	DataSetTypeKustoCluster                 DataSetType = "KustoCluster"
	DataSetTypeKustoDatabase                DataSetType = "KustoDatabase"
	DataSetTypeKustoTable                   DataSetType = "KustoTable"
	DataSetTypeSqlDBTable                   DataSetType = "SqlDBTable"
	DataSetTypeSqlDWTable                   DataSetType = "SqlDWTable"
	DataSetTypeSynapseWorkspaceSqlPoolTable DataSetType = "SynapseWorkspaceSqlPoolTable"
)

func PossibleValuesForDataSetType() []string {
	return []string{
		string(DataSetTypeAdlsGenOneFile),
		string(DataSetTypeAdlsGenOneFolder),
		string(DataSetTypeAdlsGenTwoFile),
		string(DataSetTypeAdlsGenTwoFileSystem),
		string(DataSetTypeAdlsGenTwoFolder),
		string(DataSetTypeBlob),
		string(DataSetTypeBlobFolder),
		string(DataSetTypeContainer),
		string(DataSetTypeKustoCluster),
		string(DataSetTypeKustoDatabase),
		string(DataSetTypeKustoTable),
		string(DataSetTypeSqlDBTable),
		string(DataSetTypeSqlDWTable),
		string(DataSetTypeSynapseWorkspaceSqlPoolTable),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateSucceeded),
	}
}

type RecurrenceInterval string

const (
	RecurrenceIntervalDay  RecurrenceInterval = "Day"
	RecurrenceIntervalHour RecurrenceInterval = "Hour"
)

func PossibleValuesForRecurrenceInterval() []string {
	return []string{
		string(RecurrenceIntervalDay),
		string(RecurrenceIntervalHour),
	}
}

type ShareKind string

const (
	ShareKindCopyBased ShareKind = "CopyBased"
	ShareKindInPlace   ShareKind = "InPlace"
)

func PossibleValuesForShareKind() []string {
	return []string{
		string(ShareKindCopyBased),
		string(ShareKindInPlace),
	}
}

type ShareSubscriptionStatus string

const (
	ShareSubscriptionStatusActive        ShareSubscriptionStatus = "Active"
	ShareSubscriptionStatusRevoked       ShareSubscriptionStatus = "Revoked"
	ShareSubscriptionStatusRevoking      ShareSubscriptionStatus = "Revoking"
	ShareSubscriptionStatusSourceDeleted ShareSubscriptionStatus = "SourceDeleted"
)

func PossibleValuesForShareSubscriptionStatus() []string {
	return []string{
		string(ShareSubscriptionStatusActive),
		string(ShareSubscriptionStatusRevoked),
		string(ShareSubscriptionStatusRevoking),
		string(ShareSubscriptionStatusSourceDeleted),
	}
}

type SourceShareSynchronizationSettingKind string

const (
	SourceShareSynchronizationSettingKindScheduleBased SourceShareSynchronizationSettingKind = "ScheduleBased"
)

func PossibleValuesForSourceShareSynchronizationSettingKind() []string {
	return []string{
		string(SourceShareSynchronizationSettingKindScheduleBased),
	}
}

type Status string

const (
	StatusAccepted         Status = "Accepted"
	StatusCanceled         Status = "Canceled"
	StatusFailed           Status = "Failed"
	StatusInProgress       Status = "InProgress"
	StatusSucceeded        Status = "Succeeded"
	StatusTransientFailure Status = "TransientFailure"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusAccepted),
		string(StatusCanceled),
		string(StatusFailed),
		string(StatusInProgress),
		string(StatusSucceeded),
		string(StatusTransientFailure),
	}
}

type SynchronizationMode string

const (
	SynchronizationModeFullSync    SynchronizationMode = "FullSync"
	SynchronizationModeIncremental SynchronizationMode = "Incremental"
)

func PossibleValuesForSynchronizationMode() []string {
	return []string{
		string(SynchronizationModeFullSync),
		string(SynchronizationModeIncremental),
	}
}
