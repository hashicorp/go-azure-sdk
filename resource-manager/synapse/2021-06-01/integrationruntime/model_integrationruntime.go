package integrationruntime

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntime interface {
}

func unmarshalIntegrationRuntimeImplementation(input []byte) (IntegrationRuntime, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IntegrationRuntime into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Managed") {
		var out ManagedIntegrationRuntime
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIntegrationRuntime: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SelfHosted") {
		var out SelfHostedIntegrationRuntime
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SelfHostedIntegrationRuntime: %+v", err)
		}
		return out, nil
	}

	type RawIntegrationRuntimeImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawIntegrationRuntimeImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
