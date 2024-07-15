package javacomponents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JavaComponentProperties interface {
}

// RawJavaComponentPropertiesImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJavaComponentPropertiesImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalJavaComponentPropertiesImplementation(input []byte) (JavaComponentProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JavaComponentProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["componentType"].(string)
	if !ok {
		return nil, nil
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

	out := RawJavaComponentPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
