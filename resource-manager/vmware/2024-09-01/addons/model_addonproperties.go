package addons

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddonProperties interface {
	AddonProperties() BaseAddonPropertiesImpl
}

var _ AddonProperties = BaseAddonPropertiesImpl{}

type BaseAddonPropertiesImpl struct {
	AddonType         AddonType               `json:"addonType"`
	ProvisioningState *AddonProvisioningState `json:"provisioningState,omitempty"`
}

func (s BaseAddonPropertiesImpl) AddonProperties() BaseAddonPropertiesImpl {
	return s
}

var _ AddonProperties = RawAddonPropertiesImpl{}

// RawAddonPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAddonPropertiesImpl struct {
	addonProperties BaseAddonPropertiesImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawAddonPropertiesImpl) AddonProperties() BaseAddonPropertiesImpl {
	return s.addonProperties
}

func UnmarshalAddonPropertiesImplementation(input []byte) (AddonProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AddonProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["addonType"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseAddonPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAddonPropertiesImpl: %+v", err)
	}

	return RawAddonPropertiesImpl{
		addonProperties: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
