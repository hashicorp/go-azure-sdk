package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSchoolOperationPredicate struct {
	Description          *string
	DisplayName          *string
	ExternalId           *string
	ExternalPrincipalId  *string
	ExternalSourceDetail *string
	Fax                  *string
	HighestGrade         *string
	Id                   *string
	LowestGrade          *string
	ODataType            *string
	Phone                *string
	PrincipalEmail       *string
	PrincipalName        *string
	SchoolNumber         *string
}

func (p EducationSchoolOperationPredicate) Matches(input EducationSchool) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ExternalId != nil && (input.ExternalId == nil || *p.ExternalId != *input.ExternalId) {
		return false
	}

	if p.ExternalPrincipalId != nil && (input.ExternalPrincipalId == nil || *p.ExternalPrincipalId != *input.ExternalPrincipalId) {
		return false
	}

	if p.ExternalSourceDetail != nil && (input.ExternalSourceDetail == nil || *p.ExternalSourceDetail != *input.ExternalSourceDetail) {
		return false
	}

	if p.Fax != nil && (input.Fax == nil || *p.Fax != *input.Fax) {
		return false
	}

	if p.HighestGrade != nil && (input.HighestGrade == nil || *p.HighestGrade != *input.HighestGrade) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LowestGrade != nil && (input.LowestGrade == nil || *p.LowestGrade != *input.LowestGrade) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Phone != nil && (input.Phone == nil || *p.Phone != *input.Phone) {
		return false
	}

	if p.PrincipalEmail != nil && (input.PrincipalEmail == nil || *p.PrincipalEmail != *input.PrincipalEmail) {
		return false
	}

	if p.PrincipalName != nil && (input.PrincipalName == nil || *p.PrincipalName != *input.PrincipalName) {
		return false
	}

	if p.SchoolNumber != nil && (input.SchoolNumber == nil || *p.SchoolNumber != *input.SchoolNumber) {
		return false
	}

	return true
}
