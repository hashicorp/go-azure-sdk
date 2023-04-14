package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutoMLVertical = ImageClassification{}

type ImageClassification struct {
	LimitSettings      *ImageLimitSettings                             `json:"limitSettings,omitempty"`
	ModelSettings      *ImageModelSettingsClassification               `json:"modelSettings,omitempty"`
	PrimaryMetric      *ClassificationPrimaryMetrics                   `json:"primaryMetric,omitempty"`
	SearchSpace        *[]ImageModelDistributionSettingsClassification `json:"searchSpace,omitempty"`
	SweepSettings      *ImageSweepSettings                             `json:"sweepSettings,omitempty"`
	ValidationData     *JobInput                                       `json:"validationData,omitempty"`
	ValidationDataSize *float64                                        `json:"validationDataSize,omitempty"`

	// Fields inherited from AutoMLVertical
	LogVerbosity     *LogVerbosity `json:"logVerbosity,omitempty"`
	TargetColumnName *string       `json:"targetColumnName,omitempty"`
	TrainingData     JobInput      `json:"trainingData"`
}

var _ json.Marshaler = ImageClassification{}

func (s ImageClassification) MarshalJSON() ([]byte, error) {
	type wrapper ImageClassification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImageClassification: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImageClassification: %+v", err)
	}
	decoded["taskType"] = "ImageClassification"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImageClassification: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ImageClassification{}

func (s *ImageClassification) UnmarshalJSON(bytes []byte) error {
	type alias ImageClassification
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ImageClassification: %+v", err)
	}

	s.LimitSettings = decoded.LimitSettings
	s.LogVerbosity = decoded.LogVerbosity
	s.ModelSettings = decoded.ModelSettings
	s.PrimaryMetric = decoded.PrimaryMetric
	s.SearchSpace = decoded.SearchSpace
	s.SweepSettings = decoded.SweepSettings
	s.TargetColumnName = decoded.TargetColumnName
	s.ValidationDataSize = decoded.ValidationDataSize

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ImageClassification into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["trainingData"]; ok {
		impl, err := unmarshalJobInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'TrainingData' for 'ImageClassification': %+v", err)
		}
		s.TrainingData = impl
	}

	if v, ok := temp["validationData"]; ok {
		impl, err := unmarshalJobInputImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ValidationData' for 'ImageClassification': %+v", err)
		}
		s.ValidationData = impl
	}
	return nil
}
