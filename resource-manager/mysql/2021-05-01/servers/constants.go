package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMode string

const (
	CreateModeDefault            CreateMode = "Default"
	CreateModeGeoRestore         CreateMode = "GeoRestore"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeReplica            CreateMode = "Replica"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeDefault),
		string(CreateModeGeoRestore),
		string(CreateModePointInTimeRestore),
		string(CreateModeReplica),
	}
}

type DataEncryptionType string

const (
	DataEncryptionTypeAzureKeyVault DataEncryptionType = "AzureKeyVault"
	DataEncryptionTypeSystemManaged DataEncryptionType = "SystemManaged"
)

func PossibleValuesForDataEncryptionType() []string {
	return []string{
		string(DataEncryptionTypeAzureKeyVault),
		string(DataEncryptionTypeSystemManaged),
	}
}

type EnableStatusEnum string

const (
	EnableStatusEnumDisabled EnableStatusEnum = "Disabled"
	EnableStatusEnumEnabled  EnableStatusEnum = "Enabled"
)

func PossibleValuesForEnableStatusEnum() []string {
	return []string{
		string(EnableStatusEnumDisabled),
		string(EnableStatusEnumEnabled),
	}
}

type HighAvailabilityMode string

const (
	HighAvailabilityModeDisabled      HighAvailabilityMode = "Disabled"
	HighAvailabilityModeSameZone      HighAvailabilityMode = "SameZone"
	HighAvailabilityModeZoneRedundant HighAvailabilityMode = "ZoneRedundant"
)

func PossibleValuesForHighAvailabilityMode() []string {
	return []string{
		string(HighAvailabilityModeDisabled),
		string(HighAvailabilityModeSameZone),
		string(HighAvailabilityModeZoneRedundant),
	}
}

type HighAvailabilityState string

const (
	HighAvailabilityStateCreatingStandby HighAvailabilityState = "CreatingStandby"
	HighAvailabilityStateFailingOver     HighAvailabilityState = "FailingOver"
	HighAvailabilityStateHealthy         HighAvailabilityState = "Healthy"
	HighAvailabilityStateNotEnabled      HighAvailabilityState = "NotEnabled"
	HighAvailabilityStateRemovingStandby HighAvailabilityState = "RemovingStandby"
)

func PossibleValuesForHighAvailabilityState() []string {
	return []string{
		string(HighAvailabilityStateCreatingStandby),
		string(HighAvailabilityStateFailingOver),
		string(HighAvailabilityStateHealthy),
		string(HighAvailabilityStateNotEnabled),
		string(HighAvailabilityStateRemovingStandby),
	}
}

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

func PossibleValuesForManagedServiceIdentityType() []string {
	return []string{
		string(ManagedServiceIdentityTypeUserAssigned),
	}
}

type ReplicationRole string

const (
	ReplicationRoleNone    ReplicationRole = "None"
	ReplicationRoleReplica ReplicationRole = "Replica"
	ReplicationRoleSource  ReplicationRole = "Source"
)

func PossibleValuesForReplicationRole() []string {
	return []string{
		string(ReplicationRoleNone),
		string(ReplicationRoleReplica),
		string(ReplicationRoleSource),
	}
}

type ServerState string

const (
	ServerStateDisabled ServerState = "Disabled"
	ServerStateDropping ServerState = "Dropping"
	ServerStateReady    ServerState = "Ready"
	ServerStateStarting ServerState = "Starting"
	ServerStateStopped  ServerState = "Stopped"
	ServerStateStopping ServerState = "Stopping"
	ServerStateUpdating ServerState = "Updating"
)

func PossibleValuesForServerState() []string {
	return []string{
		string(ServerStateDisabled),
		string(ServerStateDropping),
		string(ServerStateReady),
		string(ServerStateStarting),
		string(ServerStateStopped),
		string(ServerStateStopping),
		string(ServerStateUpdating),
	}
}

type ServerVersion string

const (
	ServerVersionEightPointZeroPointTwoOne ServerVersion = "8.0.21"
	ServerVersionFivePointSeven            ServerVersion = "5.7"
)

func PossibleValuesForServerVersion() []string {
	return []string{
		string(ServerVersionEightPointZeroPointTwoOne),
		string(ServerVersionFivePointSeven),
	}
}

type SkuTier string

const (
	SkuTierBurstable       SkuTier = "Burstable"
	SkuTierGeneralPurpose  SkuTier = "GeneralPurpose"
	SkuTierMemoryOptimized SkuTier = "MemoryOptimized"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBurstable),
		string(SkuTierGeneralPurpose),
		string(SkuTierMemoryOptimized),
	}
}
