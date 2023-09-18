package modelversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InferencingServer interface {
}

// RawInferencingServerImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawInferencingServerImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalInferencingServerImplementation(input []byte) (InferencingServer, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling InferencingServer into map[string]interface: %+v", err)
	}

	value, ok := temp["serverType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureMLBatch") {
		var out AzureMLBatchInferencingServer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureMLBatchInferencingServer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureMLOnline") {
		var out AzureMLOnlineInferencingServer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureMLOnlineInferencingServer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out CustomInferencingServer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomInferencingServer: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Triton") {
		var out TritonInferencingServer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TritonInferencingServer: %+v", err)
		}
		return out, nil
	}

	out := RawInferencingServerImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
