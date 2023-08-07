package labelingjob

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutoMLVertical = ImageObjectDetection{}

type ImageObjectDetection struct {
	LimitSettings      *ImageLimitSettings                              `json:"limitSettings,omitempty"`
	ModelSettings      *ImageModelSettingsObjectDetection               `json:"modelSettings,omitempty"`
	PrimaryMetric      *ObjectDetectionPrimaryMetrics                   `json:"primaryMetric,omitempty"`
	SearchSpace        *[]ImageModelDistributionSettingsObjectDetection `json:"searchSpace,omitempty"`
	SweepSettings      *ImageSweepSettings                              `json:"sweepSettings,omitempty"`
	ValidationData     *MLTableJobInput                                 `json:"validationData,omitempty"`
	ValidationDataSize *float64                                         `json:"validationDataSize,omitempty"`

	// Fields inherited from AutoMLVertical
	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TrainingData     MLTableJobInput `json:"trainingData"`
}

var _ json.Marshaler = ImageObjectDetection{}

func (s ImageObjectDetection) MarshalJSON() ([]byte, error) {
	type wrapper ImageObjectDetection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImageObjectDetection: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImageObjectDetection: %+v", err)
	}
	decoded["taskType"] = "ImageObjectDetection"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImageObjectDetection: %+v", err)
	}

	return encoded, nil
}
