package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMode string

const (
	CreateModeCreate             CreateMode = "Create"
	CreateModeDefault            CreateMode = "Default"
	CreateModePointInTimeRestore CreateMode = "PointInTimeRestore"
	CreateModeUpdate             CreateMode = "Update"
)

func PossibleValuesForCreateMode() []string {
	return []string{
		string(CreateModeCreate),
		string(CreateModeDefault),
		string(CreateModePointInTimeRestore),
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
	HighAvailabilityModeZoneRedundant HighAvailabilityMode = "ZoneRedundant"
)

func PossibleValuesForHighAvailabilityMode() []string {
	return []string{
		string(HighAvailabilityModeDisabled),
		string(HighAvailabilityModeZoneRedundant),
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
