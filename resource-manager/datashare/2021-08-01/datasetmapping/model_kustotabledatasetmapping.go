package datasetmapping

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = KustoTableDataSetMapping{}

type KustoTableDataSetMapping struct {
	Properties KustoTableDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = KustoTableDataSetMapping{}

func (s KustoTableDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper KustoTableDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KustoTableDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KustoTableDataSetMapping: %+v", err)
	}
	decoded["kind"] = "KustoTable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KustoTableDataSetMapping: %+v", err)
	}

	return encoded, nil
}
