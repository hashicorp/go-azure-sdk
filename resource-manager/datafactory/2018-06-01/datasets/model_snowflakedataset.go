package datasets

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Dataset = SnowflakeDataset{}

type SnowflakeDataset struct {
	TypeProperties SnowflakeDatasetTypeProperties `json:"typeProperties"`

	// Fields inherited from Dataset
	Annotations       *[]interface{}                     `json:"annotations,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Folder            *DatasetFolder                     `json:"folder,omitempty"`
	LinkedServiceName LinkedServiceReference             `json:"linkedServiceName"`
	Parameters        *map[string]ParameterSpecification `json:"parameters,omitempty"`
	Schema            *interface{}                       `json:"schema,omitempty"`
	Structure         *interface{}                       `json:"structure,omitempty"`
}

var _ json.Marshaler = SnowflakeDataset{}

func (s SnowflakeDataset) MarshalJSON() ([]byte, error) {
	type wrapper SnowflakeDataset
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SnowflakeDataset: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SnowflakeDataset: %+v", err)
	}
	decoded["type"] = "SnowflakeTable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SnowflakeDataset: %+v", err)
	}

	return encoded, nil
}
