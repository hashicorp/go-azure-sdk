package sourcecontrolsyncjob

import "strings"

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

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"completed": ProvisioningStateCompleted,
		"failed":    ProvisioningStateFailed,
		"running":   ProvisioningStateRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
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

func parseSyncType(input string) (*SyncType, error) {
	vals := map[string]SyncType{
		"fullsync":    SyncTypeFullSync,
		"partialsync": SyncTypePartialSync,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SyncType(input)
	return &out, nil
}
