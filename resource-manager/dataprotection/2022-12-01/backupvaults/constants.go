package backupvaults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertsState string

const (
	AlertsStateDisabled AlertsState = "Disabled"
	AlertsStateEnabled  AlertsState = "Enabled"
)

func PossibleValuesForAlertsState() []string {
	return []string{
		string(AlertsStateDisabled),
		string(AlertsStateEnabled),
	}
}

type CrossSubscriptionRestoreState string

const (
	CrossSubscriptionRestoreStateDisabled            CrossSubscriptionRestoreState = "Disabled"
	CrossSubscriptionRestoreStateEnabled             CrossSubscriptionRestoreState = "Enabled"
	CrossSubscriptionRestoreStatePermanentlyDisabled CrossSubscriptionRestoreState = "PermanentlyDisabled"
)

func PossibleValuesForCrossSubscriptionRestoreState() []string {
	return []string{
		string(CrossSubscriptionRestoreStateDisabled),
		string(CrossSubscriptionRestoreStateEnabled),
		string(CrossSubscriptionRestoreStatePermanentlyDisabled),
	}
}

type ImmutabilityState string

const (
	ImmutabilityStateDisabled ImmutabilityState = "Disabled"
	ImmutabilityStateLocked   ImmutabilityState = "Locked"
	ImmutabilityStateUnlocked ImmutabilityState = "Unlocked"
)

func PossibleValuesForImmutabilityState() []string {
	return []string{
		string(ImmutabilityStateDisabled),
		string(ImmutabilityStateLocked),
		string(ImmutabilityStateUnlocked),
	}
}

type ProvisioningState string

const (
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUnknown      ProvisioningState = "Unknown"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
		string(ProvisioningStateUpdating),
	}
}

type ResourceMoveState string

const (
	ResourceMoveStateCommitFailed    ResourceMoveState = "CommitFailed"
	ResourceMoveStateCommitTimedout  ResourceMoveState = "CommitTimedout"
	ResourceMoveStateCriticalFailure ResourceMoveState = "CriticalFailure"
	ResourceMoveStateFailed          ResourceMoveState = "Failed"
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
		string(ResourceMoveStateFailed),
		string(ResourceMoveStateInProgress),
		string(ResourceMoveStateMoveSucceeded),
		string(ResourceMoveStatePartialSuccess),
		string(ResourceMoveStatePrepareFailed),
		string(ResourceMoveStatePrepareTimedout),
		string(ResourceMoveStateUnknown),
	}
}

type SoftDeleteState string

const (
	SoftDeleteStateAlwaysOn SoftDeleteState = "AlwaysOn"
	SoftDeleteStateOff      SoftDeleteState = "Off"
	SoftDeleteStateOn       SoftDeleteState = "On"
)

func PossibleValuesForSoftDeleteState() []string {
	return []string{
		string(SoftDeleteStateAlwaysOn),
		string(SoftDeleteStateOff),
		string(SoftDeleteStateOn),
	}
}

type StorageSettingStoreTypes string

const (
	StorageSettingStoreTypesArchiveStore     StorageSettingStoreTypes = "ArchiveStore"
	StorageSettingStoreTypesOperationalStore StorageSettingStoreTypes = "OperationalStore"
	StorageSettingStoreTypesVaultStore       StorageSettingStoreTypes = "VaultStore"
)

func PossibleValuesForStorageSettingStoreTypes() []string {
	return []string{
		string(StorageSettingStoreTypesArchiveStore),
		string(StorageSettingStoreTypesOperationalStore),
		string(StorageSettingStoreTypesVaultStore),
	}
}

type StorageSettingTypes string

const (
	StorageSettingTypesGeoRedundant     StorageSettingTypes = "GeoRedundant"
	StorageSettingTypesLocallyRedundant StorageSettingTypes = "LocallyRedundant"
	StorageSettingTypesZoneRedundant    StorageSettingTypes = "ZoneRedundant"
)

func PossibleValuesForStorageSettingTypes() []string {
	return []string{
		string(StorageSettingTypesGeoRedundant),
		string(StorageSettingTypesLocallyRedundant),
		string(StorageSettingTypesZoneRedundant),
	}
}
