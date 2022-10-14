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

	type RawAddonPropertiesImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawAddonPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
