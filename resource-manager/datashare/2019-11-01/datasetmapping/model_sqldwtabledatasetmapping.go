package datasetmapping

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = SqlDWTableDataSetMapping{}

type SqlDWTableDataSetMapping struct {
	Properties SqlDWTableDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

var _ json.Marshaler = SqlDWTableDataSetMapping{}

func (s SqlDWTableDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper SqlDWTableDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SqlDWTableDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SqlDWTableDataSetMapping: %+v", err)
	}
	decoded["kind"] = "SqlDWTable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SqlDWTableDataSetMapping: %+v", err)
	}

	return encoded, nil
}
