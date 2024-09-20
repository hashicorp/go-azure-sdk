package datasetmapping

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = KustoDatabaseDataSetMapping{}

type KustoDatabaseDataSetMapping struct {
	Properties KustoDatabaseDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping

	Id   *string            `json:"id,omitempty"`
	Kind DataSetMappingKind `json:"kind"`
	Name *string            `json:"name,omitempty"`
	Type *string            `json:"type,omitempty"`
}

func (s KustoDatabaseDataSetMapping) DataSetMapping() BaseDataSetMappingImpl {
	return BaseDataSetMappingImpl{
		Id:   s.Id,
		Kind: s.Kind,
		Name: s.Name,
		Type: s.Type,
	}
}

var _ json.Marshaler = KustoDatabaseDataSetMapping{}

func (s KustoDatabaseDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper KustoDatabaseDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KustoDatabaseDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KustoDatabaseDataSetMapping: %+v", err)
	}

	decoded["kind"] = "KustoDatabase"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KustoDatabaseDataSetMapping: %+v", err)
	}

	return encoded, nil
}
