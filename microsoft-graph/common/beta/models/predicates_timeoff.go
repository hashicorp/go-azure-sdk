package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffOperationPredicate struct {
	CreatedDateTime      *string
	Id                   *string
	IsStagedForDeletion  *bool
	LastModifiedDateTime *string
	ODataType            *string
	UserId               *string
}

func (p TimeOffOperationPredicate) Matches(input TimeOff) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsStagedForDeletion != nil && (input.IsStagedForDeletion == nil || *p.IsStagedForDeletion != *input.IsStagedForDeletion) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
