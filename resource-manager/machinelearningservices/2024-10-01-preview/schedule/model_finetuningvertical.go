package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FineTuningVertical interface {
	FineTuningVertical() BaseFineTuningVerticalImpl
}

var _ FineTuningVertical = BaseFineTuningVerticalImpl{}

type BaseFineTuningVerticalImpl struct {
	Model          JobInput           `json:"model"`
	ModelProvider  ModelProvider      `json:"modelProvider"`
	TaskType       FineTuningTaskType `json:"taskType"`
	TrainingData   JobInput           `json:"trainingData"`
	ValidationData JobInput           `json:"validationData"`
}

func (s BaseFineTuningVerticalImpl) FineTuningVertical() BaseFineTuningVerticalImpl {
	return s
}

var _ FineTuningVertical = RawFineTuningVerticalImpl{}

// RawFineTuningVerticalImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawFineTuningVerticalImpl struct {
	fineTuningVertical BaseFineTuningVerticalImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawFineTuningVerticalImpl) FineTuningVertical() BaseFineTuningVerticalImpl {
	return s.fineTuningVertical
}

var _ json.Unmarshaler = &BaseFineTuningVerticalImpl{}

func (s *BaseFineTuningVerticalImpl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ModelProvider ModelProvider      `json:"modelProvider"`
		TaskType      FineTuningTaskType `json:"taskType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ModelProvider = decoded.ModelProvider
	s.TaskType = decoded.TaskType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling BaseFineTuningVerticalImpl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["model"]; ok {
		impl, err := UnmarshalJobInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Model' for 'BaseFineTuningVerticalImpl': %+v", err)
		}
		s.Model = impl
	}

	if v, ok := temp["trainingData"]; ok {
		impl, err := UnmarshalJobInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TrainingData' for 'BaseFineTuningVerticalImpl': %+v", err)
		}
		s.TrainingData = impl
	}

	if v, ok := temp["validationData"]; ok {
		impl, err := UnmarshalJobInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ValidationData' for 'BaseFineTuningVerticalImpl': %+v", err)
		}
		s.ValidationData = impl
	}

	return nil
}

func UnmarshalFineTuningVerticalImplementation(input []byte) (FineTuningVertical, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling FineTuningVertical into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["modelProvider"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AzureOpenAI") {
		var out AzureOpenAiFineTuning
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureOpenAiFineTuning: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out CustomModelFineTuning
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomModelFineTuning: %+v", err)
		}
		return out, nil
	}

	var parent BaseFineTuningVerticalImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseFineTuningVerticalImpl: %+v", err)
	}

	return RawFineTuningVerticalImpl{
		fineTuningVertical: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
