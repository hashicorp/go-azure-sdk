package labelingjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageModelSettingsObjectDetection struct {
	AdvancedSettings            *string                `json:"advancedSettings,omitempty"`
	AmsGradient                 *bool                  `json:"amsGradient,omitempty"`
	Augmentations               *string                `json:"augmentations,omitempty"`
	Beta1                       *float64               `json:"beta1,omitempty"`
	Beta2                       *float64               `json:"beta2,omitempty"`
	BoxDetectionsPerImage       *int64                 `json:"boxDetectionsPerImage,omitempty"`
	BoxScoreThreshold           *float64               `json:"boxScoreThreshold,omitempty"`
	CheckpointFrequency         *int64                 `json:"checkpointFrequency,omitempty"`
	CheckpointModel             *MLFlowModelJobInput   `json:"checkpointModel,omitempty"`
	CheckpointRunId             *string                `json:"checkpointRunId,omitempty"`
	Distributed                 *bool                  `json:"distributed,omitempty"`
	EarlyStopping               *bool                  `json:"earlyStopping,omitempty"`
	EarlyStoppingDelay          *int64                 `json:"earlyStoppingDelay,omitempty"`
	EarlyStoppingPatience       *int64                 `json:"earlyStoppingPatience,omitempty"`
	EnableOnnxNormalization     *bool                  `json:"enableOnnxNormalization,omitempty"`
	EvaluationFrequency         *int64                 `json:"evaluationFrequency,omitempty"`
	GradientAccumulationStep    *int64                 `json:"gradientAccumulationStep,omitempty"`
	ImageSize                   *int64                 `json:"imageSize,omitempty"`
	LayersToFreeze              *int64                 `json:"layersToFreeze,omitempty"`
	LearningRate                *float64               `json:"learningRate,omitempty"`
	LearningRateScheduler       *LearningRateScheduler `json:"learningRateScheduler,omitempty"`
	LogTrainingMetrics          *LogTrainingMetrics    `json:"logTrainingMetrics,omitempty"`
	LogValidationLoss           *LogValidationLoss     `json:"logValidationLoss,omitempty"`
	MaxSize                     *int64                 `json:"maxSize,omitempty"`
	MinSize                     *int64                 `json:"minSize,omitempty"`
	ModelName                   *string                `json:"modelName,omitempty"`
	ModelSize                   *ModelSize             `json:"modelSize,omitempty"`
	Momentum                    *float64               `json:"momentum,omitempty"`
	MultiScale                  *bool                  `json:"multiScale,omitempty"`
	Nesterov                    *bool                  `json:"nesterov,omitempty"`
	NmsIouThreshold             *float64               `json:"nmsIouThreshold,omitempty"`
	NumberOfEpochs              *int64                 `json:"numberOfEpochs,omitempty"`
	NumberOfWorkers             *int64                 `json:"numberOfWorkers,omitempty"`
	Optimizer                   *StochasticOptimizer   `json:"optimizer,omitempty"`
	RandomSeed                  *int64                 `json:"randomSeed,omitempty"`
	StepLRGamma                 *float64               `json:"stepLRGamma,omitempty"`
	StepLRStepSize              *int64                 `json:"stepLRStepSize,omitempty"`
	TileGridSize                *string                `json:"tileGridSize,omitempty"`
	TileOverlapRatio            *float64               `json:"tileOverlapRatio,omitempty"`
	TilePredictionsNmsThreshold *float64               `json:"tilePredictionsNmsThreshold,omitempty"`
	TrainingBatchSize           *int64                 `json:"trainingBatchSize,omitempty"`
	ValidationBatchSize         *int64                 `json:"validationBatchSize,omitempty"`
	ValidationIouThreshold      *float64               `json:"validationIouThreshold,omitempty"`
	ValidationMetricType        *ValidationMetricType  `json:"validationMetricType,omitempty"`
	WarmupCosineLRCycles        *float64               `json:"warmupCosineLRCycles,omitempty"`
	WarmupCosineLRWarmupEpochs  *int64                 `json:"warmupCosineLRWarmupEpochs,omitempty"`
	WeightDecay                 *float64               `json:"weightDecay,omitempty"`
}
