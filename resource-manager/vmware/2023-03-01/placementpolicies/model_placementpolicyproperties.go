package placementpolicies

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlacementPolicyProperties interface {
	PlacementPolicyProperties() BasePlacementPolicyPropertiesImpl
}

var _ PlacementPolicyProperties = BasePlacementPolicyPropertiesImpl{}

type BasePlacementPolicyPropertiesImpl struct {
	DisplayName       *string                           `json:"displayName,omitempty"`
	ProvisioningState *PlacementPolicyProvisioningState `json:"provisioningState,omitempty"`
	State             *PlacementPolicyState             `json:"state,omitempty"`
	Type              PlacementPolicyType               `json:"type"`
}

func (s BasePlacementPolicyPropertiesImpl) PlacementPolicyProperties() BasePlacementPolicyPropertiesImpl {
	return s
}

var _ PlacementPolicyProperties = RawPlacementPolicyPropertiesImpl{}

// RawPlacementPolicyPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPlacementPolicyPropertiesImpl struct {
	placementPolicyProperties BasePlacementPolicyPropertiesImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawPlacementPolicyPropertiesImpl) PlacementPolicyProperties() BasePlacementPolicyPropertiesImpl {
	return s.placementPolicyProperties
}

func UnmarshalPlacementPolicyPropertiesImplementation(input []byte) (PlacementPolicyProperties, error) {
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

	var parent BasePlacementPolicyPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePlacementPolicyPropertiesImpl: %+v", err)
	}

	return RawPlacementPolicyPropertiesImpl{
		placementPolicyProperties: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
