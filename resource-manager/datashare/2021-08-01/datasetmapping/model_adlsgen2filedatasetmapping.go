package datasetmapping

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = ADLSGen2FileDataSetMapping{}

type ADLSGen2FileDataSetMapping struct {
	Properties ADLSGen2FileDataSetMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ json.Marshaler = ADLSGen2FileDataSetMapping{}

func (s ADLSGen2FileDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper ADLSGen2FileDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ADLSGen2FileDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ADLSGen2FileDataSetMapping: %+v", err)
	}
	decoded["kind"] = "AdlsGen2File"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ADLSGen2FileDataSetMapping: %+v", err)
	}

	return encoded, nil
}
