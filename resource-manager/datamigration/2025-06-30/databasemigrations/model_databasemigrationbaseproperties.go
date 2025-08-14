package databasemigrations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseMigrationBaseProperties interface {
	DatabaseMigrationBaseProperties() BaseDatabaseMigrationBasePropertiesImpl
}

var _ DatabaseMigrationBaseProperties = BaseDatabaseMigrationBasePropertiesImpl{}

type BaseDatabaseMigrationBasePropertiesImpl struct {
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

func (s BaseDatabaseMigrationBasePropertiesImpl) DatabaseMigrationBaseProperties() BaseDatabaseMigrationBasePropertiesImpl {
	return s
}

var _ DatabaseMigrationBaseProperties = RawDatabaseMigrationBasePropertiesImpl{}

// RawDatabaseMigrationBasePropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDatabaseMigrationBasePropertiesImpl struct {
	databaseMigrationBaseProperties BaseDatabaseMigrationBasePropertiesImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawDatabaseMigrationBasePropertiesImpl) DatabaseMigrationBaseProperties() BaseDatabaseMigrationBasePropertiesImpl {
	return s.databaseMigrationBaseProperties
}

func UnmarshalDatabaseMigrationBasePropertiesImplementation(input []byte) (DatabaseMigrationBaseProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DatabaseMigrationBaseProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "DatabaseMigrationProperties") {
		var out DatabaseMigrationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatabaseMigrationProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MongoToCosmosDbMongo") {
		var out DatabaseMigrationPropertiesCosmosDbMongo
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DatabaseMigrationPropertiesCosmosDbMongo: %+v", err)
		}
		return out, nil
	}

	var parent BaseDatabaseMigrationBasePropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDatabaseMigrationBasePropertiesImpl: %+v", err)
	}

	return RawDatabaseMigrationBasePropertiesImpl{
		databaseMigrationBaseProperties: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
