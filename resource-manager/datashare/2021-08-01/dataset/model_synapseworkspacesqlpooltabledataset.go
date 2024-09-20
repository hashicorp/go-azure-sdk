package dataset

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSet = SynapseWorkspaceSqlPoolTableDataSet{}

type SynapseWorkspaceSqlPoolTableDataSet struct {
	Properties SynapseWorkspaceSqlPoolTableDataSetProperties `json:"properties"`

	// Fields inherited from DataSet

	Id         *string                `json:"id,omitempty"`
	Kind       DataSetKind            `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s SynapseWorkspaceSqlPoolTableDataSet) DataSet() BaseDataSetImpl {
	return BaseDataSetImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = SynapseWorkspaceSqlPoolTableDataSet{}

func (s SynapseWorkspaceSqlPoolTableDataSet) MarshalJSON() ([]byte, error) {
	type wrapper SynapseWorkspaceSqlPoolTableDataSet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SynapseWorkspaceSqlPoolTableDataSet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SynapseWorkspaceSqlPoolTableDataSet: %+v", err)
	}

	decoded["kind"] = "SynapseWorkspaceSqlPoolTable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SynapseWorkspaceSqlPoolTableDataSet: %+v", err)
	}

	return encoded, nil
}
