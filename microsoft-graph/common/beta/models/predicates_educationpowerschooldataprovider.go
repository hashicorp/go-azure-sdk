package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationPowerSchoolDataProviderOperationPredicate struct {
	AllowTeachersInMultipleSchools *bool
	ClientId                       *string
	ClientSecret                   *string
	ConnectionUrl                  *string
	ODataType                      *string
	SchoolYear                     *string
}

func (p EducationPowerSchoolDataProviderOperationPredicate) Matches(input EducationPowerSchoolDataProvider) bool {

	if p.AllowTeachersInMultipleSchools != nil && (input.AllowTeachersInMultipleSchools == nil || *p.AllowTeachersInMultipleSchools != *input.AllowTeachersInMultipleSchools) {
		return false
	}

	if p.ClientId != nil && (input.ClientId == nil || *p.ClientId != *input.ClientId) {
		return false
	}

	if p.ClientSecret != nil && (input.ClientSecret == nil || *p.ClientSecret != *input.ClientSecret) {
		return false
	}

	if p.ConnectionUrl != nil && (input.ConnectionUrl == nil || *p.ConnectionUrl != *input.ConnectionUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SchoolYear != nil && (input.SchoolYear == nil || *p.SchoolYear != *input.SchoolYear) {
		return false
	}

	return true
}
