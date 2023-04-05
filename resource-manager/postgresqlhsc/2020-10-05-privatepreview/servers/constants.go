package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CitusVersion string

const (
	CitusVersionEightPointThree CitusVersion = "8.3"
	CitusVersionNinePointFive   CitusVersion = "9.5"
	CitusVersionNinePointFour   CitusVersion = "9.4"
	CitusVersionNinePointOne    CitusVersion = "9.1"
	CitusVersionNinePointThree  CitusVersion = "9.3"
	CitusVersionNinePointTwo    CitusVersion = "9.2"
	CitusVersionNinePointZero   CitusVersion = "9.0"
)

func PossibleValuesForCitusVersion() []string {
	return []string{
		string(CitusVersionEightPointThree),
		string(CitusVersionNinePointFive),
		string(CitusVersionNinePointFour),
		string(CitusVersionNinePointOne),
		string(CitusVersionNinePointThree),
		string(CitusVersionNinePointTwo),
		string(CitusVersionNinePointZero),
	}
}

type PostgreSQLVersion string

const (
	PostgreSQLVersionOneOne PostgreSQLVersion = "11"
	PostgreSQLVersionOneTwo PostgreSQLVersion = "12"
)

func PossibleValuesForPostgreSQLVersion() []string {
	return []string{
		string(PostgreSQLVersionOneOne),
		string(PostgreSQLVersionOneTwo),
	}
}

type ServerEdition string

const (
	ServerEditionGeneralPurpose  ServerEdition = "GeneralPurpose"
	ServerEditionMemoryOptimized ServerEdition = "MemoryOptimized"
)

func PossibleValuesForServerEdition() []string {
	return []string{
		string(ServerEditionGeneralPurpose),
		string(ServerEditionMemoryOptimized),
	}
}

type ServerHaState string

const (
	ServerHaStateCreatingStandby ServerHaState = "CreatingStandby"
	ServerHaStateFailingOver     ServerHaState = "FailingOver"
	ServerHaStateHealthy         ServerHaState = "Healthy"
	ServerHaStateNotEnabled      ServerHaState = "NotEnabled"
	ServerHaStateNotSync         ServerHaState = "NotSync"
	ServerHaStateRemovingStandby ServerHaState = "RemovingStandby"
	ServerHaStateReplicatingData ServerHaState = "ReplicatingData"
)

func PossibleValuesForServerHaState() []string {
	return []string{
		string(ServerHaStateCreatingStandby),
		string(ServerHaStateFailingOver),
		string(ServerHaStateHealthy),
		string(ServerHaStateNotEnabled),
		string(ServerHaStateNotSync),
		string(ServerHaStateRemovingStandby),
		string(ServerHaStateReplicatingData),
	}
}

type ServerRole string

const (
	ServerRoleCoordinator ServerRole = "Coordinator"
	ServerRoleWorker      ServerRole = "Worker"
)

func PossibleValuesForServerRole() []string {
	return []string{
		string(ServerRoleCoordinator),
		string(ServerRoleWorker),
	}
}

type ServerState string

const (
	ServerStateDisabled     ServerState = "Disabled"
	ServerStateDropping     ServerState = "Dropping"
	ServerStateProvisioning ServerState = "Provisioning"
	ServerStateReady        ServerState = "Ready"
	ServerStateStarting     ServerState = "Starting"
	ServerStateStopped      ServerState = "Stopped"
	ServerStateStopping     ServerState = "Stopping"
	ServerStateUpdating     ServerState = "Updating"
)

func PossibleValuesForServerState() []string {
	return []string{
		string(ServerStateDisabled),
		string(ServerStateDropping),
		string(ServerStateProvisioning),
		string(ServerStateReady),
		string(ServerStateStarting),
		string(ServerStateStopped),
		string(ServerStateStopping),
		string(ServerStateUpdating),
	}
}
