package datasetmapping

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = SynapseWorkspaceSqlPoolTableDataSetMapping{}

type SynapseWorkspaceSqlPoolTableDataSetMapping struct {
	Properties SynapseWorkspaceSqlPoolTableDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = SynapseWorkspaceSqlPoolTableDataSetMapping{}

func (s SynapseWorkspaceSqlPoolTableDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper SynapseWorkspaceSqlPoolTableDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SynapseWorkspaceSqlPoolTableDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SynapseWorkspaceSqlPoolTableDataSetMapping: %+v", err)
	}
	decoded["kind"] = "SynapseWorkspaceSqlPoolTable"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SynapseWorkspaceSqlPoolTableDataSetMapping: %+v", err)
	}

	return encoded, nil
}
