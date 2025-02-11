package experiments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChaosTargetFilter interface {
	ChaosTargetFilter() BaseChaosTargetFilterImpl
}

var _ ChaosTargetFilter = BaseChaosTargetFilterImpl{}

type BaseChaosTargetFilterImpl struct {
	Type FilterType `json:"type"`
}

func (s BaseChaosTargetFilterImpl) ChaosTargetFilter() BaseChaosTargetFilterImpl {
	return s
}

var _ ChaosTargetFilter = RawChaosTargetFilterImpl{}

// RawChaosTargetFilterImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawChaosTargetFilterImpl struct {
	chaosTargetFilter BaseChaosTargetFilterImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawChaosTargetFilterImpl) ChaosTargetFilter() BaseChaosTargetFilterImpl {
	return s.chaosTargetFilter
}

func UnmarshalChaosTargetFilterImplementation(input []byte) (ChaosTargetFilter, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ChaosTargetFilter into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Simple") {
		var out ChaosTargetSimpleFilter
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ChaosTargetSimpleFilter: %+v", err)
		}
		return out, nil
	}

	var parent BaseChaosTargetFilterImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseChaosTargetFilterImpl: %+v", err)
	}

	return RawChaosTargetFilterImpl{
		chaosTargetFilter: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
