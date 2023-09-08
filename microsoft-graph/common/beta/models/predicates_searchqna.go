package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchQnaOperationPredicate struct {
	AvailabilityEndDateTime   *string
	AvailabilityStartDateTime *string
	Description               *string
	DisplayName               *string
	Id                        *string
	IsSuggested               *bool
	LastModifiedDateTime      *string
	ODataType                 *string
	WebUrl                    *string
}

func (p SearchQnaOperationPredicate) Matches(input SearchQna) bool {

	if p.AvailabilityEndDateTime != nil && (input.AvailabilityEndDateTime == nil || *p.AvailabilityEndDateTime != *input.AvailabilityEndDateTime) {
		return false
	}

	if p.AvailabilityStartDateTime != nil && (input.AvailabilityStartDateTime == nil || *p.AvailabilityStartDateTime != *input.AvailabilityStartDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSuggested != nil && (input.IsSuggested == nil || *p.IsSuggested != *input.IsSuggested) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
