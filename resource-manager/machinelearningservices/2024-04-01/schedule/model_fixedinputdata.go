package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringInputDataBase = FixedInputData{}

type FixedInputData struct {

	// Fields inherited from MonitoringInputDataBase

	Columns       *map[string]string      `json:"columns,omitempty"`
	DataContext   *string                 `json:"dataContext,omitempty"`
	InputDataType MonitoringInputDataType `json:"inputDataType"`
	JobInputType  JobInputType            `json:"jobInputType"`
	Uri           string                  `json:"uri"`
}

func (s FixedInputData) MonitoringInputDataBase() BaseMonitoringInputDataBaseImpl {
	return BaseMonitoringInputDataBaseImpl{
		Columns:       s.Columns,
		DataContext:   s.DataContext,
		InputDataType: s.InputDataType,
		JobInputType:  s.JobInputType,
		Uri:           s.Uri,
	}
}

var _ json.Marshaler = FixedInputData{}

func (s FixedInputData) MarshalJSON() ([]byte, error) {
	type wrapper FixedInputData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FixedInputData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FixedInputData: %+v", err)
	}

	decoded["inputDataType"] = "Fixed"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FixedInputData: %+v", err)
	}

	return encoded, nil
}
