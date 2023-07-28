package labelingjob

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobBase = SparkJob{}

type SparkJob struct {
	Archives      *[]string                   `json:"archives,omitempty"`
	Args          *string                     `json:"args,omitempty"`
	CodeId        string                      `json:"codeId"`
	Conf          *map[string]string          `json:"conf,omitempty"`
	Entry         SparkJobEntry               `json:"entry"`
	EnvironmentId *string                     `json:"environmentId,omitempty"`
	Files         *[]string                   `json:"files,omitempty"`
	Inputs        *map[string]JobInput        `json:"inputs,omitempty"`
	Jars          *[]string                   `json:"jars,omitempty"`
	Outputs       *map[string]JobOutput       `json:"outputs,omitempty"`
	PyFiles       *[]string                   `json:"pyFiles,omitempty"`
	QueueSettings *QueueSettings              `json:"queueSettings,omitempty"`
	Resources     *SparkResourceConfiguration `json:"resources,omitempty"`

	// Fields inherited from JobBase
	ComponentId          *string                         `json:"componentId,omitempty"`
	ComputeId            *string                         `json:"computeId,omitempty"`
	Description          *string                         `json:"description,omitempty"`
	DisplayName          *string                         `json:"displayName,omitempty"`
	ExperimentName       *string                         `json:"experimentName,omitempty"`
	Identity             IdentityConfiguration           `json:"identity"`
	IsArchived           *bool                           `json:"isArchived,omitempty"`
	NotificationSetting  *NotificationSetting            `json:"notificationSetting,omitempty"`
	Properties           *map[string]string              `json:"properties,omitempty"`
	SecretsConfiguration *map[string]SecretConfiguration `json:"secretsConfiguration,omitempty"`
	Services             *map[string]JobService          `json:"services,omitempty"`
	Status               *JobStatus                      `json:"status,omitempty"`
	Tags                 *map[string]string              `json:"tags,omitempty"`
}

var _ json.Marshaler = SparkJob{}

func (s SparkJob) MarshalJSON() ([]byte, error) {
	type wrapper SparkJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SparkJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SparkJob: %+v", err)
	}
	decoded["jobType"] = "Spark"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SparkJob: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SparkJob{}

func (s *SparkJob) UnmarshalJSON(bytes []byte) error {
	type alias SparkJob
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into SparkJob: %+v", err)
	}

	s.Archives = decoded.Archives
	s.Args = decoded.Args
	s.CodeId = decoded.CodeId
	s.ComponentId = decoded.ComponentId
	s.ComputeId = decoded.ComputeId
	s.Conf = decoded.Conf
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.EnvironmentId = decoded.EnvironmentId
	s.ExperimentName = decoded.ExperimentName
	s.Files = decoded.Files
	s.IsArchived = decoded.IsArchived
	s.Jars = decoded.Jars
	s.NotificationSetting = decoded.NotificationSetting
	s.Properties = decoded.Properties
	s.PyFiles = decoded.PyFiles
	s.QueueSettings = decoded.QueueSettings
	s.Resources = decoded.Resources
	s.SecretsConfiguration = decoded.SecretsConfiguration
	s.Services = decoded.Services
	s.Status = decoded.Status
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SparkJob into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["entry"]; ok {
		impl, err := unmarshalSparkJobEntryImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Entry' for 'SparkJob': %+v", err)
		}
		s.Entry = impl
	}

	if v, ok := temp["identity"]; ok {
		impl, err := unmarshalIdentityConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'SparkJob': %+v", err)
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
			impl, err := unmarshalJobInputImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'Inputs' for 'SparkJob': %+v", key, err)
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
			impl, err := unmarshalJobOutputImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'Outputs' for 'SparkJob': %+v", key, err)
			}
			output[key] = impl
		}
		s.Outputs = &output
	}
	return nil
}
