package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamSummaryOperationPredicate struct {
	GuestsCount  *int64
	MembersCount *int64
	ODataType    *string
	OwnersCount  *int64
}

func (p TeamSummaryOperationPredicate) Matches(input TeamSummary) bool {

	if p.GuestsCount != nil && (input.GuestsCount == nil || *p.GuestsCount != *input.GuestsCount) {
		return false
	}

	if p.MembersCount != nil && (input.MembersCount == nil || *p.MembersCount != *input.MembersCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OwnersCount != nil && (input.OwnersCount == nil || *p.OwnersCount != *input.OwnersCount) {
		return false
	}

	return true
}
