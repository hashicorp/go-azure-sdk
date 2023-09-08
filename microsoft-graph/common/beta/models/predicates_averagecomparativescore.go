package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AverageComparativeScoreOperationPredicate struct {
	Basis     *string
	ODataType *string
}

func (p AverageComparativeScoreOperationPredicate) Matches(input AverageComparativeScore) bool {

	if p.Basis != nil && (input.Basis == nil || *p.Basis != *input.Basis) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
