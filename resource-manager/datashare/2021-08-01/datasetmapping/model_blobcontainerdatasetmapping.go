package datasetmapping

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataSetMapping = BlobContainerDataSetMapping{}

type BlobContainerDataSetMapping struct {
	Properties BlobContainerMappingProperties `json:"properties"`

	// Fields inherited from DataSetMapping

	Id         *string                `json:"id,omitempty"`
	Kind       DataSetMappingKind     `json:"kind"`
	Name       *string                `json:"name,omitempty"`
	SystemData *systemdata.SystemData `json:"systemData,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

func (s BlobContainerDataSetMapping) DataSetMapping() BaseDataSetMappingImpl {
	return BaseDataSetMappingImpl{
		Id:         s.Id,
		Kind:       s.Kind,
		Name:       s.Name,
		SystemData: s.SystemData,
		Type:       s.Type,
	}
}

var _ json.Marshaler = BlobContainerDataSetMapping{}

func (s BlobContainerDataSetMapping) MarshalJSON() ([]byte, error) {
	type wrapper BlobContainerDataSetMapping
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BlobContainerDataSetMapping: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BlobContainerDataSetMapping: %+v", err)
	}

	decoded["kind"] = "Container"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BlobContainerDataSetMapping: %+v", err)
	}

	return encoded, nil
}
