package experiments

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ChaosTargetSelector = ChaosTargetListSelector{}

type ChaosTargetListSelector struct {
	Targets []TargetReference `json:"targets"`

	// Fields inherited from ChaosTargetSelector

	Filter ChaosTargetFilter `json:"filter"`
	Id     string            `json:"id"`
	Type   SelectorType      `json:"type"`
}

func (s ChaosTargetListSelector) ChaosTargetSelector() BaseChaosTargetSelectorImpl {
	return BaseChaosTargetSelectorImpl{
		Filter: s.Filter,
		Id:     s.Id,
		Type:   s.Type,
	}
}

var _ json.Marshaler = ChaosTargetListSelector{}

func (s ChaosTargetListSelector) MarshalJSON() ([]byte, error) {
	type wrapper ChaosTargetListSelector
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ChaosTargetListSelector: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ChaosTargetListSelector: %+v", err)
	}

	decoded["type"] = "List"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ChaosTargetListSelector: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ChaosTargetListSelector{}

func (s *ChaosTargetListSelector) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Targets []TargetReference `json:"targets"`
		Id      string            `json:"id"`
		Type    SelectorType      `json:"type"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Targets = decoded.Targets
	s.Id = decoded.Id
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ChaosTargetListSelector into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["filter"]; ok {
		impl, err := UnmarshalChaosTargetFilterImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Filter' for 'ChaosTargetListSelector': %+v", err)
		}
		s.Filter = impl
	}

	return nil
}
