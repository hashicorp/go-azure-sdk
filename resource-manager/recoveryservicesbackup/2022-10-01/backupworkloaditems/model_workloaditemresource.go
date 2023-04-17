package backupworkloaditems

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadItemResource struct {
	ETag       *string            `json:"eTag,omitempty"`
	Id         *string            `json:"id,omitempty"`
	Location   *string            `json:"location,omitempty"`
	Name       *string            `json:"name,omitempty"`
	Properties *WorkloadItem      `json:"properties,omitempty"`
	Tags       *map[string]string `json:"tags,omitempty"`
	Type       *string            `json:"type,omitempty"`
}

var _ json.Unmarshaler = &WorkloadItemResource{}

func (s *WorkloadItemResource) UnmarshalJSON(bytes []byte) error {
	type alias WorkloadItemResource
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into WorkloadItemResource: %+v", err)
	}

	s.ETag = decoded.ETag
	s.Id = decoded.Id
	s.Location = decoded.Location
	s.Name = decoded.Name
	s.Tags = decoded.Tags
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling WorkloadItemResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalWorkloadItemImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'WorkloadItemResource': %+v", err)
		}
		s.Properties = &impl
	}
	return nil
}
