package post

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateSqlServerSqlMISyncTaskProperties{}

type MigrateSqlServerSqlMISyncTaskProperties struct {
	Input  *SqlServerSqlMISyncTaskInput           `json:"input,omitempty"`
	Output *[]MigrateSqlServerSqlMISyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties

	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
	TaskType   string               `json:"taskType"`
}

func (s MigrateSqlServerSqlMISyncTaskProperties) ProjectTaskProperties() BaseProjectTaskPropertiesImpl {
	return BaseProjectTaskPropertiesImpl{
		ClientData: s.ClientData,
		Commands:   s.Commands,
		Errors:     s.Errors,
		State:      s.State,
		TaskType:   s.TaskType,
	}
}

var _ json.Marshaler = MigrateSqlServerSqlMISyncTaskProperties{}

func (s MigrateSqlServerSqlMISyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSqlServerSqlMISyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSqlServerSqlMISyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSqlServerSqlMISyncTaskProperties: %+v", err)
	}

	decoded["taskType"] = "Migrate.SqlServer.AzureSqlDbMI.Sync.LRS"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSqlServerSqlMISyncTaskProperties: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &MigrateSqlServerSqlMISyncTaskProperties{}

func (s *MigrateSqlServerSqlMISyncTaskProperties) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Input      *SqlServerSqlMISyncTaskInput `json:"input,omitempty"`
		ClientData *map[string]string           `json:"clientData,omitempty"`
		Errors     *[]ODataError                `json:"errors,omitempty"`
		State      *TaskState                   `json:"state,omitempty"`
		TaskType   string                       `json:"taskType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Input = decoded.Input
	s.ClientData = decoded.ClientData
	s.Errors = decoded.Errors
	s.State = decoded.State
	s.TaskType = decoded.TaskType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MigrateSqlServerSqlMISyncTaskProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["commands"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Commands into list []json.RawMessage: %+v", err)
		}

		output := make([]CommandProperties, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalCommandPropertiesImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Commands' for 'MigrateSqlServerSqlMISyncTaskProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Commands = &output
	}

	if v, ok := temp["output"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Output into list []json.RawMessage: %+v", err)
		}

		output := make([]MigrateSqlServerSqlMISyncTaskOutput, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMigrateSqlServerSqlMISyncTaskOutputImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Output' for 'MigrateSqlServerSqlMISyncTaskProperties': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Output = &output
	}

	return nil
}
