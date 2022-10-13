package extensions

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionAggregateState string

const (
	ExtensionAggregateStateCanceled           ExtensionAggregateState = "Canceled"
	ExtensionAggregateStateConnected          ExtensionAggregateState = "Connected"
	ExtensionAggregateStateCreating           ExtensionAggregateState = "Creating"
	ExtensionAggregateStateDeleted            ExtensionAggregateState = "Deleted"
	ExtensionAggregateStateDeleting           ExtensionAggregateState = "Deleting"
	ExtensionAggregateStateDisconnected       ExtensionAggregateState = "Disconnected"
	ExtensionAggregateStateError              ExtensionAggregateState = "Error"
	ExtensionAggregateStateFailed             ExtensionAggregateState = "Failed"
	ExtensionAggregateStateInProgress         ExtensionAggregateState = "InProgress"
	ExtensionAggregateStateMoving             ExtensionAggregateState = "Moving"
	ExtensionAggregateStateNotSpecified       ExtensionAggregateState = "NotSpecified"
	ExtensionAggregateStatePartiallyConnected ExtensionAggregateState = "PartiallyConnected"
	ExtensionAggregateStatePartiallySucceeded ExtensionAggregateState = "PartiallySucceeded"
	ExtensionAggregateStateSucceeded          ExtensionAggregateState = "Succeeded"
	ExtensionAggregateStateUpdating           ExtensionAggregateState = "Updating"
)

func PossibleValuesForExtensionAggregateState() []string {
	return []string{
		string(ExtensionAggregateStateCanceled),
		string(ExtensionAggregateStateConnected),
		string(ExtensionAggregateStateCreating),
		string(ExtensionAggregateStateDeleted),
		string(ExtensionAggregateStateDeleting),
		string(ExtensionAggregateStateDisconnected),
		string(ExtensionAggregateStateError),
		string(ExtensionAggregateStateFailed),
		string(ExtensionAggregateStateInProgress),
		string(ExtensionAggregateStateMoving),
		string(ExtensionAggregateStateNotSpecified),
		string(ExtensionAggregateStatePartiallyConnected),
		string(ExtensionAggregateStatePartiallySucceeded),
		string(ExtensionAggregateStateSucceeded),
		string(ExtensionAggregateStateUpdating),
	}
}

func parseExtensionAggregateState(input string) (*ExtensionAggregateState, error) {
	vals := map[string]ExtensionAggregateState{
		"canceled":           ExtensionAggregateStateCanceled,
		"connected":          ExtensionAggregateStateConnected,
		"creating":           ExtensionAggregateStateCreating,
		"deleted":            ExtensionAggregateStateDeleted,
		"deleting":           ExtensionAggregateStateDeleting,
		"disconnected":       ExtensionAggregateStateDisconnected,
		"error":              ExtensionAggregateStateError,
		"failed":             ExtensionAggregateStateFailed,
		"inprogress":         ExtensionAggregateStateInProgress,
		"moving":             ExtensionAggregateStateMoving,
		"notspecified":       ExtensionAggregateStateNotSpecified,
		"partiallyconnected": ExtensionAggregateStatePartiallyConnected,
		"partiallysucceeded": ExtensionAggregateStatePartiallySucceeded,
		"succeeded":          ExtensionAggregateStateSucceeded,
		"updating":           ExtensionAggregateStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExtensionAggregateState(input)
	return &out, nil
}

type NodeExtensionState string

const (
	NodeExtensionStateCanceled     NodeExtensionState = "Canceled"
	NodeExtensionStateConnected    NodeExtensionState = "Connected"
	NodeExtensionStateCreating     NodeExtensionState = "Creating"
	NodeExtensionStateDeleted      NodeExtensionState = "Deleted"
	NodeExtensionStateDeleting     NodeExtensionState = "Deleting"
	NodeExtensionStateDisconnected NodeExtensionState = "Disconnected"
	NodeExtensionStateError        NodeExtensionState = "Error"
	NodeExtensionStateFailed       NodeExtensionState = "Failed"
	NodeExtensionStateMoving       NodeExtensionState = "Moving"
	NodeExtensionStateNotSpecified NodeExtensionState = "NotSpecified"
	NodeExtensionStateSucceeded    NodeExtensionState = "Succeeded"
	NodeExtensionStateUpdating     NodeExtensionState = "Updating"
)

func PossibleValuesForNodeExtensionState() []string {
	return []string{
		string(NodeExtensionStateCanceled),
		string(NodeExtensionStateConnected),
		string(NodeExtensionStateCreating),
		string(NodeExtensionStateDeleted),
		string(NodeExtensionStateDeleting),
		string(NodeExtensionStateDisconnected),
		string(NodeExtensionStateError),
		string(NodeExtensionStateFailed),
		string(NodeExtensionStateMoving),
		string(NodeExtensionStateNotSpecified),
		string(NodeExtensionStateSucceeded),
		string(NodeExtensionStateUpdating),
	}
}

func parseNodeExtensionState(input string) (*NodeExtensionState, error) {
	vals := map[string]NodeExtensionState{
		"canceled":     NodeExtensionStateCanceled,
		"connected":    NodeExtensionStateConnected,
		"creating":     NodeExtensionStateCreating,
		"deleted":      NodeExtensionStateDeleted,
		"deleting":     NodeExtensionStateDeleting,
		"disconnected": NodeExtensionStateDisconnected,
		"error":        NodeExtensionStateError,
		"failed":       NodeExtensionStateFailed,
		"moving":       NodeExtensionStateMoving,
		"notspecified": NodeExtensionStateNotSpecified,
		"succeeded":    NodeExtensionStateSucceeded,
		"updating":     NodeExtensionStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodeExtensionState(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":     ProvisioningStateAccepted,
		"canceled":     ProvisioningStateCanceled,
		"failed":       ProvisioningStateFailed,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
