package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDispositionReviewStageOperationPredicate struct {
	Id          *string
	Name        *string
	ODataType   *string
	StageNumber *int64
}

func (p SecurityDispositionReviewStageOperationPredicate) Matches(input SecurityDispositionReviewStage) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StageNumber != nil && (input.StageNumber == nil || *p.StageNumber != *input.StageNumber) {
		return false
	}

	return true
}
