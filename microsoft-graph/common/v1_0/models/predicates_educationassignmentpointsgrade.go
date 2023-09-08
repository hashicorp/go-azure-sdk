package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentPointsGradeOperationPredicate struct {
	GradedDateTime *string
	ODataType      *string
}

func (p EducationAssignmentPointsGradeOperationPredicate) Matches(input EducationAssignmentPointsGrade) bool {

	if p.GradedDateTime != nil && (input.GradedDateTime == nil || *p.GradedDateTime != *input.GradedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
