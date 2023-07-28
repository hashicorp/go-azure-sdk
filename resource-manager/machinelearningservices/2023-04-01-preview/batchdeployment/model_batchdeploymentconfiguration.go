package batchdeployment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchDeploymentConfiguration interface {
}

func unmarshalBatchDeploymentConfigurationImplementation(input []byte) (BatchDeploymentConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BatchDeploymentConfiguration into map[string]interface: %+v", err)
	}

	value, ok := temp["deploymentConfigurationType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "PipelineComponent") {
		var out BatchPipelineComponentDeploymentConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BatchPipelineComponentDeploymentConfiguration: %+v", err)
		}
		return out, nil
	}

	type RawBatchDeploymentConfigurationImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawBatchDeploymentConfigurationImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
