package dataset

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSet = KustoTableDataSet{}

type KustoTableDataSet struct {
	Properties KustoTableDataSetProperties `json:"properties"`

	// Fields inherited from DataSet

	Id         *string                `json:"id,omitempty"`
	Kind       DataSetKind            `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s KustoTableDataSet) DataSet() BaseDataSetImpl {
	return BaseDataSetImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = KustoTableDataSet{}

func (s KustoTableDataSet) MarshalJSON() ([]byte, error) {
	type wrapper KustoTableDataSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KustoTableDataSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KustoTableDataSet: %+v", err)
	}

	decoded["kind"] = "KustoTable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KustoTableDataSet: %+v", err)
	}

	return encoded, nil
}
