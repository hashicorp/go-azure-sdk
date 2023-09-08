package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationSubmissionOperationPredicate struct {
	Id                  *string
	ODataType           *string
	ReassignedDateTime  *string
	ResourcesFolderUrl  *string
	ReturnedDateTime    *string
	SubmittedDateTime   *string
	UnsubmittedDateTime *string
	WebUrl              *string
}

func (p EducationSubmissionOperationPredicate) Matches(input EducationSubmission) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReassignedDateTime != nil && (input.ReassignedDateTime == nil || *p.ReassignedDateTime != *input.ReassignedDateTime) {
		return false
	}

	if p.ResourcesFolderUrl != nil && (input.ResourcesFolderUrl == nil || *p.ResourcesFolderUrl != *input.ResourcesFolderUrl) {
		return false
	}

	if p.ReturnedDateTime != nil && (input.ReturnedDateTime == nil || *p.ReturnedDateTime != *input.ReturnedDateTime) {
		return false
	}

	if p.SubmittedDateTime != nil && (input.SubmittedDateTime == nil || *p.SubmittedDateTime != *input.SubmittedDateTime) {
		return false
	}

	if p.UnsubmittedDateTime != nil && (input.UnsubmittedDateTime == nil || *p.UnsubmittedDateTime != *input.UnsubmittedDateTime) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
