package job

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureOpenAiHyperParameters struct {
	BatchSize              *int64   `json:"batchSize,omitempty"`
	LearningRateMultiplier *float64 `json:"learningRateMultiplier,omitempty"`
	NEpochs                *int64   `json:"nEpochs,omitempty"`
}
