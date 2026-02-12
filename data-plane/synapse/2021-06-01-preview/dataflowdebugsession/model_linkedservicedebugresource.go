package dataflowdebugsession

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedServiceDebugResource struct {
	Name       *string       `json:"name,omitempty"`
	Properties LinkedService `json:"properties"`
}

var _ json.Unmarshaler = &LinkedServiceDebugResource{}

func (s *LinkedServiceDebugResource) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Name *string `json:"name,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Name = decoded.Name

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LinkedServiceDebugResource into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := UnmarshalLinkedServiceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'LinkedServiceDebugResource': %+v", err)
		}
		s.Properties = impl
	}

	return nil
}
