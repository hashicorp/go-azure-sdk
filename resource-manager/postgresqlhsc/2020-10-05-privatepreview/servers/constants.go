package servers

import (
	"encoding/json"
	"fmt"
	"strings"
)

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

func (s *CitusVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCitusVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCitusVersion(input string) (*CitusVersion, error) {
	vals := map[string]CitusVersion{
		"8.3": CitusVersionEightPointThree,
		"9.5": CitusVersionNinePointFive,
		"9.4": CitusVersionNinePointFour,
		"9.1": CitusVersionNinePointOne,
		"9.3": CitusVersionNinePointThree,
		"9.2": CitusVersionNinePointTwo,
		"9.0": CitusVersionNinePointZero,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CitusVersion(input)
	return &out, nil
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

func (s *PostgreSQLVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePostgreSQLVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePostgreSQLVersion(input string) (*PostgreSQLVersion, error) {
	vals := map[string]PostgreSQLVersion{
		"11": PostgreSQLVersionOneOne,
		"12": PostgreSQLVersionOneTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PostgreSQLVersion(input)
	return &out, nil
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

func (s *ServerEdition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerEdition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerEdition(input string) (*ServerEdition, error) {
	vals := map[string]ServerEdition{
		"generalpurpose":  ServerEditionGeneralPurpose,
		"memoryoptimized": ServerEditionMemoryOptimized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerEdition(input)
	return &out, nil
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

func (s *ServerHaState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerHaState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerHaState(input string) (*ServerHaState, error) {
	vals := map[string]ServerHaState{
		"creatingstandby": ServerHaStateCreatingStandby,
		"failingover":     ServerHaStateFailingOver,
		"healthy":         ServerHaStateHealthy,
		"notenabled":      ServerHaStateNotEnabled,
		"notsync":         ServerHaStateNotSync,
		"removingstandby": ServerHaStateRemovingStandby,
		"replicatingdata": ServerHaStateReplicatingData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerHaState(input)
	return &out, nil
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

func (s *ServerRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerRole(input string) (*ServerRole, error) {
	vals := map[string]ServerRole{
		"coordinator": ServerRoleCoordinator,
		"worker":      ServerRoleWorker,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerRole(input)
	return &out, nil
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

func (s *ServerState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerState(input string) (*ServerState, error) {
	vals := map[string]ServerState{
		"disabled":     ServerStateDisabled,
		"dropping":     ServerStateDropping,
		"provisioning": ServerStateProvisioning,
		"ready":        ServerStateReady,
		"starting":     ServerStateStarting,
		"stopped":      ServerStateStopped,
		"stopping":     ServerStateStopping,
		"updating":     ServerStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerState(input)
	return &out, nil
}
