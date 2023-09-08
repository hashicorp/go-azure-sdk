package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationCourseOperationPredicate struct {
	CourseNumber *string
	Description  *string
	DisplayName  *string
	ExternalId   *string
	ODataType    *string
	Subject      *string
}

func (p EducationCourseOperationPredicate) Matches(input EducationCourse) bool {

	if p.CourseNumber != nil && (input.CourseNumber == nil || *p.CourseNumber != *input.CourseNumber) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ExternalId != nil && (input.ExternalId == nil || *p.ExternalId != *input.ExternalId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	return true
}
