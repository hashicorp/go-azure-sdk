package integrationruntime

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimeStatus interface {
}

func unmarshalIntegrationRuntimeStatusImplementation(input []byte) (IntegrationRuntimeStatus, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling IntegrationRuntimeStatus into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Managed") {
		var out ManagedIntegrationRuntimeStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManagedIntegrationRuntimeStatus: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SelfHosted") {
		var out SelfHostedIntegrationRuntimeStatus
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SelfHostedIntegrationRuntimeStatus: %+v", err)
		}
		return out, nil
	}

	type RawIntegrationRuntimeStatusImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawIntegrationRuntimeStatusImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
