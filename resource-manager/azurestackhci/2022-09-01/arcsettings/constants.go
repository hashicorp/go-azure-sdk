package arcsettings

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArcSettingAggregateState string

const (
	ArcSettingAggregateStateCanceled           ArcSettingAggregateState = "Canceled"
	ArcSettingAggregateStateConnected          ArcSettingAggregateState = "Connected"
	ArcSettingAggregateStateCreating           ArcSettingAggregateState = "Creating"
	ArcSettingAggregateStateDeleted            ArcSettingAggregateState = "Deleted"
	ArcSettingAggregateStateDeleting           ArcSettingAggregateState = "Deleting"
	ArcSettingAggregateStateDisconnected       ArcSettingAggregateState = "Disconnected"
	ArcSettingAggregateStateError              ArcSettingAggregateState = "Error"
	ArcSettingAggregateStateFailed             ArcSettingAggregateState = "Failed"
	ArcSettingAggregateStateInProgress         ArcSettingAggregateState = "InProgress"
	ArcSettingAggregateStateMoving             ArcSettingAggregateState = "Moving"
	ArcSettingAggregateStateNotSpecified       ArcSettingAggregateState = "NotSpecified"
	ArcSettingAggregateStatePartiallyConnected ArcSettingAggregateState = "PartiallyConnected"
	ArcSettingAggregateStatePartiallySucceeded ArcSettingAggregateState = "PartiallySucceeded"
	ArcSettingAggregateStateSucceeded          ArcSettingAggregateState = "Succeeded"
	ArcSettingAggregateStateUpdating           ArcSettingAggregateState = "Updating"
)

func PossibleValuesForArcSettingAggregateState() []string {
	return []string{
		string(ArcSettingAggregateStateCanceled),
		string(ArcSettingAggregateStateConnected),
		string(ArcSettingAggregateStateCreating),
		string(ArcSettingAggregateStateDeleted),
		string(ArcSettingAggregateStateDeleting),
		string(ArcSettingAggregateStateDisconnected),
		string(ArcSettingAggregateStateError),
		string(ArcSettingAggregateStateFailed),
		string(ArcSettingAggregateStateInProgress),
		string(ArcSettingAggregateStateMoving),
		string(ArcSettingAggregateStateNotSpecified),
		string(ArcSettingAggregateStatePartiallyConnected),
		string(ArcSettingAggregateStatePartiallySucceeded),
		string(ArcSettingAggregateStateSucceeded),
		string(ArcSettingAggregateStateUpdating),
	}
}

func parseArcSettingAggregateState(input string) (*ArcSettingAggregateState, error) {
	vals := map[string]ArcSettingAggregateState{
		"canceled":           ArcSettingAggregateStateCanceled,
		"connected":          ArcSettingAggregateStateConnected,
		"creating":           ArcSettingAggregateStateCreating,
		"deleted":            ArcSettingAggregateStateDeleted,
		"deleting":           ArcSettingAggregateStateDeleting,
		"disconnected":       ArcSettingAggregateStateDisconnected,
		"error":              ArcSettingAggregateStateError,
		"failed":             ArcSettingAggregateStateFailed,
		"inprogress":         ArcSettingAggregateStateInProgress,
		"moving":             ArcSettingAggregateStateMoving,
		"notspecified":       ArcSettingAggregateStateNotSpecified,
		"partiallyconnected": ArcSettingAggregateStatePartiallyConnected,
		"partiallysucceeded": ArcSettingAggregateStatePartiallySucceeded,
		"succeeded":          ArcSettingAggregateStateSucceeded,
		"updating":           ArcSettingAggregateStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ArcSettingAggregateState(input)
	return &out, nil
}

type NodeArcState string

const (
	NodeArcStateCanceled     NodeArcState = "Canceled"
	NodeArcStateConnected    NodeArcState = "Connected"
	NodeArcStateCreating     NodeArcState = "Creating"
	NodeArcStateDeleted      NodeArcState = "Deleted"
	NodeArcStateDeleting     NodeArcState = "Deleting"
	NodeArcStateDisconnected NodeArcState = "Disconnected"
	NodeArcStateError        NodeArcState = "Error"
	NodeArcStateFailed       NodeArcState = "Failed"
	NodeArcStateMoving       NodeArcState = "Moving"
	NodeArcStateNotSpecified NodeArcState = "NotSpecified"
	NodeArcStateSucceeded    NodeArcState = "Succeeded"
	NodeArcStateUpdating     NodeArcState = "Updating"
)

func PossibleValuesForNodeArcState() []string {
	return []string{
		string(NodeArcStateCanceled),
		string(NodeArcStateConnected),
		string(NodeArcStateCreating),
		string(NodeArcStateDeleted),
		string(NodeArcStateDeleting),
		string(NodeArcStateDisconnected),
		string(NodeArcStateError),
		string(NodeArcStateFailed),
		string(NodeArcStateMoving),
		string(NodeArcStateNotSpecified),
		string(NodeArcStateSucceeded),
		string(NodeArcStateUpdating),
	}
}

func parseNodeArcState(input string) (*NodeArcState, error) {
	vals := map[string]NodeArcState{
		"canceled":     NodeArcStateCanceled,
		"connected":    NodeArcStateConnected,
		"creating":     NodeArcStateCreating,
		"deleted":      NodeArcStateDeleted,
		"deleting":     NodeArcStateDeleting,
		"disconnected": NodeArcStateDisconnected,
		"error":        NodeArcStateError,
		"failed":       NodeArcStateFailed,
		"moving":       NodeArcStateMoving,
		"notspecified": NodeArcStateNotSpecified,
		"succeeded":    NodeArcStateSucceeded,
		"updating":     NodeArcStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodeArcState(input)
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
