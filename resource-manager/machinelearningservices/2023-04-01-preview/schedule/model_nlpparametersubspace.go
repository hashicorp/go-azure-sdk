package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NlpParameterSubspace struct {
	GradientAccumulationSteps *string `json:"gradientAccumulationSteps,omitempty"`
	LearningRate              *string `json:"learningRate,omitempty"`
	LearningRateScheduler     *string `json:"learningRateScheduler,omitempty"`
	ModelName                 *string `json:"modelName,omitempty"`
	NumberOfEpochs            *string `json:"numberOfEpochs,omitempty"`
	TrainingBatchSize         *string `json:"trainingBatchSize,omitempty"`
	ValidationBatchSize       *string `json:"validationBatchSize,omitempty"`
	WarmupRatio               *string `json:"warmupRatio,omitempty"`
	WeightDecay               *string `json:"weightDecay,omitempty"`
}
