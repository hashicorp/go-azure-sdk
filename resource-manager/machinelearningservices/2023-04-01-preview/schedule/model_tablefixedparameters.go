package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TableFixedParameters struct {
	Booster          *string  `json:"booster,omitempty"`
	BoostingType     *string  `json:"boostingType,omitempty"`
	GrowPolicy       *string  `json:"growPolicy,omitempty"`
	LearningRate     *float64 `json:"learningRate,omitempty"`
	MaxBin           *int64   `json:"maxBin,omitempty"`
	MaxDepth         *int64   `json:"maxDepth,omitempty"`
	MaxLeaves        *int64   `json:"maxLeaves,omitempty"`
	MinDataInLeaf    *int64   `json:"minDataInLeaf,omitempty"`
	MinSplitGain     *float64 `json:"minSplitGain,omitempty"`
	ModelName        *string  `json:"modelName,omitempty"`
	NEstimators      *int64   `json:"nEstimators,omitempty"`
	NumLeaves        *int64   `json:"numLeaves,omitempty"`
	PreprocessorName *string  `json:"preprocessorName,omitempty"`
	RegAlpha         *float64 `json:"regAlpha,omitempty"`
	RegLambda        *float64 `json:"regLambda,omitempty"`
	Subsample        *float64 `json:"subsample,omitempty"`
	SubsampleFreq    *float64 `json:"subsampleFreq,omitempty"`
	TreeMethod       *string  `json:"treeMethod,omitempty"`
	WithMean         *bool    `json:"withMean,omitempty"`
	WithStd          *bool    `json:"withStd,omitempty"`
}
