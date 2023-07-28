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

	type RawInferencingServerImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawInferencingServerImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
