package addons

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddonProperties interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAddonPropertiesImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalAddonPropertiesImplementation(input []byte) (AddonProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AddonProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["addonType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Arc") {
		var out AddonArcProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddonArcProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "HCX") {
		var out AddonHcxProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddonHcxProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SRM") {
		var out AddonSrmProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddonSrmProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VR") {
		var out AddonVrProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddonVrProperties: %+v", err)
		}
		return out, nil
	}

	out := RawAddonPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
