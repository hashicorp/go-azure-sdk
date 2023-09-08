package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationInnerErrorOperationPredicate struct {
	ActivityId      *string
	ClientRequestId *string
	Code            *string
	ErrorDateTime   *string
	ODataType       *string
}

func (p ClassificationInnerErrorOperationPredicate) Matches(input ClassificationInnerError) bool {

	if p.ActivityId != nil && (input.ActivityId == nil || *p.ActivityId != *input.ActivityId) {
		return false
	}

	if p.ClientRequestId != nil && (input.ClientRequestId == nil || *p.ClientRequestId != *input.ClientRequestId) {
		return false
	}

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.ErrorDateTime != nil && (input.ErrorDateTime == nil || *p.ErrorDateTime != *input.ErrorDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
