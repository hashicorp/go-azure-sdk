package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringInputDataBase = StaticInputData{}

type StaticInputData struct {
	PreprocessingComponentId *string `json:"preprocessingComponentId,omitempty"`
	WindowEnd                string  `json:"windowEnd"`
	WindowStart              string  `json:"windowStart"`

	// Fields inherited from MonitoringInputDataBase
	Columns      *map[string]string `json:"columns,omitempty"`
	DataContext  *string            `json:"dataContext,omitempty"`
	JobInputType JobInputType       `json:"jobInputType"`
	Uri          string             `json:"uri"`
}

var _ json.Marshaler = StaticInputData{}

func (s StaticInputData) MarshalJSON() ([]byte, error) {
	type wrapper StaticInputData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StaticInputData: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StaticInputData: %+v", err)
	}
	decoded["inputDataType"] = "Static"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StaticInputData: %+v", err)
	}

	return encoded, nil
}
