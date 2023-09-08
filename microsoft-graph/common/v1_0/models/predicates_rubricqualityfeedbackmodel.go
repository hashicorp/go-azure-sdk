package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RubricQualityFeedbackModelOperationPredicate struct {
	ODataType *string
	QualityId *string
}

func (p RubricQualityFeedbackModelOperationPredicate) Matches(input RubricQualityFeedbackModel) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.QualityId != nil && (input.QualityId == nil || *p.QualityId != *input.QualityId) {
		return false
	}

	return true
}
