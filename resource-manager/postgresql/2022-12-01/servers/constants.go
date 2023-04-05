package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveDirectoryAuthEnum string

const (
	ActiveDirectoryAuthEnumDisabled ActiveDirectoryAuthEnum = "Disabled"
	ActiveDirectoryAuthEnumEnabled  ActiveDirectoryAuthEnum = "Enabled"
)

func PossibleValuesForActiveDirectoryAuthEnum() []string {
	return []string{
		string(ActiveDirectoryAuthEnumDisabled),
		string(ActiveDirectoryAuthEnumEnabled),
	}
}

type ArmServerKeyType string

const (
	ArmServerKeyTypeAzureKeyVault ArmServerKeyType = "AzureKeyVault"
	ArmServerKeyTypeSystemManaged ArmServerKeyType = "SystemManaged"
)

func PossibleValuesForArmServerKeyType() []string {
	return []string{
		string(ArmServerKeyTypeAzureKeyVault),
		string(ArmServerKeyTypeSystemManaged),
	}
}

type CreateMode string

const (
	CreateModeCreate             CreateMode = "Create"
	CreateModeDefault            CreateMode = "Default"
	CreateModeGeoRestore         CreateMode = "GeoRestore"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeReplica            CreateMode = "Replica"
	CreateModeUpdate             CreateMode = "Update"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeCreate),
		string(CreateModeDefault),
		string(CreateModeGeoRestore),
		string(CreateModePointInTimeRestore),
		string(CreateModeReplica),
		string(CreateModeUpdate),
	}
}

type CreateModeForUpdate string

const (
	CreateModeForUpdateDefault CreateModeForUpdate = "Default"
	CreateModeForUpdateUpdate  CreateModeForUpdate = "Update"
)

func PossibleValuesForCreateModeForUpdate() []string {
	return []string{
		string(CreateModeForUpdateDefault),
		string(CreateModeForUpdateUpdate),
	}
}

type GeoRedundantBackupEnum string

const (
	GeoRedundantBackupEnumDisabled GeoRedundantBackupEnum = "Disabled"
	GeoRedundantBackupEnumEnabled  GeoRedundantBackupEnum = "Enabled"
)

func PossibleValuesForGeoRedundantBackupEnum() []string {
	return []string{
		string(GeoRedundantBackupEnumDisabled),
		string(GeoRedundantBackupEnumEnabled),
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

type IdentityType string

const (
	IdentityTypeNone           IdentityType = "None"
	IdentityTypeSystemAssigned IdentityType = "SystemAssigned"
	IdentityTypeUserAssigned   IdentityType = "UserAssigned"
)

func PossibleValuesForIdentityType() []string {
	return []string{
		string(IdentityTypeNone),
		string(IdentityTypeSystemAssigned),
		string(IdentityTypeUserAssigned),
	}
}

type PasswordAuthEnum string

const (
	PasswordAuthEnumDisabled PasswordAuthEnum = "Disabled"
	PasswordAuthEnumEnabled  PasswordAuthEnum = "Enabled"
)

func PossibleValuesForPasswordAuthEnum() []string {
	return []string{
		string(PasswordAuthEnumDisabled),
		string(PasswordAuthEnumEnabled),
	}
}

type ReplicationRole string

const (
	ReplicationRoleAsyncReplica    ReplicationRole = "AsyncReplica"
	ReplicationRoleGeoAsyncReplica ReplicationRole = "GeoAsyncReplica"
	ReplicationRoleNone            ReplicationRole = "None"
	ReplicationRolePrimary         ReplicationRole = "Primary"
)

func PossibleValuesForReplicationRole() []string {
	return []string{
		string(ReplicationRoleAsyncReplica),
		string(ReplicationRoleGeoAsyncReplica),
		string(ReplicationRoleNone),
		string(ReplicationRolePrimary),
	}
}

type ServerHAState string

const (
	ServerHAStateCreatingStandby ServerHAState = "CreatingStandby"
	ServerHAStateFailingOver     ServerHAState = "FailingOver"
	ServerHAStateHealthy         ServerHAState = "Healthy"
	ServerHAStateNotEnabled      ServerHAState = "NotEnabled"
	ServerHAStateRemovingStandby ServerHAState = "RemovingStandby"
	ServerHAStateReplicatingData ServerHAState = "ReplicatingData"
)

func PossibleValuesForServerHAState() []string {
	return []string{
		string(ServerHAStateCreatingStandby),
		string(ServerHAStateFailingOver),
		string(ServerHAStateHealthy),
		string(ServerHAStateNotEnabled),
		string(ServerHAStateRemovingStandby),
		string(ServerHAStateReplicatingData),
	}
}

type ServerPublicNetworkAccessState string

const (
	ServerPublicNetworkAccessStateDisabled ServerPublicNetworkAccessState = "Disabled"
	ServerPublicNetworkAccessStateEnabled  ServerPublicNetworkAccessState = "Enabled"
)

func PossibleValuesForServerPublicNetworkAccessState() []string {
	return []string{
		string(ServerPublicNetworkAccessStateDisabled),
		string(ServerPublicNetworkAccessStateEnabled),
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
	ServerVersionOneFour  ServerVersion = "14"
	ServerVersionOneOne   ServerVersion = "11"
	ServerVersionOneThree ServerVersion = "13"
	ServerVersionOneTwo   ServerVersion = "12"
)

func PossibleValuesForServerVersion() []string {
	return []string{
		string(ServerVersionOneFour),
		string(ServerVersionOneOne),
		string(ServerVersionOneThree),
		string(ServerVersionOneTwo),
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
