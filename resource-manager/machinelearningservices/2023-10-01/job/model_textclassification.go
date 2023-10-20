package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutoMLVertical = TextClassification{}

type TextClassification struct {
	FeaturizationSettings *FeaturizationSettings        `json:"featurizationSettings,omitempty"`
	LimitSettings         *NlpVerticalLimitSettings     `json:"limitSettings,omitempty"`
	PrimaryMetric         *ClassificationPrimaryMetrics `json:"primaryMetric,omitempty"`
	ValidationData        *MLTableJobInput              `json:"validationData,omitempty"`

	// Fields inherited from AutoMLVertical
	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TrainingData     MLTableJobInput `json:"trainingData"`
}

var _ json.Marshaler = TextClassification{}

func (s TextClassification) MarshalJSON() ([]byte, error) {
	type wrapper TextClassification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TextClassification: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TextClassification: %+v", err)
	}
	decoded["taskType"] = "TextClassification"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TextClassification: %+v", err)
	}

	return encoded, nil
}
