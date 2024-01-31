package migrateschemasqlserversqldbtasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateSchemaSqlServerSqlDbTaskProperties{}

type MigrateSchemaSqlServerSqlDbTaskProperties struct {
	Input  *MigrateSchemaSqlServerSqlDbTaskInput    `json:"input,omitempty"`
	Output *[]MigrateSchemaSqlServerSqlDbTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateSchemaSqlServerSqlDbTaskProperties{}

func (s MigrateSchemaSqlServerSqlDbTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSchemaSqlServerSqlDbTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSchemaSqlServerSqlDbTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSchemaSqlServerSqlDbTaskProperties: %+v", err)
	}
	decoded["taskType"] = "MigrateSchemaSqlServerSqlDb"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSchemaSqlServerSqlDbTaskProperties: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MigrateSchemaSqlServerSqlDbTaskProperties{}

func (s *MigrateSchemaSqlServerSqlDbTaskProperties) UnmarshalJSON(bytes []byte) error {
	type alias MigrateSchemaSqlServerSqlDbTaskProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MigrateSchemaSqlServerSqlDbTaskProperties: %+v", err)
	}

	s.ClientData = decoded.ClientData
	s.Commands = decoded.Commands
	s.Errors = decoded.Errors
	s.Input = decoded.Input
	s.State = decoded.State

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MigrateSchemaSqlServerSqlDbTaskProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["output"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Output into list []json.RawMessage: %+v", err)
		}

		output := make([]MigrateSchemaSqlServerSqlDbTaskOutput, 0)
		for i, val := range listTemp {
			impl, err := unmarshalMigrateSchemaSqlServerSqlDbTaskOutputImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Output' for 'MigrateSchemaSqlServerSqlDbTaskProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Output = &output
	}
	return nil
}
