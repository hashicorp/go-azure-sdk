package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffItemOperationPredicate struct {
	EndDateTime     *string
	ODataType       *string
	StartDateTime   *string
	TimeOffReasonId *string
}

func (p TimeOffItemOperationPredicate) Matches(input TimeOffItem) bool {

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.TimeOffReasonId != nil && (input.TimeOffReasonId == nil || *p.TimeOffReasonId != *input.TimeOffReasonId) {
		return false
	}

	return true
}
