package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutoMLVertical = ImageInstanceSegmentation{}

type ImageInstanceSegmentation struct {
	LimitSettings      ImageLimitSettings                               `json:"limitSettings"`
	ModelSettings      *ImageModelSettingsObjectDetection               `json:"modelSettings,omitempty"`
	PrimaryMetric      *InstanceSegmentationPrimaryMetrics              `json:"primaryMetric,omitempty"`
	SearchSpace        *[]ImageModelDistributionSettingsObjectDetection `json:"searchSpace,omitempty"`
	SweepSettings      *ImageSweepSettings                              `json:"sweepSettings,omitempty"`
	ValidationData     *MLTableJobInput                                 `json:"validationData,omitempty"`
	ValidationDataSize *float64                                         `json:"validationDataSize,omitempty"`

	// Fields inherited from AutoMLVertical

	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TaskType         TaskType        `json:"taskType"`
	TrainingData     MLTableJobInput `json:"trainingData"`
}

func (s ImageInstanceSegmentation) AutoMLVertical() BaseAutoMLVerticalImpl {
	return BaseAutoMLVerticalImpl{
		LogVerbosity:     s.LogVerbosity,
		TargetColumnName: s.TargetColumnName,
		TaskType:         s.TaskType,
		TrainingData:     s.TrainingData,
	}
}

var _ json.Marshaler = ImageInstanceSegmentation{}

func (s ImageInstanceSegmentation) MarshalJSON() ([]byte, error) {
	type wrapper ImageInstanceSegmentation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImageInstanceSegmentation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImageInstanceSegmentation: %+v", err)
	}

	decoded["taskType"] = "ImageInstanceSegmentation"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImageInstanceSegmentation: %+v", err)
	}

	return encoded, nil
}
