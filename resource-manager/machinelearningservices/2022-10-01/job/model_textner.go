package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutoMLVertical = TextNer{}

type TextNer struct {
	FeaturizationSettings *FeaturizationSettings        `json:"featurizationSettings,omitempty"`
	LimitSettings         *NlpVerticalLimitSettings     `json:"limitSettings,omitempty"`
	PrimaryMetric         *ClassificationPrimaryMetrics `json:"primaryMetric,omitempty"`
	ValidationData        *MLTableJobInput              `json:"validationData,omitempty"`

	// Fields inherited from AutoMLVertical
	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TrainingData     MLTableJobInput `json:"trainingData"`
}

var _ json.Marshaler = TextNer{}

func (s TextNer) MarshalJSON() ([]byte, error) {
	type wrapper TextNer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TextNer: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TextNer: %+v", err)
	}
	decoded["taskType"] = "TextNER"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TextNer: %+v", err)
	}

	return encoded, nil
}
