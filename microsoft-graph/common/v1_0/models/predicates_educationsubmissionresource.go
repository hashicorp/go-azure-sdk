package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSubmissionResourceOperationPredicate struct {
	AssignmentResourceUrl *string
	Id                    *string
	ODataType             *string
}

func (p EducationSubmissionResourceOperationPredicate) Matches(input EducationSubmissionResource) bool {

	if p.AssignmentResourceUrl != nil && (input.AssignmentResourceUrl == nil || *p.AssignmentResourceUrl != *input.AssignmentResourceUrl) {
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
