package packetcaptures

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCaptureStatus string

const (
	PacketCaptureStatusError      PacketCaptureStatus = "Error"
	PacketCaptureStatusNotStarted PacketCaptureStatus = "NotStarted"
	PacketCaptureStatusRunning    PacketCaptureStatus = "Running"
	PacketCaptureStatusStopped    PacketCaptureStatus = "Stopped"
)

func PossibleValuesForPacketCaptureStatus() []string {
	return []string{
		string(PacketCaptureStatusError),
		string(PacketCaptureStatusNotStarted),
		string(PacketCaptureStatusRunning),
		string(PacketCaptureStatusStopped),
	}
}

func parsePacketCaptureStatus(input string) (*PacketCaptureStatus, error) {
	vals := map[string]PacketCaptureStatus{
		"error":      PacketCaptureStatusError,
		"notstarted": PacketCaptureStatusNotStarted,
		"running":    PacketCaptureStatusRunning,
		"stopped":    PacketCaptureStatusStopped,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PacketCaptureStatus(input)
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
