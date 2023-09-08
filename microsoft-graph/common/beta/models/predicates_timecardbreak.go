package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeCardBreakOperationPredicate struct {
	BreakId   *string
	ODataType *string
}

func (p TimeCardBreakOperationPredicate) Matches(input TimeCardBreak) bool {

	if p.BreakId != nil && (input.BreakId == nil || *p.BreakId != *input.BreakId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
