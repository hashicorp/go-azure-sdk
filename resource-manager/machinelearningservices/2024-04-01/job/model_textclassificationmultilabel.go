package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutoMLVertical = TextClassificationMultilabel{}

type TextClassificationMultilabel struct {
	FeaturizationSettings *FeaturizationSettings                  `json:"featurizationSettings,omitempty"`
	LimitSettings         *NlpVerticalLimitSettings               `json:"limitSettings,omitempty"`
	PrimaryMetric         *ClassificationMultilabelPrimaryMetrics `json:"primaryMetric,omitempty"`
	ValidationData        *MLTableJobInput                        `json:"validationData,omitempty"`

	// Fields inherited from AutoMLVertical

	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TaskType         TaskType        `json:"taskType"`
	TrainingData     MLTableJobInput `json:"trainingData"`
}

func (s TextClassificationMultilabel) AutoMLVertical() BaseAutoMLVerticalImpl {
	return BaseAutoMLVerticalImpl{
		LogVerbosity:     s.LogVerbosity,
		TargetColumnName: s.TargetColumnName,
		TaskType:         s.TaskType,
		TrainingData:     s.TrainingData,
	}
}

var _ json.Marshaler = TextClassificationMultilabel{}

func (s TextClassificationMultilabel) MarshalJSON() ([]byte, error) {
	type wrapper TextClassificationMultilabel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TextClassificationMultilabel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TextClassificationMultilabel: %+v", err)
	}

	decoded["taskType"] = "TextClassificationMultilabel"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TextClassificationMultilabel: %+v", err)
	}

	return encoded, nil
}
