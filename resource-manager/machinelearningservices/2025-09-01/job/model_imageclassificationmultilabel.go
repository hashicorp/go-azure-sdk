package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutoMLVertical = ImageClassificationMultilabel{}

type ImageClassificationMultilabel struct {
	LimitSettings      ImageLimitSettings                              `json:"limitSettings"`
	ModelSettings      *ImageModelSettingsClassification               `json:"modelSettings,omitempty"`
	PrimaryMetric      *ClassificationMultilabelPrimaryMetrics         `json:"primaryMetric,omitempty"`
	SearchSpace        *[]ImageModelDistributionSettingsClassification `json:"searchSpace,omitempty"`
	SweepSettings      *ImageSweepSettings                             `json:"sweepSettings,omitempty"`
	ValidationData     *MLTableJobInput                                `json:"validationData,omitempty"`
	ValidationDataSize *float64                                        `json:"validationDataSize,omitempty"`

	// Fields inherited from AutoMLVertical

	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TaskType         TaskType        `json:"taskType"`
	TrainingData     MLTableJobInput `json:"trainingData"`
}

func (s ImageClassificationMultilabel) AutoMLVertical() BaseAutoMLVerticalImpl {
	return BaseAutoMLVerticalImpl{
		LogVerbosity:     s.LogVerbosity,
		TargetColumnName: s.TargetColumnName,
		TaskType:         s.TaskType,
		TrainingData:     s.TrainingData,
	}
}

var _ json.Marshaler = ImageClassificationMultilabel{}

func (s ImageClassificationMultilabel) MarshalJSON() ([]byte, error) {
	type wrapper ImageClassificationMultilabel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImageClassificationMultilabel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImageClassificationMultilabel: %+v", err)
	}

	decoded["taskType"] = "ImageClassificationMultilabel"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImageClassificationMultilabel: %+v", err)
	}

	return encoded, nil
}
