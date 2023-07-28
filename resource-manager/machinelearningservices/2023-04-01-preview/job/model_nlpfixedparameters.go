package job

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NlpFixedParameters struct {
	GradientAccumulationSteps *int64                    `json:"gradientAccumulationSteps,omitempty"`
	LearningRate              *float64                  `json:"learningRate,omitempty"`
	LearningRateScheduler     *NlpLearningRateScheduler `json:"learningRateScheduler,omitempty"`
	ModelName                 *string                   `json:"modelName,omitempty"`
	NumberOfEpochs            *int64                    `json:"numberOfEpochs,omitempty"`
	TrainingBatchSize         *int64                    `json:"trainingBatchSize,omitempty"`
	ValidationBatchSize       *int64                    `json:"validationBatchSize,omitempty"`
	WarmupRatio               *float64                  `json:"warmupRatio,omitempty"`
	WeightDecay               *float64                  `json:"weightDecay,omitempty"`
}
