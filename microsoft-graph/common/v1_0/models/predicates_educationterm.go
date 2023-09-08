package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationTermOperationPredicate struct {
	DisplayName *string
	EndDate     *string
	ExternalId  *string
	ODataType   *string
	StartDate   *string
}

func (p EducationTermOperationPredicate) Matches(input EducationTerm) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.EndDate != nil && (input.EndDate == nil || *p.EndDate != *input.EndDate) {
		return false
	}

	if p.ExternalId != nil && (input.ExternalId == nil || *p.ExternalId != *input.ExternalId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDate != nil && (input.StartDate == nil || *p.StartDate != *input.StartDate) {
		return false
	}

	return true
}
