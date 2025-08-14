package sqlmigrationservices

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatabaseMigrationBaseProperties = DatabaseMigrationPropertiesSqlDb{}

type DatabaseMigrationPropertiesSqlDb struct {
	MigrationStatusDetails *SqlDbMigrationStatusDetails `json:"migrationStatusDetails,omitempty"`
	OfflineConfiguration   *SqlDbOfflineConfiguration   `json:"offlineConfiguration,omitempty"`
	TableList              *[]string                    `json:"tableList,omitempty"`
	TargetSqlConnection    *SqlConnectionInformation    `json:"targetSqlConnection,omitempty"`

	// Fields inherited from DatabaseMigrationBaseProperties

	EndedOn               *string            `json:"endedOn,omitempty"`
	Kind                  ResourceType       `json:"kind"`
	MigrationFailureError *ErrorInfo         `json:"migrationFailureError,omitempty"`
	MigrationOperationId  *string            `json:"migrationOperationId,omitempty"`
	MigrationService      *string            `json:"migrationService,omitempty"`
	MigrationStatus       *string            `json:"migrationStatus,omitempty"`
	ProvisioningError     *string            `json:"provisioningError,omitempty"`
	ProvisioningState     *ProvisioningState `json:"provisioningState,omitempty"`
	Scope                 *string            `json:"scope,omitempty"`
	StartedOn             *string            `json:"startedOn,omitempty"`
}

func (s DatabaseMigrationPropertiesSqlDb) DatabaseMigrationBaseProperties() BaseDatabaseMigrationBasePropertiesImpl {
	return BaseDatabaseMigrationBasePropertiesImpl{
		EndedOn:               s.EndedOn,
		Kind:                  s.Kind,
		MigrationFailureError: s.MigrationFailureError,
		MigrationOperationId:  s.MigrationOperationId,
		MigrationService:      s.MigrationService,
		MigrationStatus:       s.MigrationStatus,
		ProvisioningError:     s.ProvisioningError,
		ProvisioningState:     s.ProvisioningState,
		Scope:                 s.Scope,
		StartedOn:             s.StartedOn,
	}
}

var _ json.Marshaler = DatabaseMigrationPropertiesSqlDb{}

func (s DatabaseMigrationPropertiesSqlDb) MarshalJSON() ([]byte, error) {
	type wrapper DatabaseMigrationPropertiesSqlDb
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatabaseMigrationPropertiesSqlDb: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatabaseMigrationPropertiesSqlDb: %+v", err)
	}

	decoded["kind"] = "SqlDb"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatabaseMigrationPropertiesSqlDb: %+v", err)
	}

	return encoded, nil
}
