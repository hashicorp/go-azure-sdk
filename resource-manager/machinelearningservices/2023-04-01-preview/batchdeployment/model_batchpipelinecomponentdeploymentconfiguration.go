package batchdeployment

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BatchDeploymentConfiguration = BatchPipelineComponentDeploymentConfiguration{}

type BatchPipelineComponentDeploymentConfiguration struct {
	ComponentId AssetReferenceBase `json:"componentId"`
	Description *string            `json:"description,omitempty"`
	Settings    *map[string]string `json:"settings,omitempty"`
	Tags        *map[string]string `json:"tags,omitempty"`

	// Fields inherited from BatchDeploymentConfiguration
}

var _ json.Marshaler = BatchPipelineComponentDeploymentConfiguration{}

func (s BatchPipelineComponentDeploymentConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper BatchPipelineComponentDeploymentConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BatchPipelineComponentDeploymentConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BatchPipelineComponentDeploymentConfiguration: %+v", err)
	}
	decoded["deploymentConfigurationType"] = "PipelineComponent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BatchPipelineComponentDeploymentConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &BatchPipelineComponentDeploymentConfiguration{}

func (s *BatchPipelineComponentDeploymentConfiguration) UnmarshalJSON(bytes []byte) error {
	type alias BatchPipelineComponentDeploymentConfiguration
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BatchPipelineComponentDeploymentConfiguration: %+v", err)
	}

	s.Description = decoded.Description
	s.Settings = decoded.Settings
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BatchPipelineComponentDeploymentConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["componentId"]; ok {
		impl, err := unmarshalAssetReferenceBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ComponentId' for 'BatchPipelineComponentDeploymentConfiguration': %+v", err)
		}
		s.ComponentId = impl
	}
	return nil
}
