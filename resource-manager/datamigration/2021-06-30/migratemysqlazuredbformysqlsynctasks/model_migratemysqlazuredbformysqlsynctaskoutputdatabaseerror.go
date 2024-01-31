package migratemysqlazuredbformysqlsynctasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigrateMySqlAzureDbForMySqlSyncTaskOutput = MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseError{}

type MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseError struct {
	ErrorMessage *string                            `json:"errorMessage,omitempty"`
	Events       *[]SyncMigrationDatabaseErrorEvent `json:"events,omitempty"`

	// Fields inherited from MigrateMySqlAzureDbForMySqlSyncTaskOutput
	Id *string `json:"id,omitempty"`
}

var _ json.Marshaler = MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseError{}

func (s MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseError) MarshalJSON() ([]byte, error) {
	type wrapper MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseError: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseError: %+v", err)
	}
	decoded["resultType"] = "DatabaseLevelErrorOutput"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateMySqlAzureDbForMySqlSyncTaskOutputDatabaseError: %+v", err)
	}

	return encoded, nil
}
