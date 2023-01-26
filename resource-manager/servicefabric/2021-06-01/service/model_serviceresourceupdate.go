package service

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceResourceUpdate struct {
	Etag       *string                         `json:"etag,omitempty"`
	Id         *string                         `json:"id,omitempty"`
	Location   *string                         `json:"location,omitempty"`
	Name       *string                         `json:"name,omitempty"`
	Properties ServiceResourceUpdateProperties `json:"properties"`
	SystemData *SystemData                     `json:"systemData,omitempty"`
	Tags       *map[string]string              `json:"tags,omitempty"`
	Type       *string                         `json:"type,omitempty"`
}

var _ json.Unmarshaler = &ServiceResourceUpdate{}

func (s *ServiceResourceUpdate) UnmarshalJSON(bytes []byte) error {
	type alias ServiceResourceUpdate
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ServiceResourceUpdate: %+v", err)
	}

	s.Etag = decoded.Etag
	s.Id = decoded.Id
	s.Location = decoded.Location
	s.Name = decoded.Name
	s.SystemData = decoded.SystemData
	s.Tags = decoded.Tags
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ServiceResourceUpdate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalServiceResourceUpdatePropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'ServiceResourceUpdate': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
