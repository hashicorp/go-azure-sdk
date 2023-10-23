package mhsmlistprivateendpointconnections

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionsRequired string

const (
	ActionsRequiredNone ActionsRequired = "None"
)

func PossibleValuesForActionsRequired() []string {
	return []string{
		string(ActionsRequiredNone),
	}
}

func parseActionsRequired(input string) (*ActionsRequired, error) {
	vals := map[string]ActionsRequired{
		"none": ActionsRequiredNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActionsRequired(input)
	return &out, nil
}

type ManagedHsmSkuFamily string

const (
	ManagedHsmSkuFamilyB ManagedHsmSkuFamily = "B"
)

func PossibleValuesForManagedHsmSkuFamily() []string {
	return []string{
		string(ManagedHsmSkuFamilyB),
	}
}

func parseManagedHsmSkuFamily(input string) (*ManagedHsmSkuFamily, error) {
	vals := map[string]ManagedHsmSkuFamily{
		"b": ManagedHsmSkuFamilyB,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedHsmSkuFamily(input)
	return &out, nil
}

type ManagedHsmSkuName string

const (
	ManagedHsmSkuNameCustomBSix      ManagedHsmSkuName = "Custom_B6"
	ManagedHsmSkuNameCustomBThreeTwo ManagedHsmSkuName = "Custom_B32"
	ManagedHsmSkuNameStandardBOne    ManagedHsmSkuName = "Standard_B1"
)

func PossibleValuesForManagedHsmSkuName() []string {
	return []string{
		string(ManagedHsmSkuNameCustomBSix),
		string(ManagedHsmSkuNameCustomBThreeTwo),
		string(ManagedHsmSkuNameStandardBOne),
	}
}

func parseManagedHsmSkuName(input string) (*ManagedHsmSkuName, error) {
	vals := map[string]ManagedHsmSkuName{
		"custom_b6":   ManagedHsmSkuNameCustomBSix,
		"custom_b32":  ManagedHsmSkuNameCustomBThreeTwo,
		"standard_b1": ManagedHsmSkuNameStandardBOne,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedHsmSkuName(input)
	return &out, nil
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating     PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting     PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateDisconnected PrivateEndpointConnectionProvisioningState = "Disconnected"
	PrivateEndpointConnectionProvisioningStateFailed       PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded    PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating     PrivateEndpointConnectionProvisioningState = "Updating"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateDisconnected),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
		string(PrivateEndpointConnectionProvisioningStateUpdating),
	}
}

func parsePrivateEndpointConnectionProvisioningState(input string) (*PrivateEndpointConnectionProvisioningState, error) {
	vals := map[string]PrivateEndpointConnectionProvisioningState{
		"creating":     PrivateEndpointConnectionProvisioningStateCreating,
		"deleting":     PrivateEndpointConnectionProvisioningStateDeleting,
		"disconnected": PrivateEndpointConnectionProvisioningStateDisconnected,
		"failed":       PrivateEndpointConnectionProvisioningStateFailed,
		"succeeded":    PrivateEndpointConnectionProvisioningStateSucceeded,
		"updating":     PrivateEndpointConnectionProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointConnectionProvisioningState(input)
	return &out, nil
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusDisconnected),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}

func parsePrivateEndpointServiceConnectionStatus(input string) (*PrivateEndpointServiceConnectionStatus, error) {
	vals := map[string]PrivateEndpointServiceConnectionStatus{
		"approved":     PrivateEndpointServiceConnectionStatusApproved,
		"disconnected": PrivateEndpointServiceConnectionStatusDisconnected,
		"pending":      PrivateEndpointServiceConnectionStatusPending,
		"rejected":     PrivateEndpointServiceConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointServiceConnectionStatus(input)
	return &out, nil
}
