package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScheduleActionBase = ImportDataAction{}

type ImportDataAction struct {
	DataImportDefinition DataVersionBase `json:"dataImportDefinition"`

	// Fields inherited from ScheduleActionBase
}

var _ json.Marshaler = ImportDataAction{}

func (s ImportDataAction) MarshalJSON() ([]byte, error) {
	type wrapper ImportDataAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImportDataAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImportDataAction: %+v", err)
	}
	decoded["actionType"] = "ImportData"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImportDataAction: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ImportDataAction{}

func (s *ImportDataAction) UnmarshalJSON(bytes []byte) error {

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ImportDataAction into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["dataImportDefinition"]; ok {
		impl, err := unmarshalDataVersionBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DataImportDefinition' for 'ImportDataAction': %+v", err)
		}
		s.DataImportDefinition = impl
	}
	return nil
}
