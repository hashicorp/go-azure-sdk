package job

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AutoMLVertical = Classification{}

type Classification struct {
	CvSplitColumnNames    *[]string                           `json:"cvSplitColumnNames,omitempty"`
	FeaturizationSettings *TableVerticalFeaturizationSettings `json:"featurizationSettings,omitempty"`
	LimitSettings         *TableVerticalLimitSettings         `json:"limitSettings,omitempty"`
	NCrossValidations     NCrossValidations                   `json:"nCrossValidations"`
	PositiveLabel         *string                             `json:"positiveLabel,omitempty"`
	PrimaryMetric         *ClassificationPrimaryMetrics       `json:"primaryMetric,omitempty"`
	TestData              *MLTableJobInput                    `json:"testData,omitempty"`
	TestDataSize          *float64                            `json:"testDataSize,omitempty"`
	TrainingSettings      *ClassificationTrainingSettings     `json:"trainingSettings,omitempty"`
	ValidationData        *MLTableJobInput                    `json:"validationData,omitempty"`
	ValidationDataSize    *float64                            `json:"validationDataSize,omitempty"`
	WeightColumnName      *string                             `json:"weightColumnName,omitempty"`

	// Fields inherited from AutoMLVertical

	LogVerbosity     *LogVerbosity   `json:"logVerbosity,omitempty"`
	TargetColumnName *string         `json:"targetColumnName,omitempty"`
	TaskType         TaskType        `json:"taskType"`
	TrainingData     MLTableJobInput `json:"trainingData"`
}

func (s Classification) AutoMLVertical() BaseAutoMLVerticalImpl {
	return BaseAutoMLVerticalImpl{
		LogVerbosity:     s.LogVerbosity,
		TargetColumnName: s.TargetColumnName,
		TaskType:         s.TaskType,
		TrainingData:     s.TrainingData,
	}
}

var _ json.Marshaler = Classification{}

func (s Classification) MarshalJSON() ([]byte, error) {
	type wrapper Classification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Classification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Classification: %+v", err)
	}

	decoded["taskType"] = "Classification"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Classification: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Classification{}

func (s *Classification) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CvSplitColumnNames    *[]string                           `json:"cvSplitColumnNames,omitempty"`
		FeaturizationSettings *TableVerticalFeaturizationSettings `json:"featurizationSettings,omitempty"`
		LimitSettings         *TableVerticalLimitSettings         `json:"limitSettings,omitempty"`
		PositiveLabel         *string                             `json:"positiveLabel,omitempty"`
		PrimaryMetric         *ClassificationPrimaryMetrics       `json:"primaryMetric,omitempty"`
		TestData              *MLTableJobInput                    `json:"testData,omitempty"`
		TestDataSize          *float64                            `json:"testDataSize,omitempty"`
		TrainingSettings      *ClassificationTrainingSettings     `json:"trainingSettings,omitempty"`
		ValidationData        *MLTableJobInput                    `json:"validationData,omitempty"`
		ValidationDataSize    *float64                            `json:"validationDataSize,omitempty"`
		WeightColumnName      *string                             `json:"weightColumnName,omitempty"`
		LogVerbosity          *LogVerbosity                       `json:"logVerbosity,omitempty"`
		TargetColumnName      *string                             `json:"targetColumnName,omitempty"`
		TaskType              TaskType                            `json:"taskType"`
		TrainingData          MLTableJobInput                     `json:"trainingData"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CvSplitColumnNames = decoded.CvSplitColumnNames
	s.FeaturizationSettings = decoded.FeaturizationSettings
	s.LimitSettings = decoded.LimitSettings
	s.PositiveLabel = decoded.PositiveLabel
	s.PrimaryMetric = decoded.PrimaryMetric
	s.TestData = decoded.TestData
	s.TestDataSize = decoded.TestDataSize
	s.TrainingSettings = decoded.TrainingSettings
	s.ValidationData = decoded.ValidationData
	s.ValidationDataSize = decoded.ValidationDataSize
	s.WeightColumnName = decoded.WeightColumnName
	s.LogVerbosity = decoded.LogVerbosity
	s.TargetColumnName = decoded.TargetColumnName
	s.TaskType = decoded.TaskType
	s.TrainingData = decoded.TrainingData

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Classification into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["nCrossValidations"]; ok {
		impl, err := UnmarshalNCrossValidationsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'NCrossValidations' for 'Classification': %+v", err)
		}
		s.NCrossValidations = impl
	}

	return nil
}
