package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LearningAssignmentOperationPredicate struct {
	AssignedDateTime         *string
	AssignerUserId           *string
	CompletedDateTime        *string
	CompletionPercentage     *int64
	ExternalcourseActivityId *string
	Id                       *string
	LearnerUserId            *string
	LearningContentId        *string
	LearningProviderId       *string
	ODataType                *string
}

func (p LearningAssignmentOperationPredicate) Matches(input LearningAssignment) bool {

	if p.AssignedDateTime != nil && (input.AssignedDateTime == nil || *p.AssignedDateTime != *input.AssignedDateTime) {
		return false
	}

	if p.AssignerUserId != nil && (input.AssignerUserId == nil || *p.AssignerUserId != *input.AssignerUserId) {
		return false
	}

	if p.CompletedDateTime != nil && (input.CompletedDateTime == nil || *p.CompletedDateTime != *input.CompletedDateTime) {
		return false
	}

	if p.CompletionPercentage != nil && (input.CompletionPercentage == nil || *p.CompletionPercentage != *input.CompletionPercentage) {
		return false
	}

	if p.ExternalcourseActivityId != nil && (input.ExternalcourseActivityId == nil || *p.ExternalcourseActivityId != *input.ExternalcourseActivityId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LearnerUserId != nil && (input.LearnerUserId == nil || *p.LearnerUserId != *input.LearnerUserId) {
		return false
	}

	if p.LearningContentId != nil && (input.LearningContentId == nil || *p.LearningContentId != *input.LearningContentId) {
		return false
	}

	if p.LearningProviderId != nil && (input.LearningProviderId == nil || *p.LearningProviderId != *input.LearningProviderId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
