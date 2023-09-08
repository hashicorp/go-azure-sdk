package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentResourceOperationPredicate struct {
	DistributeForStudentWork *bool
	Id                       *string
	ODataType                *string
}

func (p EducationAssignmentResourceOperationPredicate) Matches(input EducationAssignmentResource) bool {

	if p.DistributeForStudentWork != nil && (input.DistributeForStudentWork == nil || *p.DistributeForStudentWork != *input.DistributeForStudentWork) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
