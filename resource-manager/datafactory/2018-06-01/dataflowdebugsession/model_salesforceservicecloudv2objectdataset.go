package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Dataset = SalesforceServiceCloudV2ObjectDataset{}

type SalesforceServiceCloudV2ObjectDataset struct {
	TypeProperties *SalesforceServiceCloudV2ObjectDatasetTypeProperties `json:"typeProperties,omitempty"`

	// Fields inherited from Dataset

	Annotations       *[]interface{}                     `json:"annotations,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	Folder            *DatasetFolder                     `json:"folder,omitempty"`
	LinkedServiceName Reference                          `json:"linkedServiceName"`
	Parameters        *map[string]ParameterSpecification `json:"parameters,omitempty"`
	Schema            *interface{}                       `json:"schema,omitempty"`
	Structure         *interface{}                       `json:"structure,omitempty"`
	Type              string                             `json:"type"`
}

func (s SalesforceServiceCloudV2ObjectDataset) Dataset() BaseDatasetImpl {
	return BaseDatasetImpl{
		Annotations:       s.Annotations,
		Description:       s.Description,
		Folder:            s.Folder,
		LinkedServiceName: s.LinkedServiceName,
		Parameters:        s.Parameters,
		Schema:            s.Schema,
		Structure:         s.Structure,
		Type:              s.Type,
	}
}

var _ json.Marshaler = SalesforceServiceCloudV2ObjectDataset{}

func (s SalesforceServiceCloudV2ObjectDataset) MarshalJSON() ([]byte, error) {
	type wrapper SalesforceServiceCloudV2ObjectDataset
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SalesforceServiceCloudV2ObjectDataset: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SalesforceServiceCloudV2ObjectDataset: %+v", err)
	}

	decoded["type"] = "SalesforceServiceCloudV2Object"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SalesforceServiceCloudV2ObjectDataset: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SalesforceServiceCloudV2ObjectDataset{}

func (s *SalesforceServiceCloudV2ObjectDataset) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		TypeProperties *SalesforceServiceCloudV2ObjectDatasetTypeProperties `json:"typeProperties,omitempty"`
		Annotations    *[]interface{}                                       `json:"annotations,omitempty"`
		Description    *string                                              `json:"description,omitempty"`
		Folder         *DatasetFolder                                       `json:"folder,omitempty"`
		Parameters     *map[string]ParameterSpecification                   `json:"parameters,omitempty"`
		Schema         *interface{}                                         `json:"schema,omitempty"`
		Structure      *interface{}                                         `json:"structure,omitempty"`
		Type           string                                               `json:"type"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.TypeProperties = decoded.TypeProperties
	s.Annotations = decoded.Annotations
	s.Description = decoded.Description
	s.Folder = decoded.Folder
	s.Parameters = decoded.Parameters
	s.Schema = decoded.Schema
	s.Structure = decoded.Structure
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SalesforceServiceCloudV2ObjectDataset into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["linkedServiceName"]; ok {
		impl, err := UnmarshalReferenceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LinkedServiceName' for 'SalesforceServiceCloudV2ObjectDataset': %+v", err)
		}
		s.LinkedServiceName = impl
	}

	return nil
}
