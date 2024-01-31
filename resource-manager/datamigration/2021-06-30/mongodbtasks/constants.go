package mongodbtasks

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

type MongoDbErrorType string

const (
	MongoDbErrorTypeError           MongoDbErrorType = "Error"
	MongoDbErrorTypeValidationError MongoDbErrorType = "ValidationError"
	MongoDbErrorTypeWarning         MongoDbErrorType = "Warning"
)

func PossibleValuesForMongoDbErrorType() []string {
	return []string{
		string(MongoDbErrorTypeError),
		string(MongoDbErrorTypeValidationError),
		string(MongoDbErrorTypeWarning),
	}
}

func (s *MongoDbErrorType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMongoDbErrorType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMongoDbErrorType(input string) (*MongoDbErrorType, error) {
	vals := map[string]MongoDbErrorType{
		"error":           MongoDbErrorTypeError,
		"validationerror": MongoDbErrorTypeValidationError,
		"warning":         MongoDbErrorTypeWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MongoDbErrorType(input)
	return &out, nil
}

type MongoDbMigrationState string

const (
	MongoDbMigrationStateCanceled        MongoDbMigrationState = "Canceled"
	MongoDbMigrationStateComplete        MongoDbMigrationState = "Complete"
	MongoDbMigrationStateCopying         MongoDbMigrationState = "Copying"
	MongoDbMigrationStateFailed          MongoDbMigrationState = "Failed"
	MongoDbMigrationStateFinalizing      MongoDbMigrationState = "Finalizing"
	MongoDbMigrationStateInitialReplay   MongoDbMigrationState = "InitialReplay"
	MongoDbMigrationStateInitializing    MongoDbMigrationState = "Initializing"
	MongoDbMigrationStateNotStarted      MongoDbMigrationState = "NotStarted"
	MongoDbMigrationStateReplaying       MongoDbMigrationState = "Replaying"
	MongoDbMigrationStateRestarting      MongoDbMigrationState = "Restarting"
	MongoDbMigrationStateValidatingInput MongoDbMigrationState = "ValidatingInput"
)

func PossibleValuesForMongoDbMigrationState() []string {
	return []string{
		string(MongoDbMigrationStateCanceled),
		string(MongoDbMigrationStateComplete),
		string(MongoDbMigrationStateCopying),
		string(MongoDbMigrationStateFailed),
		string(MongoDbMigrationStateFinalizing),
		string(MongoDbMigrationStateInitialReplay),
		string(MongoDbMigrationStateInitializing),
		string(MongoDbMigrationStateNotStarted),
		string(MongoDbMigrationStateReplaying),
		string(MongoDbMigrationStateRestarting),
		string(MongoDbMigrationStateValidatingInput),
	}
}

func (s *MongoDbMigrationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMongoDbMigrationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMongoDbMigrationState(input string) (*MongoDbMigrationState, error) {
	vals := map[string]MongoDbMigrationState{
		"canceled":        MongoDbMigrationStateCanceled,
		"complete":        MongoDbMigrationStateComplete,
		"copying":         MongoDbMigrationStateCopying,
		"failed":          MongoDbMigrationStateFailed,
		"finalizing":      MongoDbMigrationStateFinalizing,
		"initialreplay":   MongoDbMigrationStateInitialReplay,
		"initializing":    MongoDbMigrationStateInitializing,
		"notstarted":      MongoDbMigrationStateNotStarted,
		"replaying":       MongoDbMigrationStateReplaying,
		"restarting":      MongoDbMigrationStateRestarting,
		"validatinginput": MongoDbMigrationStateValidatingInput,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MongoDbMigrationState(input)
	return &out, nil
}

type ResultType string

const (
	ResultTypeCollection ResultType = "Collection"
	ResultTypeDatabase   ResultType = "Database"
	ResultTypeMigration  ResultType = "Migration"
)

func PossibleValuesForResultType() []string {
	return []string{
		string(ResultTypeCollection),
		string(ResultTypeDatabase),
		string(ResultTypeMigration),
	}
}

func (s *ResultType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResultType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResultType(input string) (*ResultType, error) {
	vals := map[string]ResultType{
		"collection": ResultTypeCollection,
		"database":   ResultTypeDatabase,
		"migration":  ResultTypeMigration,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResultType(input)
	return &out, nil
}
