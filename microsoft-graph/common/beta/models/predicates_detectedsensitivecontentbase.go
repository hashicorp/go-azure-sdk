package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedSensitiveContentBaseOperationPredicate struct {
	Confidence            *int64
	DisplayName           *string
	Id                    *string
	ODataType             *string
	RecommendedConfidence *int64
	UniqueCount           *int64
}

func (p DetectedSensitiveContentBaseOperationPredicate) Matches(input DetectedSensitiveContentBase) bool {

	if p.Confidence != nil && (input.Confidence == nil || *p.Confidence != *input.Confidence) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecommendedConfidence != nil && (input.RecommendedConfidence == nil || *p.RecommendedConfidence != *input.RecommendedConfidence) {
		return false
	}

	if p.UniqueCount != nil && (input.UniqueCount == nil || *p.UniqueCount != *input.UniqueCount) {
		return false
	}

	return true
}
