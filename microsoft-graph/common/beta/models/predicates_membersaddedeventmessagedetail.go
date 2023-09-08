package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MembersAddedEventMessageDetailOperationPredicate struct {
	ODataType                   *string
	VisibleHistoryStartDateTime *string
}

func (p MembersAddedEventMessageDetailOperationPredicate) Matches(input MembersAddedEventMessageDetail) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.VisibleHistoryStartDateTime != nil && (input.VisibleHistoryStartDateTime == nil || *p.VisibleHistoryStartDateTime != *input.VisibleHistoryStartDateTime) {
		return false
	}

	return true
}
