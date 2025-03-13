package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ FineTuningVertical = CustomModelFineTuning{}

type CustomModelFineTuning struct {
	HyperParameters *map[string]string `json:"hyperParameters,omitempty"`

	// Fields inherited from FineTuningVertical

	Model          JobInput           `json:"model"`
	ModelProvider  ModelProvider      `json:"modelProvider"`
	TaskType       FineTuningTaskType `json:"taskType"`
	TrainingData   JobInput           `json:"trainingData"`
	ValidationData JobInput           `json:"validationData"`
}

func (s CustomModelFineTuning) FineTuningVertical() BaseFineTuningVerticalImpl {
	return BaseFineTuningVerticalImpl{
		Model:          s.Model,
		ModelProvider:  s.ModelProvider,
		TaskType:       s.TaskType,
		TrainingData:   s.TrainingData,
		ValidationData: s.ValidationData,
	}
}

var _ json.Marshaler = CustomModelFineTuning{}

func (s CustomModelFineTuning) MarshalJSON() ([]byte, error) {
	type wrapper CustomModelFineTuning
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomModelFineTuning: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomModelFineTuning: %+v", err)
	}

	decoded["modelProvider"] = "Custom"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomModelFineTuning: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CustomModelFineTuning{}

func (s *CustomModelFineTuning) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		HyperParameters *map[string]string `json:"hyperParameters,omitempty"`
		ModelProvider   ModelProvider      `json:"modelProvider"`
		TaskType        FineTuningTaskType `json:"taskType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.HyperParameters = decoded.HyperParameters
	s.ModelProvider = decoded.ModelProvider
	s.TaskType = decoded.TaskType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CustomModelFineTuning into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["model"]; ok {
		impl, err := UnmarshalJobInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Model' for 'CustomModelFineTuning': %+v", err)
		}
		s.Model = impl
	}

	if v, ok := temp["trainingData"]; ok {
		impl, err := UnmarshalJobInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TrainingData' for 'CustomModelFineTuning': %+v", err)
		}
		s.TrainingData = impl
	}

	if v, ok := temp["validationData"]; ok {
		impl, err := UnmarshalJobInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ValidationData' for 'CustomModelFineTuning': %+v", err)
		}
		s.ValidationData = impl
	}

	return nil
}
