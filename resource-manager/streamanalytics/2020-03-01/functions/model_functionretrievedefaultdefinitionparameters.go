package functions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FunctionRetrieveDefaultDefinitionParameters interface {
}

func unmarshalFunctionRetrieveDefaultDefinitionParametersImplementation(input []byte) (FunctionRetrieveDefaultDefinitionParameters, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FunctionRetrieveDefaultDefinitionParameters into map[string]interface: %+v", err)
	}

	value, ok := temp["bindingType"].(string)
	if !ok {
		return nil, nil
	}

	type RawFunctionRetrieveDefaultDefinitionParametersImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawFunctionRetrieveDefaultDefinitionParametersImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
