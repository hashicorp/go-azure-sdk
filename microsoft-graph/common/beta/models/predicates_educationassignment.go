package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentOperationPredicate struct {
	AllowLateSubmissions                    *bool
	AllowStudentsToAddResourcesToSubmission *bool
	AssignDateTime                          *string
	AssignedDateTime                        *string
	ClassId                                 *string
	CloseDateTime                           *string
	CreatedDateTime                         *string
	DisplayName                             *string
	DueDateTime                             *string
	FeedbackResourcesFolderUrl              *string
	Id                                      *string
	LastModifiedDateTime                    *string
	ModuleUrl                               *string
	NotificationChannelUrl                  *string
	ODataType                               *string
	ResourcesFolderUrl                      *string
	WebUrl                                  *string
}

func (p EducationAssignmentOperationPredicate) Matches(input EducationAssignment) bool {

	if p.AllowLateSubmissions != nil && (input.AllowLateSubmissions == nil || *p.AllowLateSubmissions != *input.AllowLateSubmissions) {
		return false
	}

	if p.AllowStudentsToAddResourcesToSubmission != nil && (input.AllowStudentsToAddResourcesToSubmission == nil || *p.AllowStudentsToAddResourcesToSubmission != *input.AllowStudentsToAddResourcesToSubmission) {
		return false
	}

	if p.AssignDateTime != nil && (input.AssignDateTime == nil || *p.AssignDateTime != *input.AssignDateTime) {
		return false
	}

	if p.AssignedDateTime != nil && (input.AssignedDateTime == nil || *p.AssignedDateTime != *input.AssignedDateTime) {
		return false
	}

	if p.ClassId != nil && (input.ClassId == nil || *p.ClassId != *input.ClassId) {
		return false
	}

	if p.CloseDateTime != nil && (input.CloseDateTime == nil || *p.CloseDateTime != *input.CloseDateTime) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.DueDateTime != nil && (input.DueDateTime == nil || *p.DueDateTime != *input.DueDateTime) {
		return false
	}

	if p.FeedbackResourcesFolderUrl != nil && (input.FeedbackResourcesFolderUrl == nil || *p.FeedbackResourcesFolderUrl != *input.FeedbackResourcesFolderUrl) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ModuleUrl != nil && (input.ModuleUrl == nil || *p.ModuleUrl != *input.ModuleUrl) {
		return false
	}

	if p.NotificationChannelUrl != nil && (input.NotificationChannelUrl == nil || *p.NotificationChannelUrl != *input.NotificationChannelUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResourcesFolderUrl != nil && (input.ResourcesFolderUrl == nil || *p.ResourcesFolderUrl != *input.ResourcesFolderUrl) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
