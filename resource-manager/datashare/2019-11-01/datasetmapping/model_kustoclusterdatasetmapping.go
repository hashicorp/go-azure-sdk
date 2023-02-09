package datasetmapping

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = KustoClusterDataSetMapping{}

type KustoClusterDataSetMapping struct {
	Properties KustoClusterDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping
	Id   *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

var _ json.Marshaler = KustoClusterDataSetMapping{}

func (s KustoClusterDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper KustoClusterDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KustoClusterDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KustoClusterDataSetMapping: %+v", err)
	}
	decoded["kind"] = "KustoCluster"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KustoClusterDataSetMapping: %+v", err)
	}

	return encoded, nil
}
