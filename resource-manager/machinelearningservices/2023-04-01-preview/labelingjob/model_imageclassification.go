package labelingjob

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
	ValidationData     *MLTableJobInput                                `json:"validationData,omitempty"`
	ValidationDataSize *float64                                        `json:"validationDataSize,omitempty"`

	// Fields inherited from AutoMLVertical
	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TrainingData     MLTableJobInput `json:"trainingData"`
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
