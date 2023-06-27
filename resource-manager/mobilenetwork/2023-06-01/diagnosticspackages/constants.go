package diagnosticspackages

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsPackageStatus string

const (
	DiagnosticsPackageStatusCollected  DiagnosticsPackageStatus = "Collected"
	DiagnosticsPackageStatusCollecting DiagnosticsPackageStatus = "Collecting"
	DiagnosticsPackageStatusError      DiagnosticsPackageStatus = "Error"
	DiagnosticsPackageStatusNotStarted DiagnosticsPackageStatus = "NotStarted"
)

func PossibleValuesForDiagnosticsPackageStatus() []string {
	return []string{
		string(DiagnosticsPackageStatusCollected),
		string(DiagnosticsPackageStatusCollecting),
		string(DiagnosticsPackageStatusError),
		string(DiagnosticsPackageStatusNotStarted),
	}
}

func parseDiagnosticsPackageStatus(input string) (*DiagnosticsPackageStatus, error) {
	vals := map[string]DiagnosticsPackageStatus{
		"collected":  DiagnosticsPackageStatusCollected,
		"collecting": DiagnosticsPackageStatusCollecting,
		"error":      DiagnosticsPackageStatusError,
		"notstarted": DiagnosticsPackageStatusNotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DiagnosticsPackageStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":  ProvisioningStateAccepted,
		"canceled":  ProvisioningStateCanceled,
		"deleted":   ProvisioningStateDeleted,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"unknown":   ProvisioningStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
