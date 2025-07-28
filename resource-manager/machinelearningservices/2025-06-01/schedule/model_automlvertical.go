package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutoMLVertical interface {
	AutoMLVertical() BaseAutoMLVerticalImpl
}

var _ AutoMLVertical = BaseAutoMLVerticalImpl{}

type BaseAutoMLVerticalImpl struct {
	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TaskType         TaskType        `json:"taskType"`
	TrainingData     MLTableJobInput `json:"trainingData"`
}

func (s BaseAutoMLVerticalImpl) AutoMLVertical() BaseAutoMLVerticalImpl {
	return s
}

var _ AutoMLVertical = RawAutoMLVerticalImpl{}

// RawAutoMLVerticalImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAutoMLVerticalImpl struct {
	autoMLVertical BaseAutoMLVerticalImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawAutoMLVerticalImpl) AutoMLVertical() BaseAutoMLVerticalImpl {
	return s.autoMLVertical
}

func UnmarshalAutoMLVerticalImplementation(input []byte) (AutoMLVertical, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AutoMLVertical into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["taskType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Classification") {
		var out Classification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Classification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Forecasting") {
		var out Forecasting
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Forecasting: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ImageClassification") {
		var out ImageClassification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImageClassification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ImageClassificationMultilabel") {
		var out ImageClassificationMultilabel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImageClassificationMultilabel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ImageInstanceSegmentation") {
		var out ImageInstanceSegmentation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImageInstanceSegmentation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ImageObjectDetection") {
		var out ImageObjectDetection
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImageObjectDetection: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Regression") {
		var out Regression
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Regression: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TextClassification") {
		var out TextClassification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TextClassification: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TextClassificationMultilabel") {
		var out TextClassificationMultilabel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TextClassificationMultilabel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TextNER") {
		var out TextNer
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TextNer: %+v", err)
		}
		return out, nil
	}

	var parent BaseAutoMLVerticalImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAutoMLVerticalImpl: %+v", err)
	}

	return RawAutoMLVerticalImpl{
		autoMLVertical: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
