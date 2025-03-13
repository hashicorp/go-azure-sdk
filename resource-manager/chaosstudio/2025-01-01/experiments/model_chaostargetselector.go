package experiments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChaosTargetSelector interface {
	ChaosTargetSelector() BaseChaosTargetSelectorImpl
}

var _ ChaosTargetSelector = BaseChaosTargetSelectorImpl{}

type BaseChaosTargetSelectorImpl struct {
	Filter ChaosTargetFilter `json:"filter"`
	Id     string            `json:"id"`
	Type   SelectorType      `json:"type"`
}

func (s BaseChaosTargetSelectorImpl) ChaosTargetSelector() BaseChaosTargetSelectorImpl {
	return s
}

var _ ChaosTargetSelector = RawChaosTargetSelectorImpl{}

// RawChaosTargetSelectorImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawChaosTargetSelectorImpl struct {
	chaosTargetSelector BaseChaosTargetSelectorImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawChaosTargetSelectorImpl) ChaosTargetSelector() BaseChaosTargetSelectorImpl {
	return s.chaosTargetSelector
}

var _ json.Unmarshaler = &BaseChaosTargetSelectorImpl{}

func (s *BaseChaosTargetSelectorImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Id   string       `json:"id"`
		Type SelectorType `json:"type"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Id = decoded.Id
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseChaosTargetSelectorImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["filter"]; ok {
		impl, err := UnmarshalChaosTargetFilterImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Filter' for 'BaseChaosTargetSelectorImpl': %+v", err)
		}
		s.Filter = impl
	}

	return nil
}

func UnmarshalChaosTargetSelectorImplementation(input []byte) (ChaosTargetSelector, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ChaosTargetSelector into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "List") {
		var out ChaosTargetListSelector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChaosTargetListSelector: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Query") {
		var out ChaosTargetQuerySelector
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChaosTargetQuerySelector: %+v", err)
		}
		return out, nil
	}

	var parent BaseChaosTargetSelectorImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseChaosTargetSelectorImpl: %+v", err)
	}

	return RawChaosTargetSelectorImpl{
		chaosTargetSelector: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
