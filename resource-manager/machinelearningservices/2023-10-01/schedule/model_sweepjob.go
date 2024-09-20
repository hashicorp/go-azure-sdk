package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobBase = SweepJob{}

type SweepJob struct {
	EarlyTermination  EarlyTerminationPolicy `json:"earlyTermination"`
	Inputs            *map[string]JobInput   `json:"inputs,omitempty"`
	Limits            *SweepJobLimits        `json:"limits,omitempty"`
	Objective         Objective              `json:"objective"`
	Outputs           *map[string]JobOutput  `json:"outputs,omitempty"`
	QueueSettings     *QueueSettings         `json:"queueSettings,omitempty"`
	SamplingAlgorithm SamplingAlgorithm      `json:"samplingAlgorithm"`
	SearchSpace       interface{}            `json:"searchSpace"`
	Trial             TrialComponent         `json:"trial"`

	// Fields inherited from JobBase

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

func (s SweepJob) JobBase() BaseJobBaseImpl {
	return BaseJobBaseImpl{
		ComponentId:    s.ComponentId,
		ComputeId:      s.ComputeId,
		Description:    s.Description,
		DisplayName:    s.DisplayName,
		ExperimentName: s.ExperimentName,
		Identity:       s.Identity,
		IsArchived:     s.IsArchived,
		JobType:        s.JobType,
		Properties:     s.Properties,
		Services:       s.Services,
		Status:         s.Status,
		Tags:           s.Tags,
	}
}

var _ json.Marshaler = SweepJob{}

func (s SweepJob) MarshalJSON() ([]byte, error) {
	type wrapper SweepJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SweepJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SweepJob: %+v", err)
	}

	decoded["jobType"] = "Sweep"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SweepJob: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SweepJob{}

func (s *SweepJob) UnmarshalJSON(bytes []byte) error {
	type alias SweepJob
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SweepJob: %+v", err)
	}

	s.ComponentId = decoded.ComponentId
	s.ComputeId = decoded.ComputeId
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExperimentName = decoded.ExperimentName
	s.IsArchived = decoded.IsArchived
	s.JobType = decoded.JobType
	s.Limits = decoded.Limits
	s.Objective = decoded.Objective
	s.Properties = decoded.Properties
	s.QueueSettings = decoded.QueueSettings
	s.SearchSpace = decoded.SearchSpace
	s.Services = decoded.Services
	s.Status = decoded.Status
	s.Tags = decoded.Tags
	s.Trial = decoded.Trial

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SweepJob into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["earlyTermination"]; ok {
		impl, err := UnmarshalEarlyTerminationPolicyImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EarlyTermination' for 'SweepJob': %+v", err)
		}
		s.EarlyTermination = impl
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentityConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'SweepJob': %+v", err)
		}
		s.Identity = impl
	}

	if v, ok := temp["inputs"]; ok {
		var dictionaryTemp map[string]json.RawMessage
		if err := json.Unmarshal(v, &dictionaryTemp); err != nil {
			return fmt.Errorf("unmarshaling Inputs into dictionary map[string]json.RawMessage: %+v", err)
		}

		output := make(map[string]JobInput)
		for key, val := range dictionaryTemp {
			impl, err := UnmarshalJobInputImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'Inputs' for 'SweepJob': %+v", key, err)
			}
			output[key] = impl
		}
		s.Inputs = &output
	}

	if v, ok := temp["outputs"]; ok {
		var dictionaryTemp map[string]json.RawMessage
		if err := json.Unmarshal(v, &dictionaryTemp); err != nil {
			return fmt.Errorf("unmarshaling Outputs into dictionary map[string]json.RawMessage: %+v", err)
		}

		output := make(map[string]JobOutput)
		for key, val := range dictionaryTemp {
			impl, err := UnmarshalJobOutputImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'Outputs' for 'SweepJob': %+v", key, err)
			}
			output[key] = impl
		}
		s.Outputs = &output
	}

	if v, ok := temp["samplingAlgorithm"]; ok {
		impl, err := UnmarshalSamplingAlgorithmImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SamplingAlgorithm' for 'SweepJob': %+v", err)
		}
		s.SamplingAlgorithm = impl
	}
	return nil
}
