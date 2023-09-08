package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceOperationPredicate struct {
	EndDateTime   *string
	Id            *string
	ODataType     *string
	StartDateTime *string
	Status        *string
}

func (p AccessReviewInstanceOperationPredicate) Matches(input AccessReviewInstance) bool {

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
