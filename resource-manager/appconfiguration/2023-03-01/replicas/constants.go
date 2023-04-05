package replicas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicaProvisioningState string

const (
	ReplicaProvisioningStateCanceled  ReplicaProvisioningState = "Canceled"
	ReplicaProvisioningStateCreating  ReplicaProvisioningState = "Creating"
	ReplicaProvisioningStateDeleting  ReplicaProvisioningState = "Deleting"
	ReplicaProvisioningStateFailed    ReplicaProvisioningState = "Failed"
	ReplicaProvisioningStateSucceeded ReplicaProvisioningState = "Succeeded"
)

func PossibleValuesForReplicaProvisioningState() []string {
	return []string{
		string(ReplicaProvisioningStateCanceled),
		string(ReplicaProvisioningStateCreating),
		string(ReplicaProvisioningStateDeleting),
		string(ReplicaProvisioningStateFailed),
		string(ReplicaProvisioningStateSucceeded),
	}
}
