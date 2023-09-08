package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestHistoryOperationPredicate struct {
	EventDateTime *string
	ODataType     *string
	Type          *string
}

func (p SubjectRightsRequestHistoryOperationPredicate) Matches(input SubjectRightsRequestHistory) bool {

	if p.EventDateTime != nil && (input.EventDateTime == nil || *p.EventDateTime != *input.EventDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	return true
}
