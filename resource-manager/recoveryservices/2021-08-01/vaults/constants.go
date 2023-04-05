package vaults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InfrastructureEncryptionState string

const (
	InfrastructureEncryptionStateDisabled InfrastructureEncryptionState = "Disabled"
	InfrastructureEncryptionStateEnabled  InfrastructureEncryptionState = "Enabled"
)

func PossibleValuesForInfrastructureEncryptionState() []string {
	return []string{
		string(InfrastructureEncryptionStateDisabled),
		string(InfrastructureEncryptionStateEnabled),
	}
}

type PrivateEndpointConnectionStatus string

const (
	PrivateEndpointConnectionStatusApproved     PrivateEndpointConnectionStatus = "Approved"
	PrivateEndpointConnectionStatusDisconnected PrivateEndpointConnectionStatus = "Disconnected"
	PrivateEndpointConnectionStatusPending      PrivateEndpointConnectionStatus = "Pending"
	PrivateEndpointConnectionStatusRejected     PrivateEndpointConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointConnectionStatus() []string {
	return []string{
		string(PrivateEndpointConnectionStatusApproved),
		string(PrivateEndpointConnectionStatusDisconnected),
		string(PrivateEndpointConnectionStatusPending),
		string(PrivateEndpointConnectionStatusRejected),
	}
}

type ProvisioningState string

const (
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStatePending   ProvisioningState = "Pending"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStatePending),
		string(ProvisioningStateSucceeded),
	}
}

type ResourceMoveState string

const (
	ResourceMoveStateCommitFailed    ResourceMoveState = "CommitFailed"
	ResourceMoveStateCommitTimedout  ResourceMoveState = "CommitTimedout"
	ResourceMoveStateCriticalFailure ResourceMoveState = "CriticalFailure"
	ResourceMoveStateFailure         ResourceMoveState = "Failure"
	ResourceMoveStateInProgress      ResourceMoveState = "InProgress"
	ResourceMoveStateMoveSucceeded   ResourceMoveState = "MoveSucceeded"
	ResourceMoveStatePartialSuccess  ResourceMoveState = "PartialSuccess"
	ResourceMoveStatePrepareFailed   ResourceMoveState = "PrepareFailed"
	ResourceMoveStatePrepareTimedout ResourceMoveState = "PrepareTimedout"
	ResourceMoveStateUnknown         ResourceMoveState = "Unknown"
)

func PossibleValuesForResourceMoveState() []string {
	return []string{
		string(ResourceMoveStateCommitFailed),
		string(ResourceMoveStateCommitTimedout),
		string(ResourceMoveStateCriticalFailure),
		string(ResourceMoveStateFailure),
		string(ResourceMoveStateInProgress),
		string(ResourceMoveStateMoveSucceeded),
		string(ResourceMoveStatePartialSuccess),
		string(ResourceMoveStatePrepareFailed),
		string(ResourceMoveStatePrepareTimedout),
		string(ResourceMoveStateUnknown),
	}
}

type SkuName string

const (
	SkuNameRSZero   SkuName = "RS0"
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameRSZero),
		string(SkuNameStandard),
	}
}

type TriggerType string

const (
	TriggerTypeForcedUpgrade TriggerType = "ForcedUpgrade"
	TriggerTypeUserTriggered TriggerType = "UserTriggered"
)

func PossibleValuesForTriggerType() []string {
	return []string{
		string(TriggerTypeForcedUpgrade),
		string(TriggerTypeUserTriggered),
	}
}

type VaultPrivateEndpointState string

const (
	VaultPrivateEndpointStateEnabled VaultPrivateEndpointState = "Enabled"
	VaultPrivateEndpointStateNone    VaultPrivateEndpointState = "None"
)

func PossibleValuesForVaultPrivateEndpointState() []string {
	return []string{
		string(VaultPrivateEndpointStateEnabled),
		string(VaultPrivateEndpointStateNone),
	}
}

type VaultUpgradeState string

const (
	VaultUpgradeStateFailed     VaultUpgradeState = "Failed"
	VaultUpgradeStateInProgress VaultUpgradeState = "InProgress"
	VaultUpgradeStateUnknown    VaultUpgradeState = "Unknown"
	VaultUpgradeStateUpgraded   VaultUpgradeState = "Upgraded"
)

func PossibleValuesForVaultUpgradeState() []string {
	return []string{
		string(VaultUpgradeStateFailed),
		string(VaultUpgradeStateInProgress),
		string(VaultUpgradeStateUnknown),
		string(VaultUpgradeStateUpgraded),
	}
}
