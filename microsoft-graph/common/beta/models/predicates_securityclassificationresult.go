package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityClassificationResultOperationPredicate struct {
	ConfidenceLevel *int64
	Count           *int64
	ODataType       *string
	SensitiveTypeId *string
}

func (p SecurityClassificationResultOperationPredicate) Matches(input SecurityClassificationResult) bool {

	if p.ConfidenceLevel != nil && (input.ConfidenceLevel == nil || *p.ConfidenceLevel != *input.ConfidenceLevel) {
		return false
	}

	if p.Count != nil && (input.Count == nil || *p.Count != *input.Count) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SensitiveTypeId != nil && (input.SensitiveTypeId == nil || *p.SensitiveTypeId != *input.SensitiveTypeId) {
		return false
	}

	return true
}
