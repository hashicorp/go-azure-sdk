package connecttosourcemysqltasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandState string

const (
	CommandStateAccepted  CommandState = "Accepted"
	CommandStateFailed    CommandState = "Failed"
	CommandStateRunning   CommandState = "Running"
	CommandStateSucceeded CommandState = "Succeeded"
	CommandStateUnknown   CommandState = "Unknown"
)

func PossibleValuesForCommandState() []string {
	return []string{
		string(CommandStateAccepted),
		string(CommandStateFailed),
		string(CommandStateRunning),
		string(CommandStateSucceeded),
		string(CommandStateUnknown),
	}
}

func (s *CommandState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCommandState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCommandState(input string) (*CommandState, error) {
	vals := map[string]CommandState{
		"accepted":  CommandStateAccepted,
		"failed":    CommandStateFailed,
		"running":   CommandStateRunning,
		"succeeded": CommandStateSucceeded,
		"unknown":   CommandStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommandState(input)
	return &out, nil
}

type MySqlTargetPlatformType string

const (
	MySqlTargetPlatformTypeAzureDbForMySQL MySqlTargetPlatformType = "AzureDbForMySQL"
	MySqlTargetPlatformTypeSqlServer       MySqlTargetPlatformType = "SqlServer"
)

func PossibleValuesForMySqlTargetPlatformType() []string {
	return []string{
		string(MySqlTargetPlatformTypeAzureDbForMySQL),
		string(MySqlTargetPlatformTypeSqlServer),
	}
}

func (s *MySqlTargetPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMySqlTargetPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMySqlTargetPlatformType(input string) (*MySqlTargetPlatformType, error) {
	vals := map[string]MySqlTargetPlatformType{
		"azuredbformysql": MySqlTargetPlatformTypeAzureDbForMySQL,
		"sqlserver":       MySqlTargetPlatformTypeSqlServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MySqlTargetPlatformType(input)
	return &out, nil
}

type ServerLevelPermissionsGroup string

const (
	ServerLevelPermissionsGroupDefault                             ServerLevelPermissionsGroup = "Default"
	ServerLevelPermissionsGroupMigrationFromMySQLToAzureDBForMySQL ServerLevelPermissionsGroup = "MigrationFromMySQLToAzureDBForMySQL"
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB     ServerLevelPermissionsGroup = "MigrationFromSqlServerToAzureDB"
	ServerLevelPermissionsGroupMigrationFromSqlServerToAzureMI     ServerLevelPermissionsGroup = "MigrationFromSqlServerToAzureMI"
)

func PossibleValuesForServerLevelPermissionsGroup() []string {
	return []string{
		string(ServerLevelPermissionsGroupDefault),
		string(ServerLevelPermissionsGroupMigrationFromMySQLToAzureDBForMySQL),
		string(ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB),
		string(ServerLevelPermissionsGroupMigrationFromSqlServerToAzureMI),
	}
}

func (s *ServerLevelPermissionsGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServerLevelPermissionsGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServerLevelPermissionsGroup(input string) (*ServerLevelPermissionsGroup, error) {
	vals := map[string]ServerLevelPermissionsGroup{
		"default":                             ServerLevelPermissionsGroupDefault,
		"migrationfrommysqltoazuredbformysql": ServerLevelPermissionsGroupMigrationFromMySQLToAzureDBForMySQL,
		"migrationfromsqlservertoazuredb":     ServerLevelPermissionsGroupMigrationFromSqlServerToAzureDB,
		"migrationfromsqlservertoazuremi":     ServerLevelPermissionsGroupMigrationFromSqlServerToAzureMI,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServerLevelPermissionsGroup(input)
	return &out, nil
}

type TaskState string

const (
	TaskStateCanceled              TaskState = "Canceled"
	TaskStateFailed                TaskState = "Failed"
	TaskStateFailedInputValidation TaskState = "FailedInputValidation"
	TaskStateFaulted               TaskState = "Faulted"
	TaskStateQueued                TaskState = "Queued"
	TaskStateRunning               TaskState = "Running"
	TaskStateSucceeded             TaskState = "Succeeded"
	TaskStateUnknown               TaskState = "Unknown"
)

func PossibleValuesForTaskState() []string {
	return []string{
		string(TaskStateCanceled),
		string(TaskStateFailed),
		string(TaskStateFailedInputValidation),
		string(TaskStateFaulted),
		string(TaskStateQueued),
		string(TaskStateRunning),
		string(TaskStateSucceeded),
		string(TaskStateUnknown),
	}
}

func (s *TaskState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTaskState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTaskState(input string) (*TaskState, error) {
	vals := map[string]TaskState{
		"canceled":              TaskStateCanceled,
		"failed":                TaskStateFailed,
		"failedinputvalidation": TaskStateFailedInputValidation,
		"faulted":               TaskStateFaulted,
		"queued":                TaskStateQueued,
		"running":               TaskStateRunning,
		"succeeded":             TaskStateSucceeded,
		"unknown":               TaskStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TaskState(input)
	return &out, nil
}
