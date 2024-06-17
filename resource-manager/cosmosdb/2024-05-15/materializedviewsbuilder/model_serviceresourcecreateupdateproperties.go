package materializedviewsbuilder

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceResourceCreateUpdateProperties interface {
}

// RawServiceResourceCreateUpdatePropertiesImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceResourceCreateUpdatePropertiesImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalServiceResourceCreateUpdatePropertiesImplementation(input []byte) (ServiceResourceCreateUpdateProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceResourceCreateUpdateProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["serviceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "DataTransfer") {
		var out DataTransferServiceResourceCreateUpdateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataTransferServiceResourceCreateUpdateProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GraphAPICompute") {
		var out GraphAPIComputeServiceResourceCreateUpdateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GraphAPIComputeServiceResourceCreateUpdateProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MaterializedViewsBuilder") {
		var out MaterializedViewsBuilderServiceResourceCreateUpdateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MaterializedViewsBuilderServiceResourceCreateUpdateProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SqlDedicatedGateway") {
		var out SqlDedicatedGatewayServiceResourceCreateUpdateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SqlDedicatedGatewayServiceResourceCreateUpdateProperties: %+v", err)
		}
		return out, nil
	}

	out := RawServiceResourceCreateUpdatePropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
