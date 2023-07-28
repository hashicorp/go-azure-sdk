package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobBase = LabelingJob{}

type LabelingJob struct {
	CreatedDateTime            *string                    `json:"createdDateTime,omitempty"`
	DataConfiguration          *LabelingDataConfiguration `json:"dataConfiguration,omitempty"`
	JobInstructions            *LabelingJobInstructions   `json:"jobInstructions,omitempty"`
	LabelCategories            *map[string]LabelCategory  `json:"labelCategories,omitempty"`
	LabelingJobMediaProperties LabelingJobMediaProperties `json:"labelingJobMediaProperties"`
	MlAssistConfiguration      MLAssistConfiguration      `json:"mlAssistConfiguration"`
	ProgressMetrics            *ProgressMetrics           `json:"progressMetrics,omitempty"`
	ProjectId                  *string                    `json:"projectId,omitempty"`
	ProvisioningState          *JobProvisioningState      `json:"provisioningState,omitempty"`
	StatusMessages             *[]StatusMessage           `json:"statusMessages,omitempty"`

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

var _ json.Marshaler = LabelingJob{}

func (s LabelingJob) MarshalJSON() ([]byte, error) {
	type wrapper LabelingJob
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LabelingJob: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LabelingJob: %+v", err)
	}
	decoded["jobType"] = "Labeling"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LabelingJob: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &LabelingJob{}

func (s *LabelingJob) UnmarshalJSON(bytes []byte) error {
	type alias LabelingJob
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into LabelingJob: %+v", err)
	}

	s.ComponentId = decoded.ComponentId
	s.ComputeId = decoded.ComputeId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DataConfiguration = decoded.DataConfiguration
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExperimentName = decoded.ExperimentName
	s.IsArchived = decoded.IsArchived
	s.JobInstructions = decoded.JobInstructions
	s.LabelCategories = decoded.LabelCategories
	s.NotificationSetting = decoded.NotificationSetting
	s.ProgressMetrics = decoded.ProgressMetrics
	s.ProjectId = decoded.ProjectId
	s.Properties = decoded.Properties
	s.ProvisioningState = decoded.ProvisioningState
	s.SecretsConfiguration = decoded.SecretsConfiguration
	s.Services = decoded.Services
	s.Status = decoded.Status
	s.StatusMessages = decoded.StatusMessages
	s.Tags = decoded.Tags

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LabelingJob into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := unmarshalIdentityConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'LabelingJob': %+v", err)
		}
		s.Identity = impl
	}

	if v, ok := temp["labelingJobMediaProperties"]; ok {
		impl, err := unmarshalLabelingJobMediaPropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LabelingJobMediaProperties' for 'LabelingJob': %+v", err)
		}
		s.LabelingJobMediaProperties = impl
	}

	if v, ok := temp["mlAssistConfiguration"]; ok {
		impl, err := unmarshalMLAssistConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MlAssistConfiguration' for 'LabelingJob': %+v", err)
		}
		s.MlAssistConfiguration = impl
	}
	return nil
}
