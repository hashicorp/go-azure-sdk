package datastores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatastoreProvisioningState string

const (
	DatastoreProvisioningStateCanceled  DatastoreProvisioningState = "Canceled"
	DatastoreProvisioningStateCancelled DatastoreProvisioningState = "Cancelled"
	DatastoreProvisioningStateCreating  DatastoreProvisioningState = "Creating"
	DatastoreProvisioningStateDeleting  DatastoreProvisioningState = "Deleting"
	DatastoreProvisioningStateFailed    DatastoreProvisioningState = "Failed"
	DatastoreProvisioningStatePending   DatastoreProvisioningState = "Pending"
	DatastoreProvisioningStateSucceeded DatastoreProvisioningState = "Succeeded"
	DatastoreProvisioningStateUpdating  DatastoreProvisioningState = "Updating"
)

func PossibleValuesForDatastoreProvisioningState() []string {
	return []string{
		string(DatastoreProvisioningStateCanceled),
		string(DatastoreProvisioningStateCancelled),
		string(DatastoreProvisioningStateCreating),
		string(DatastoreProvisioningStateDeleting),
		string(DatastoreProvisioningStateFailed),
		string(DatastoreProvisioningStatePending),
		string(DatastoreProvisioningStateSucceeded),
		string(DatastoreProvisioningStateUpdating),
	}
}

type DatastoreStatus string

const (
	DatastoreStatusAccessible        DatastoreStatus = "Accessible"
	DatastoreStatusAttached          DatastoreStatus = "Attached"
	DatastoreStatusDeadOrError       DatastoreStatus = "DeadOrError"
	DatastoreStatusDetached          DatastoreStatus = "Detached"
	DatastoreStatusInaccessible      DatastoreStatus = "Inaccessible"
	DatastoreStatusLostCommunication DatastoreStatus = "LostCommunication"
	DatastoreStatusUnknown           DatastoreStatus = "Unknown"
)

func PossibleValuesForDatastoreStatus() []string {
	return []string{
		string(DatastoreStatusAccessible),
		string(DatastoreStatusAttached),
		string(DatastoreStatusDeadOrError),
		string(DatastoreStatusDetached),
		string(DatastoreStatusInaccessible),
		string(DatastoreStatusLostCommunication),
		string(DatastoreStatusUnknown),
	}
}

type MountOptionEnum string

const (
	MountOptionEnumATTACH MountOptionEnum = "ATTACH"
	MountOptionEnumMOUNT  MountOptionEnum = "MOUNT"
)

func PossibleValuesForMountOptionEnum() []string {
	return []string{
		string(MountOptionEnumATTACH),
		string(MountOptionEnumMOUNT),
	}
}
