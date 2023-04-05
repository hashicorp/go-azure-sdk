package extensions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtensionAggregateState string

const (
	ExtensionAggregateStateAccepted                       ExtensionAggregateState = "Accepted"
	ExtensionAggregateStateCanceled                       ExtensionAggregateState = "Canceled"
	ExtensionAggregateStateConnected                      ExtensionAggregateState = "Connected"
	ExtensionAggregateStateCreating                       ExtensionAggregateState = "Creating"
	ExtensionAggregateStateDeleted                        ExtensionAggregateState = "Deleted"
	ExtensionAggregateStateDeleting                       ExtensionAggregateState = "Deleting"
	ExtensionAggregateStateDisconnected                   ExtensionAggregateState = "Disconnected"
	ExtensionAggregateStateError                          ExtensionAggregateState = "Error"
	ExtensionAggregateStateFailed                         ExtensionAggregateState = "Failed"
	ExtensionAggregateStateInProgress                     ExtensionAggregateState = "InProgress"
	ExtensionAggregateStateMoving                         ExtensionAggregateState = "Moving"
	ExtensionAggregateStateNotSpecified                   ExtensionAggregateState = "NotSpecified"
	ExtensionAggregateStatePartiallyConnected             ExtensionAggregateState = "PartiallyConnected"
	ExtensionAggregateStatePartiallySucceeded             ExtensionAggregateState = "PartiallySucceeded"
	ExtensionAggregateStateProvisioning                   ExtensionAggregateState = "Provisioning"
	ExtensionAggregateStateSucceeded                      ExtensionAggregateState = "Succeeded"
	ExtensionAggregateStateUpdating                       ExtensionAggregateState = "Updating"
	ExtensionAggregateStateUpgradeFailedRollbackSucceeded ExtensionAggregateState = "UpgradeFailedRollbackSucceeded"
)

func PossibleValuesForExtensionAggregateState() []string {
	return []string{
		string(ExtensionAggregateStateAccepted),
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
		string(ExtensionAggregateStateProvisioning),
		string(ExtensionAggregateStateSucceeded),
		string(ExtensionAggregateStateUpdating),
		string(ExtensionAggregateStateUpgradeFailedRollbackSucceeded),
	}
}

type ExtensionManagedBy string

const (
	ExtensionManagedByAzure ExtensionManagedBy = "Azure"
	ExtensionManagedByUser  ExtensionManagedBy = "User"
)

func PossibleValuesForExtensionManagedBy() []string {
	return []string{
		string(ExtensionManagedByAzure),
		string(ExtensionManagedByUser),
	}
}

type NodeExtensionState string

const (
	NodeExtensionStateAccepted           NodeExtensionState = "Accepted"
	NodeExtensionStateCanceled           NodeExtensionState = "Canceled"
	NodeExtensionStateConnected          NodeExtensionState = "Connected"
	NodeExtensionStateCreating           NodeExtensionState = "Creating"
	NodeExtensionStateDeleted            NodeExtensionState = "Deleted"
	NodeExtensionStateDeleting           NodeExtensionState = "Deleting"
	NodeExtensionStateDisconnected       NodeExtensionState = "Disconnected"
	NodeExtensionStateError              NodeExtensionState = "Error"
	NodeExtensionStateFailed             NodeExtensionState = "Failed"
	NodeExtensionStateInProgress         NodeExtensionState = "InProgress"
	NodeExtensionStateMoving             NodeExtensionState = "Moving"
	NodeExtensionStateNotSpecified       NodeExtensionState = "NotSpecified"
	NodeExtensionStatePartiallyConnected NodeExtensionState = "PartiallyConnected"
	NodeExtensionStatePartiallySucceeded NodeExtensionState = "PartiallySucceeded"
	NodeExtensionStateProvisioning       NodeExtensionState = "Provisioning"
	NodeExtensionStateSucceeded          NodeExtensionState = "Succeeded"
	NodeExtensionStateUpdating           NodeExtensionState = "Updating"
)

func PossibleValuesForNodeExtensionState() []string {
	return []string{
		string(NodeExtensionStateAccepted),
		string(NodeExtensionStateCanceled),
		string(NodeExtensionStateConnected),
		string(NodeExtensionStateCreating),
		string(NodeExtensionStateDeleted),
		string(NodeExtensionStateDeleting),
		string(NodeExtensionStateDisconnected),
		string(NodeExtensionStateError),
		string(NodeExtensionStateFailed),
		string(NodeExtensionStateInProgress),
		string(NodeExtensionStateMoving),
		string(NodeExtensionStateNotSpecified),
		string(NodeExtensionStatePartiallyConnected),
		string(NodeExtensionStatePartiallySucceeded),
		string(NodeExtensionStateProvisioning),
		string(NodeExtensionStateSucceeded),
		string(NodeExtensionStateUpdating),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted           ProvisioningState = "Accepted"
	ProvisioningStateCanceled           ProvisioningState = "Canceled"
	ProvisioningStateConnected          ProvisioningState = "Connected"
	ProvisioningStateCreating           ProvisioningState = "Creating"
	ProvisioningStateDeleted            ProvisioningState = "Deleted"
	ProvisioningStateDeleting           ProvisioningState = "Deleting"
	ProvisioningStateDisconnected       ProvisioningState = "Disconnected"
	ProvisioningStateError              ProvisioningState = "Error"
	ProvisioningStateFailed             ProvisioningState = "Failed"
	ProvisioningStateInProgress         ProvisioningState = "InProgress"
	ProvisioningStateMoving             ProvisioningState = "Moving"
	ProvisioningStateNotSpecified       ProvisioningState = "NotSpecified"
	ProvisioningStatePartiallyConnected ProvisioningState = "PartiallyConnected"
	ProvisioningStatePartiallySucceeded ProvisioningState = "PartiallySucceeded"
	ProvisioningStateProvisioning       ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded          ProvisioningState = "Succeeded"
	ProvisioningStateUpdating           ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateConnected),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateDisconnected),
		string(ProvisioningStateError),
		string(ProvisioningStateFailed),
		string(ProvisioningStateInProgress),
		string(ProvisioningStateMoving),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStatePartiallyConnected),
		string(ProvisioningStatePartiallySucceeded),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

func PossibleValuesForStatusLevelTypes() []string {
	return []string{
		string(StatusLevelTypesError),
		string(StatusLevelTypesInfo),
		string(StatusLevelTypesWarning),
	}
}
