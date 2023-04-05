package sourcecontrolsyncjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCompleted ProvisioningState = "Completed"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateRunning   ProvisioningState = "Running"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCompleted),
		string(ProvisioningStateFailed),
		string(ProvisioningStateRunning),
	}
}

type SyncType string

const (
	SyncTypeFullSync    SyncType = "FullSync"
	SyncTypePartialSync SyncType = "PartialSync"
)

func PossibleValuesForSyncType() []string {
	return []string{
		string(SyncTypeFullSync),
		string(SyncTypePartialSync),
	}
}
