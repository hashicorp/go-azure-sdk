package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentClassificationOperationPredicate struct {
	Confidence      *int64
	ODataType       *string
	SensitiveTypeId *string
	UniqueCount     *int64
}

func (p ContentClassificationOperationPredicate) Matches(input ContentClassification) bool {

	if p.Confidence != nil && (input.Confidence == nil || *p.Confidence != *input.Confidence) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SensitiveTypeId != nil && (input.SensitiveTypeId == nil || *p.SensitiveTypeId != *input.SensitiveTypeId) {
		return false
	}

	if p.UniqueCount != nil && (input.UniqueCount == nil || *p.UniqueCount != *input.UniqueCount) {
		return false
	}

	return true
}
