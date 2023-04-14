package services

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceResource struct {
	Id         *string                    `json:"id,omitempty"`
	Location   *string                    `json:"location,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	Properties *ServiceResourceProperties `json:"properties,omitempty"`
	SystemData *SystemData                `json:"systemData,omitempty"`
	Tags       *map[string]string         `json:"tags,omitempty"`
	Type       *string                    `json:"type,omitempty"`
}

var _ json.Unmarshaler = &ServiceResource{}

func (s *ServiceResource) UnmarshalJSON(bytes []byte) error {
	type alias ServiceResource
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ServiceResource: %+v", err)
	}

	s.Id = decoded.Id
	s.Location = decoded.Location
	s.Name = decoded.Name
	s.SystemData = decoded.SystemData
	s.Tags = decoded.Tags
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ServiceResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalServiceResourcePropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'ServiceResource': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
