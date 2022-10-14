package placementpolicies

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementPolicy struct {
	Id         *string                   `json:"id,omitempty"`
	Name       *string                   `json:"name,omitempty"`
	Properties PlacementPolicyProperties `json:"properties"`
	Type       *string                   `json:"type,omitempty"`
}

var _ json.Unmarshaler = &PlacementPolicy{}

func (s *PlacementPolicy) UnmarshalJSON(bytes []byte) error {
	type alias PlacementPolicy
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into PlacementPolicy: %+v", err)
	}

	s.Id = decoded.Id
	s.Name = decoded.Name
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PlacementPolicy into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalPlacementPolicyPropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'PlacementPolicy': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
