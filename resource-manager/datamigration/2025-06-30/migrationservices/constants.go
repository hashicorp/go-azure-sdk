package migrationservices

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthType string

const (
	AuthTypeAccountKey      AuthType = "AccountKey"
	AuthTypeManagedIdentity AuthType = "ManagedIdentity"
)

func PossibleValuesForAuthType() []string {
	return []string{
		string(AuthTypeAccountKey),
		string(AuthTypeManagedIdentity),
	}
}

func (s *AuthType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthType(input string) (*AuthType, error) {
	vals := map[string]AuthType{
		"accountkey":      AuthTypeAccountKey,
		"managedidentity": AuthTypeManagedIdentity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthType(input)
	return &out, nil
}

type MongoMigrationStatus string

const (
	MongoMigrationStatusCanceled   MongoMigrationStatus = "Canceled"
	MongoMigrationStatusCompleted  MongoMigrationStatus = "Completed"
	MongoMigrationStatusFailed     MongoMigrationStatus = "Failed"
	MongoMigrationStatusInProgress MongoMigrationStatus = "InProgress"
	MongoMigrationStatusNotStarted MongoMigrationStatus = "NotStarted"
)

func PossibleValuesForMongoMigrationStatus() []string {
	return []string{
		string(MongoMigrationStatusCanceled),
		string(MongoMigrationStatusCompleted),
		string(MongoMigrationStatusFailed),
		string(MongoMigrationStatusInProgress),
		string(MongoMigrationStatusNotStarted),
	}
}

func (s *MongoMigrationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMongoMigrationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMongoMigrationStatus(input string) (*MongoMigrationStatus, error) {
	vals := map[string]MongoMigrationStatus{
		"canceled":   MongoMigrationStatusCanceled,
		"completed":  MongoMigrationStatusCompleted,
		"failed":     MongoMigrationStatusFailed,
		"inprogress": MongoMigrationStatusInProgress,
		"notstarted": MongoMigrationStatusNotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MongoMigrationStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"canceled":     ProvisioningStateCanceled,
		"failed":       ProvisioningStateFailed,
		"provisioning": ProvisioningStateProvisioning,
		"succeeded":    ProvisioningStateSucceeded,
		"updating":     ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type ResourceType string

const (
	ResourceTypeMongoToCosmosDbMongo ResourceType = "MongoToCosmosDbMongo"
	ResourceTypeSqlDb                ResourceType = "SqlDb"
	ResourceTypeSqlMi                ResourceType = "SqlMi"
	ResourceTypeSqlVM                ResourceType = "SqlVm"
)

func PossibleValuesForResourceType() []string {
	return []string{
		string(ResourceTypeMongoToCosmosDbMongo),
		string(ResourceTypeSqlDb),
		string(ResourceTypeSqlMi),
		string(ResourceTypeSqlVM),
	}
}

func (s *ResourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResourceType(input string) (*ResourceType, error) {
	vals := map[string]ResourceType{
		"mongotocosmosdbmongo": ResourceTypeMongoToCosmosDbMongo,
		"sqldb":                ResourceTypeSqlDb,
		"sqlmi":                ResourceTypeSqlMi,
		"sqlvm":                ResourceTypeSqlVM,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceType(input)
	return &out, nil
}
