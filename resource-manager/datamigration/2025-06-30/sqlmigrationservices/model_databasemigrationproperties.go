package sqlmigrationservices

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DatabaseMigrationBaseProperties = DatabaseMigrationProperties{}

type DatabaseMigrationProperties struct {
	SourceDatabaseName      *string                   `json:"sourceDatabaseName,omitempty"`
	SourceServerName        *string                   `json:"sourceServerName,omitempty"`
	SourceSqlConnection     *SqlConnectionInformation `json:"sourceSqlConnection,omitempty"`
	TargetDatabaseCollation *string                   `json:"targetDatabaseCollation,omitempty"`

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

func (s DatabaseMigrationProperties) DatabaseMigrationBaseProperties() BaseDatabaseMigrationBasePropertiesImpl {
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

func (o *DatabaseMigrationProperties) GetEndedOnAsTime() (*time.Time, error) {
	if o.EndedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseMigrationProperties) SetEndedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndedOn = &formatted
}

func (o *DatabaseMigrationProperties) GetStartedOnAsTime() (*time.Time, error) {
	if o.StartedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseMigrationProperties) SetStartedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartedOn = &formatted
}

var _ json.Marshaler = DatabaseMigrationProperties{}

func (s DatabaseMigrationProperties) MarshalJSON() ([]byte, error) {
	type wrapper DatabaseMigrationProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatabaseMigrationProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatabaseMigrationProperties: %+v", err)
	}

	decoded["kind"] = "DatabaseMigrationProperties"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatabaseMigrationProperties: %+v", err)
	}

	return encoded, nil
}
