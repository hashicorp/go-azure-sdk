package batchdeployment

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchDeploymentConfiguration interface {
	BatchDeploymentConfiguration() BaseBatchDeploymentConfigurationImpl
}

var _ BatchDeploymentConfiguration = BaseBatchDeploymentConfigurationImpl{}

type BaseBatchDeploymentConfigurationImpl struct {
	DeploymentConfigurationType BatchDeploymentConfigurationType `json:"deploymentConfigurationType"`
}

func (s BaseBatchDeploymentConfigurationImpl) BatchDeploymentConfiguration() BaseBatchDeploymentConfigurationImpl {
	return s
}

var _ BatchDeploymentConfiguration = RawBatchDeploymentConfigurationImpl{}

// RawBatchDeploymentConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawBatchDeploymentConfigurationImpl struct {
	batchDeploymentConfiguration BaseBatchDeploymentConfigurationImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawBatchDeploymentConfigurationImpl) BatchDeploymentConfiguration() BaseBatchDeploymentConfigurationImpl {
	return s.batchDeploymentConfiguration
}

func (s RawBatchDeploymentConfigurationImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalBatchDeploymentConfigurationImplementation(input []byte) (BatchDeploymentConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BatchDeploymentConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["deploymentConfigurationType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "PipelineComponent") {
		var out BatchPipelineComponentDeploymentConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BatchPipelineComponentDeploymentConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseBatchDeploymentConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBatchDeploymentConfigurationImpl: %+v", err)
	}

	return RawBatchDeploymentConfigurationImpl{
		batchDeploymentConfiguration: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
