package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PackageRequest struct {
	BaseEnvironmentSource    BaseEnvironmentSource `json:"baseEnvironmentSource"`
	EnvironmentVariables     *map[string]string    `json:"environmentVariables,omitempty"`
	InferencingServer        InferencingServer     `json:"inferencingServer"`
	Inputs                   *[]ModelPackageInput  `json:"inputs,omitempty"`
	ModelConfiguration       *ModelConfiguration   `json:"modelConfiguration,omitempty"`
	Tags                     *map[string]string    `json:"tags,omitempty"`
	TargetEnvironmentName    string                `json:"targetEnvironmentName"`
	TargetEnvironmentVersion *string               `json:"targetEnvironmentVersion,omitempty"`
}

var _ json.Unmarshaler = &PackageRequest{}

func (s *PackageRequest) UnmarshalJSON(bytes []byte) error {
	type alias PackageRequest
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into PackageRequest: %+v", err)
	}

	s.EnvironmentVariables = decoded.EnvironmentVariables
	s.Inputs = decoded.Inputs
	s.ModelConfiguration = decoded.ModelConfiguration
	s.Tags = decoded.Tags
	s.TargetEnvironmentName = decoded.TargetEnvironmentName
	s.TargetEnvironmentVersion = decoded.TargetEnvironmentVersion

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PackageRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["baseEnvironmentSource"]; ok {
		impl, err := unmarshalBaseEnvironmentSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'BaseEnvironmentSource' for 'PackageRequest': %+v", err)
		}
		s.BaseEnvironmentSource = impl
	}

	if v, ok := temp["inferencingServer"]; ok {
		impl, err := unmarshalInferencingServerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'InferencingServer' for 'PackageRequest': %+v", err)
		}
		s.InferencingServer = impl
	}
	return nil
}
