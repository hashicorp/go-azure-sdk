package javacomponents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JavaComponentProperties interface {
	JavaComponentProperties() BaseJavaComponentPropertiesImpl
}

var _ JavaComponentProperties = BaseJavaComponentPropertiesImpl{}

type BaseJavaComponentPropertiesImpl struct {
	ComponentType     JavaComponentType                     `json:"componentType"`
	Configurations    *[]JavaComponentConfigurationProperty `json:"configurations,omitempty"`
	ProvisioningState *JavaComponentProvisioningState       `json:"provisioningState,omitempty"`
	ServiceBinds      *[]JavaComponentServiceBind           `json:"serviceBinds,omitempty"`
}

func (s BaseJavaComponentPropertiesImpl) JavaComponentProperties() BaseJavaComponentPropertiesImpl {
	return s
}

var _ JavaComponentProperties = RawJavaComponentPropertiesImpl{}

// RawJavaComponentPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJavaComponentPropertiesImpl struct {
	javaComponentProperties BaseJavaComponentPropertiesImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawJavaComponentPropertiesImpl) JavaComponentProperties() BaseJavaComponentPropertiesImpl {
	return s.javaComponentProperties
}

func UnmarshalJavaComponentPropertiesImplementation(input []byte) (JavaComponentProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JavaComponentProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["componentType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Nacos") {
		var out NacosComponent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NacosComponent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SpringBootAdmin") {
		var out SpringBootAdminComponent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SpringBootAdminComponent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SpringCloudConfig") {
		var out SpringCloudConfigComponent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SpringCloudConfigComponent: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SpringCloudEureka") {
		var out SpringCloudEurekaComponent
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SpringCloudEurekaComponent: %+v", err)
		}
		return out, nil
	}

	var parent BaseJavaComponentPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJavaComponentPropertiesImpl: %+v", err)
	}

	return RawJavaComponentPropertiesImpl{
		javaComponentProperties: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
