package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentSettingsOperationPredicate struct {
	Id                          *string
	ODataType                   *string
	SubmissionAnimationDisabled *bool
}

func (p EducationAssignmentSettingsOperationPredicate) Matches(input EducationAssignmentSettings) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SubmissionAnimationDisabled != nil && (input.SubmissionAnimationDisabled == nil || *p.SubmissionAnimationDisabled != *input.SubmissionAnimationDisabled) {
		return false
	}

	return true
}
