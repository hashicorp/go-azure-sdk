package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobBase interface {
	JobBase() BaseJobBaseImpl
}

var _ JobBase = BaseJobBaseImpl{}

type BaseJobBaseImpl struct {
	ComponentId    *string                `json:"componentId,omitempty"`
	ComputeId      *string                `json:"computeId,omitempty"`
	Description    *string                `json:"description,omitempty"`
	DisplayName    *string                `json:"displayName,omitempty"`
	ExperimentName *string                `json:"experimentName,omitempty"`
	Identity       IdentityConfiguration  `json:"identity"`
	IsArchived     *bool                  `json:"isArchived,omitempty"`
	JobType        JobType                `json:"jobType"`
	Properties     *map[string]string     `json:"properties,omitempty"`
	Services       *map[string]JobService `json:"services,omitempty"`
	Status         *JobStatus             `json:"status,omitempty"`
	Tags           *map[string]string     `json:"tags,omitempty"`
}

func (s BaseJobBaseImpl) JobBase() BaseJobBaseImpl {
	return s
}

var _ JobBase = RawJobBaseImpl{}

// RawJobBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobBaseImpl struct {
	jobBase BaseJobBaseImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawJobBaseImpl) JobBase() BaseJobBaseImpl {
	return s.jobBase
}

var _ json.Unmarshaler = &BaseJobBaseImpl{}

func (s *BaseJobBaseImpl) UnmarshalJSON(bytes []byte) error {
	type alias BaseJobBaseImpl
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into BaseJobBaseImpl: %+v", err)
	}

	s.ComponentId = decoded.ComponentId
	s.ComputeId = decoded.ComputeId
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExperimentName = decoded.ExperimentName
	s.IsArchived = decoded.IsArchived
	s.JobType = decoded.JobType
	s.Properties = decoded.Properties
	s.Services = decoded.Services
	s.Status = decoded.Status
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseJobBaseImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentityConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'BaseJobBaseImpl': %+v", err)
		}
		s.Identity = impl
	}
	return nil
}

func UnmarshalJobBaseImplementation(input []byte) (JobBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobBase into map[string]interface: %+v", err)
	}

	value, ok := temp["jobType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AutoML") {
		var out AutoMLJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutoMLJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Command") {
		var out CommandJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommandJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Pipeline") {
		var out PipelineJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PipelineJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Sweep") {
		var out SweepJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SweepJob: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobBaseImpl: %+v", err)
	}

	return RawJobBaseImpl{
		jobBase: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
