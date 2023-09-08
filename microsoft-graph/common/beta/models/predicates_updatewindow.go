package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowOperationPredicate struct {
	ODataType             *string
	UpdateWindowEndTime   *string
	UpdateWindowStartTime *string
}

func (p UpdateWindowOperationPredicate) Matches(input UpdateWindow) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UpdateWindowEndTime != nil && (input.UpdateWindowEndTime == nil || *p.UpdateWindowEndTime != *input.UpdateWindowEndTime) {
		return false
	}

	if p.UpdateWindowStartTime != nil && (input.UpdateWindowStartTime == nil || *p.UpdateWindowStartTime != *input.UpdateWindowStartTime) {
		return false
	}

	return true
}
