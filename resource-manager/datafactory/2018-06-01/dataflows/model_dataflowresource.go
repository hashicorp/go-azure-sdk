package dataflows

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowResource struct {
	Etag       *string  `json:"etag,omitempty"`
	Id         *string  `json:"id,omitempty"`
	Name       *string  `json:"name,omitempty"`
	Properties DataFlow `json:"properties"`
	Type       *string  `json:"type,omitempty"`
}

var _ json.Unmarshaler = &DataFlowResource{}

func (s *DataFlowResource) UnmarshalJSON(bytes []byte) error {
	type alias DataFlowResource
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DataFlowResource: %+v", err)
	}

	s.Etag = decoded.Etag
	s.Id = decoded.Id
	s.Name = decoded.Name
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataFlowResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalDataFlowImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'DataFlowResource': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
