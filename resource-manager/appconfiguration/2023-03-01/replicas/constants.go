package replicas

import "strings"

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

func parseReplicaProvisioningState(input string) (*ReplicaProvisioningState, error) {
	vals := map[string]ReplicaProvisioningState{
		"canceled":  ReplicaProvisioningStateCanceled,
		"creating":  ReplicaProvisioningStateCreating,
		"deleting":  ReplicaProvisioningStateDeleting,
		"failed":    ReplicaProvisioningStateFailed,
		"succeeded": ReplicaProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReplicaProvisioningState(input)
	return &out, nil
}
