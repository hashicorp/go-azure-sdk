package extendedueinformation

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendedUeInfoProperties interface {
	ExtendedUeInfoProperties() BaseExtendedUeInfoPropertiesImpl
}

var _ ExtendedUeInfoProperties = BaseExtendedUeInfoPropertiesImpl{}

type BaseExtendedUeInfoPropertiesImpl struct {
	LastReadAt *string `json:"lastReadAt,omitempty"`
	RatType    RatType `json:"ratType"`
}

func (s BaseExtendedUeInfoPropertiesImpl) ExtendedUeInfoProperties() BaseExtendedUeInfoPropertiesImpl {
	return s
}

var _ ExtendedUeInfoProperties = RawExtendedUeInfoPropertiesImpl{}

// RawExtendedUeInfoPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawExtendedUeInfoPropertiesImpl struct {
	extendedUeInfoProperties BaseExtendedUeInfoPropertiesImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawExtendedUeInfoPropertiesImpl) ExtendedUeInfoProperties() BaseExtendedUeInfoPropertiesImpl {
	return s.extendedUeInfoProperties
}

func UnmarshalExtendedUeInfoPropertiesImplementation(input []byte) (ExtendedUeInfoProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ExtendedUeInfoProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["ratType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "4G") {
		var out UeInfo4G
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UeInfo4G: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "5G") {
		var out UeInfo5G
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UeInfo5G: %+v", err)
		}
		return out, nil
	}

	var parent BaseExtendedUeInfoPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseExtendedUeInfoPropertiesImpl: %+v", err)
	}

	return RawExtendedUeInfoPropertiesImpl{
		extendedUeInfoProperties: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
