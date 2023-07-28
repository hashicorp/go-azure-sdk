package modelversion

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PackageResponse struct {
	BaseEnvironmentSource    BaseEnvironmentSource `json:"baseEnvironmentSource"`
	BuildId                  *string               `json:"buildId,omitempty"`
	BuildState               *PackageBuildState    `json:"buildState,omitempty"`
	EnvironmentVariables     *map[string]string    `json:"environmentVariables,omitempty"`
	InferencingServer        InferencingServer     `json:"inferencingServer"`
	Inputs                   *[]ModelPackageInput  `json:"inputs,omitempty"`
	LogUrl                   *string               `json:"logUrl,omitempty"`
	ModelConfiguration       *ModelConfiguration   `json:"modelConfiguration,omitempty"`
	Tags                     *map[string]string    `json:"tags,omitempty"`
	TargetEnvironmentId      *string               `json:"targetEnvironmentId,omitempty"`
	TargetEnvironmentName    *string               `json:"targetEnvironmentName,omitempty"`
	TargetEnvironmentVersion *string               `json:"targetEnvironmentVersion,omitempty"`
}

var _ json.Unmarshaler = &PackageResponse{}

func (s *PackageResponse) UnmarshalJSON(bytes []byte) error {
	type alias PackageResponse
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into PackageResponse: %+v", err)
	}

	s.BuildId = decoded.BuildId
	s.BuildState = decoded.BuildState
	s.EnvironmentVariables = decoded.EnvironmentVariables
	s.Inputs = decoded.Inputs
	s.LogUrl = decoded.LogUrl
	s.ModelConfiguration = decoded.ModelConfiguration
	s.Tags = decoded.Tags
	s.TargetEnvironmentId = decoded.TargetEnvironmentId
	s.TargetEnvironmentName = decoded.TargetEnvironmentName
	s.TargetEnvironmentVersion = decoded.TargetEnvironmentVersion

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PackageResponse into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["baseEnvironmentSource"]; ok {
		impl, err := unmarshalBaseEnvironmentSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'BaseEnvironmentSource' for 'PackageResponse': %+v", err)
		}
		s.BaseEnvironmentSource = impl
	}

	if v, ok := temp["inferencingServer"]; ok {
		impl, err := unmarshalInferencingServerImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'InferencingServer' for 'PackageResponse': %+v", err)
		}
		s.InferencingServer = impl
	}
	return nil
}
