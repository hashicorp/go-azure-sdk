package labelingjob

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TableParameterSubspace struct {
	Booster          *string `json:"booster,omitempty"`
	BoostingType     *string `json:"boostingType,omitempty"`
	GrowPolicy       *string `json:"growPolicy,omitempty"`
	LearningRate     *string `json:"learningRate,omitempty"`
	MaxBin           *string `json:"maxBin,omitempty"`
	MaxDepth         *string `json:"maxDepth,omitempty"`
	MaxLeaves        *string `json:"maxLeaves,omitempty"`
	MinDataInLeaf    *string `json:"minDataInLeaf,omitempty"`
	MinSplitGain     *string `json:"minSplitGain,omitempty"`
	ModelName        *string `json:"modelName,omitempty"`
	NEstimators      *string `json:"nEstimators,omitempty"`
	NumLeaves        *string `json:"numLeaves,omitempty"`
	PreprocessorName *string `json:"preprocessorName,omitempty"`
	RegAlpha         *string `json:"regAlpha,omitempty"`
	RegLambda        *string `json:"regLambda,omitempty"`
	Subsample        *string `json:"subsample,omitempty"`
	SubsampleFreq    *string `json:"subsampleFreq,omitempty"`
	TreeMethod       *string `json:"treeMethod,omitempty"`
	WithMean         *string `json:"withMean,omitempty"`
	WithStd          *string `json:"withStd,omitempty"`
}
