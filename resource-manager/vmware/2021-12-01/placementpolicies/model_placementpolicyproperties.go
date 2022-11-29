package placementpolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementPolicyProperties interface {
}

func unmarshalPlacementPolicyPropertiesImplementation(input []byte) (PlacementPolicyProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PlacementPolicyProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "VmHost") {
		var out VMHostPlacementPolicyProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMHostPlacementPolicyProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VmVm") {
		var out VMVMPlacementPolicyProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMVMPlacementPolicyProperties: %+v", err)
		}
		return out, nil
	}

	type RawPlacementPolicyPropertiesImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawPlacementPolicyPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
