package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataImportSource = DatabaseSource{}

type DatabaseSource struct {
	Query                 *string              `json:"query,omitempty"`
	StoredProcedure       *string              `json:"storedProcedure,omitempty"`
	StoredProcedureParams *[]map[string]string `json:"storedProcedureParams,omitempty"`
	TableName             *string              `json:"tableName,omitempty"`

	// Fields inherited from DataImportSource
	Connection *string `json:"connection,omitempty"`
}

var _ json.Marshaler = DatabaseSource{}

func (s DatabaseSource) MarshalJSON() ([]byte, error) {
	type wrapper DatabaseSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DatabaseSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DatabaseSource: %+v", err)
	}
	decoded["sourceType"] = "database"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DatabaseSource: %+v", err)
	}

	return encoded, nil
}
