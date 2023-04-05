package timeseriesdatabaseconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CleanupConnectionArtifacts string

const (
	CleanupConnectionArtifactsFalse CleanupConnectionArtifacts = "false"
	CleanupConnectionArtifactsTrue  CleanupConnectionArtifacts = "true"
)

func PossibleValuesForCleanupConnectionArtifacts() []string {
	return []string{
		string(CleanupConnectionArtifactsFalse),
		string(CleanupConnectionArtifactsTrue),
	}
}

type ConnectionType string

const (
	ConnectionTypeAzureDataExplorer ConnectionType = "AzureDataExplorer"
)

func PossibleValuesForConnectionType() []string {
	return []string{
		string(ConnectionTypeAzureDataExplorer),
	}
}

type IdentityType string

const (
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
	IdentityTypeUserAssigned   IdentityType = "UserAssigned"
)

func PossibleValuesForIdentityType() []string {
	return []string{
		string(IdentityTypeSystemAssigned),
		string(IdentityTypeUserAssigned),
	}
}

type RecordPropertyAndItemRemovals string

const (
	RecordPropertyAndItemRemovalsFalse RecordPropertyAndItemRemovals = "false"
	RecordPropertyAndItemRemovalsTrue  RecordPropertyAndItemRemovals = "true"
)

func PossibleValuesForRecordPropertyAndItemRemovals() []string {
	return []string{
		string(RecordPropertyAndItemRemovalsFalse),
		string(RecordPropertyAndItemRemovalsTrue),
	}
}

type TimeSeriesDatabaseConnectionState string

const (
	TimeSeriesDatabaseConnectionStateCanceled     TimeSeriesDatabaseConnectionState = "Canceled"
	TimeSeriesDatabaseConnectionStateDeleted      TimeSeriesDatabaseConnectionState = "Deleted"
	TimeSeriesDatabaseConnectionStateDeleting     TimeSeriesDatabaseConnectionState = "Deleting"
	TimeSeriesDatabaseConnectionStateDisabled     TimeSeriesDatabaseConnectionState = "Disabled"
	TimeSeriesDatabaseConnectionStateFailed       TimeSeriesDatabaseConnectionState = "Failed"
	TimeSeriesDatabaseConnectionStateMoving       TimeSeriesDatabaseConnectionState = "Moving"
	TimeSeriesDatabaseConnectionStateProvisioning TimeSeriesDatabaseConnectionState = "Provisioning"
	TimeSeriesDatabaseConnectionStateRestoring    TimeSeriesDatabaseConnectionState = "Restoring"
	TimeSeriesDatabaseConnectionStateSucceeded    TimeSeriesDatabaseConnectionState = "Succeeded"
	TimeSeriesDatabaseConnectionStateSuspending   TimeSeriesDatabaseConnectionState = "Suspending"
	TimeSeriesDatabaseConnectionStateUpdating     TimeSeriesDatabaseConnectionState = "Updating"
	TimeSeriesDatabaseConnectionStateWarning      TimeSeriesDatabaseConnectionState = "Warning"
)

func PossibleValuesForTimeSeriesDatabaseConnectionState() []string {
	return []string{
		string(TimeSeriesDatabaseConnectionStateCanceled),
		string(TimeSeriesDatabaseConnectionStateDeleted),
		string(TimeSeriesDatabaseConnectionStateDeleting),
		string(TimeSeriesDatabaseConnectionStateDisabled),
		string(TimeSeriesDatabaseConnectionStateFailed),
		string(TimeSeriesDatabaseConnectionStateMoving),
		string(TimeSeriesDatabaseConnectionStateProvisioning),
		string(TimeSeriesDatabaseConnectionStateRestoring),
		string(TimeSeriesDatabaseConnectionStateSucceeded),
		string(TimeSeriesDatabaseConnectionStateSuspending),
		string(TimeSeriesDatabaseConnectionStateUpdating),
		string(TimeSeriesDatabaseConnectionStateWarning),
	}
}
