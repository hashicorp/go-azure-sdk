package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobBase = FineTuningJob{}

type FineTuningJob struct {
	FineTuningDetails FineTuningVertical   `json:"fineTuningDetails"`
	Outputs           map[string]JobOutput `json:"outputs"`
	QueueSettings     *QueueSettings       `json:"queueSettings,omitempty"`
	Resources         *JobResources        `json:"resources,omitempty"`

	// Fields inherited from JobBase

	ComponentId         *string                `json:"componentId,omitempty"`
	ComputeId           *string                `json:"computeId,omitempty"`
	Description         *string                `json:"description,omitempty"`
	DisplayName         *string                `json:"displayName,omitempty"`
	ExperimentName      *string                `json:"experimentName,omitempty"`
	Identity            IdentityConfiguration  `json:"identity"`
	IsArchived          *bool                  `json:"isArchived,omitempty"`
	JobType             JobType                `json:"jobType"`
	NotificationSetting *NotificationSetting   `json:"notificationSetting,omitempty"`
	Properties          *map[string]string     `json:"properties,omitempty"`
	Services            *map[string]JobService `json:"services,omitempty"`
	Status              *JobStatus             `json:"status,omitempty"`
	Tags                *map[string]string     `json:"tags,omitempty"`
}

func (s FineTuningJob) JobBase() BaseJobBaseImpl {
	return BaseJobBaseImpl{
		ComponentId:         s.ComponentId,
		ComputeId:           s.ComputeId,
		Description:         s.Description,
		DisplayName:         s.DisplayName,
		ExperimentName:      s.ExperimentName,
		Identity:            s.Identity,
		IsArchived:          s.IsArchived,
		JobType:             s.JobType,
		NotificationSetting: s.NotificationSetting,
		Properties:          s.Properties,
		Services:            s.Services,
		Status:              s.Status,
		Tags:                s.Tags,
	}
}

var _ json.Marshaler = FineTuningJob{}

func (s FineTuningJob) MarshalJSON() ([]byte, error) {
	type wrapper FineTuningJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FineTuningJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FineTuningJob: %+v", err)
	}

	decoded["jobType"] = "FineTuning"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FineTuningJob: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &FineTuningJob{}

func (s *FineTuningJob) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		QueueSettings       *QueueSettings         `json:"queueSettings,omitempty"`
		Resources           *JobResources          `json:"resources,omitempty"`
		ComponentId         *string                `json:"componentId,omitempty"`
		ComputeId           *string                `json:"computeId,omitempty"`
		Description         *string                `json:"description,omitempty"`
		DisplayName         *string                `json:"displayName,omitempty"`
		ExperimentName      *string                `json:"experimentName,omitempty"`
		IsArchived          *bool                  `json:"isArchived,omitempty"`
		JobType             JobType                `json:"jobType"`
		NotificationSetting *NotificationSetting   `json:"notificationSetting,omitempty"`
		Properties          *map[string]string     `json:"properties,omitempty"`
		Services            *map[string]JobService `json:"services,omitempty"`
		Status              *JobStatus             `json:"status,omitempty"`
		Tags                *map[string]string     `json:"tags,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.QueueSettings = decoded.QueueSettings
	s.Resources = decoded.Resources
	s.ComponentId = decoded.ComponentId
	s.ComputeId = decoded.ComputeId
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExperimentName = decoded.ExperimentName
	s.IsArchived = decoded.IsArchived
	s.JobType = decoded.JobType
	s.NotificationSetting = decoded.NotificationSetting
	s.Properties = decoded.Properties
	s.Services = decoded.Services
	s.Status = decoded.Status
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling FineTuningJob into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["fineTuningDetails"]; ok {
		impl, err := UnmarshalFineTuningVerticalImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'FineTuningDetails' for 'FineTuningJob': %+v", err)
		}
		s.FineTuningDetails = impl
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentityConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'FineTuningJob': %+v", err)
		}
		s.Identity = impl
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
				return fmt.Errorf("unmarshaling key %q field 'Outputs' for 'FineTuningJob': %+v", key, err)
			}
			output[key] = impl
		}
		s.Outputs = output
	}

	return nil
}
