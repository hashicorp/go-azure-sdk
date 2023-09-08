package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryRedundancyDetectionSettingsOperationPredicate struct {
	IsEnabled           *bool
	MaxWords            *int64
	MinWords            *int64
	ODataType           *string
	SimilarityThreshold *int64
}

func (p EdiscoveryRedundancyDetectionSettingsOperationPredicate) Matches(input EdiscoveryRedundancyDetectionSettings) bool {

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.MaxWords != nil && (input.MaxWords == nil || *p.MaxWords != *input.MaxWords) {
		return false
	}

	if p.MinWords != nil && (input.MinWords == nil || *p.MinWords != *input.MinWords) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SimilarityThreshold != nil && (input.SimilarityThreshold == nil || *p.SimilarityThreshold != *input.SimilarityThreshold) {
		return false
	}

	return true
}
