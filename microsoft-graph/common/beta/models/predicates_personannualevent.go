package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAnnualEventOperationPredicate struct {
	CreatedDateTime      *string
	Date                 *string
	DisplayName          *string
	Id                   *string
	IsSearchable         *bool
	LastModifiedDateTime *string
	ODataType            *string
}

func (p PersonAnnualEventOperationPredicate) Matches(input PersonAnnualEvent) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Date != nil && (input.Date == nil || *p.Date != *input.Date) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSearchable != nil && (input.IsSearchable == nil || *p.IsSearchable != *input.IsSearchable) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
