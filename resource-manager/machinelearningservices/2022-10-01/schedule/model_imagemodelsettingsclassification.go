package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageModelSettingsClassification struct {
	AdvancedSettings           *string                `json:"advancedSettings,omitempty"`
	AmsGradient                *bool                  `json:"amsGradient,omitempty"`
	Augmentations              *string                `json:"augmentations,omitempty"`
	Beta1                      *float64               `json:"beta1,omitempty"`
	Beta2                      *float64               `json:"beta2,omitempty"`
	CheckpointFrequency        *int64                 `json:"checkpointFrequency,omitempty"`
	CheckpointModel            *MLFlowModelJobInput   `json:"checkpointModel,omitempty"`
	CheckpointRunId            *string                `json:"checkpointRunId,omitempty"`
	Distributed                *bool                  `json:"distributed,omitempty"`
	EarlyStopping              *bool                  `json:"earlyStopping,omitempty"`
	EarlyStoppingDelay         *int64                 `json:"earlyStoppingDelay,omitempty"`
	EarlyStoppingPatience      *int64                 `json:"earlyStoppingPatience,omitempty"`
	EnableOnnxNormalization    *bool                  `json:"enableOnnxNormalization,omitempty"`
	EvaluationFrequency        *int64                 `json:"evaluationFrequency,omitempty"`
	GradientAccumulationStep   *int64                 `json:"gradientAccumulationStep,omitempty"`
	LayersToFreeze             *int64                 `json:"layersToFreeze,omitempty"`
	LearningRate               *float64               `json:"learningRate,omitempty"`
	LearningRateScheduler      *LearningRateScheduler `json:"learningRateScheduler,omitempty"`
	ModelName                  *string                `json:"modelName,omitempty"`
	Momentum                   *float64               `json:"momentum,omitempty"`
	Nesterov                   *bool                  `json:"nesterov,omitempty"`
	NumberOfEpochs             *int64                 `json:"numberOfEpochs,omitempty"`
	NumberOfWorkers            *int64                 `json:"numberOfWorkers,omitempty"`
	Optimizer                  *StochasticOptimizer   `json:"optimizer,omitempty"`
	RandomSeed                 *int64                 `json:"randomSeed,omitempty"`
	StepLRGamma                *float64               `json:"stepLRGamma,omitempty"`
	StepLRStepSize             *int64                 `json:"stepLRStepSize,omitempty"`
	TrainingBatchSize          *int64                 `json:"trainingBatchSize,omitempty"`
	TrainingCropSize           *int64                 `json:"trainingCropSize,omitempty"`
	ValidationBatchSize        *int64                 `json:"validationBatchSize,omitempty"`
	ValidationCropSize         *int64                 `json:"validationCropSize,omitempty"`
	ValidationResizeSize       *int64                 `json:"validationResizeSize,omitempty"`
	WarmupCosineLRCycles       *float64               `json:"warmupCosineLRCycles,omitempty"`
	WarmupCosineLRWarmupEpochs *int64                 `json:"warmupCosineLRWarmupEpochs,omitempty"`
	WeightDecay                *float64               `json:"weightDecay,omitempty"`
	WeightedLoss               *int64                 `json:"weightedLoss,omitempty"`
}
