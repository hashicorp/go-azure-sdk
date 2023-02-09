package datasetmapping

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = ADLSGen2FileSystemDataSetMapping{}

type ADLSGen2FileSystemDataSetMapping struct {
	Properties ADLSGen2FileSystemDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = ADLSGen2FileSystemDataSetMapping{}

func (s ADLSGen2FileSystemDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper ADLSGen2FileSystemDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ADLSGen2FileSystemDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ADLSGen2FileSystemDataSetMapping: %+v", err)
	}
	decoded["kind"] = "AdlsGen2FileSystem"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ADLSGen2FileSystemDataSetMapping: %+v", err)
	}

	return encoded, nil
}
