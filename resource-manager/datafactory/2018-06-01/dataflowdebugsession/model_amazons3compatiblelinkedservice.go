package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LinkedService = AmazonS3CompatibleLinkedService{}

type AmazonS3CompatibleLinkedService struct {
	TypeProperties AmazonS3CompatibleLinkedServiceTypeProperties `json:"typeProperties"`

	// Fields inherited from LinkedService

	Annotations *[]interface{}                     `json:"annotations,omitempty"`
	ConnectVia  Reference                          `json:"connectVia"`
	Description *string                            `json:"description,omitempty"`
	Parameters  *map[string]ParameterSpecification `json:"parameters,omitempty"`
	Type        string                             `json:"type"`
	Version     *string                            `json:"version,omitempty"`
}

func (s AmazonS3CompatibleLinkedService) LinkedService() BaseLinkedServiceImpl {
	return BaseLinkedServiceImpl{
		Annotations: s.Annotations,
		ConnectVia:  s.ConnectVia,
		Description: s.Description,
		Parameters:  s.Parameters,
		Type:        s.Type,
		Version:     s.Version,
	}
}

var _ json.Marshaler = AmazonS3CompatibleLinkedService{}

func (s AmazonS3CompatibleLinkedService) MarshalJSON() ([]byte, error) {
	type wrapper AmazonS3CompatibleLinkedService
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AmazonS3CompatibleLinkedService: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AmazonS3CompatibleLinkedService: %+v", err)
	}

	decoded["type"] = "AmazonS3Compatible"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AmazonS3CompatibleLinkedService: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AmazonS3CompatibleLinkedService{}

func (s *AmazonS3CompatibleLinkedService) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		TypeProperties AmazonS3CompatibleLinkedServiceTypeProperties `json:"typeProperties"`
		Annotations    *[]interface{}                                `json:"annotations,omitempty"`
		Description    *string                                       `json:"description,omitempty"`
		Parameters     *map[string]ParameterSpecification            `json:"parameters,omitempty"`
		Type           string                                        `json:"type"`
		Version        *string                                       `json:"version,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.TypeProperties = decoded.TypeProperties
	s.Annotations = decoded.Annotations
	s.Description = decoded.Description
	s.Parameters = decoded.Parameters
	s.Type = decoded.Type
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AmazonS3CompatibleLinkedService into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["connectVia"]; ok {
		impl, err := UnmarshalReferenceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ConnectVia' for 'AmazonS3CompatibleLinkedService': %+v", err)
		}
		s.ConnectVia = impl
	}

	return nil
}
