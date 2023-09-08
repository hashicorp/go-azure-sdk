package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoTaskOperationPredicate struct {
	BodyLastModifiedDateTime *string
	CreatedDateTime          *string
	HasAttachments           *bool
	Id                       *string
	IsReminderOn             *bool
	LastModifiedDateTime     *string
	ODataType                *string
	Title                    *string
}

func (p TodoTaskOperationPredicate) Matches(input TodoTask) bool {

	if p.BodyLastModifiedDateTime != nil && (input.BodyLastModifiedDateTime == nil || *p.BodyLastModifiedDateTime != *input.BodyLastModifiedDateTime) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.HasAttachments != nil && (input.HasAttachments == nil || *p.HasAttachments != *input.HasAttachments) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsReminderOn != nil && (input.IsReminderOn == nil || *p.IsReminderOn != *input.IsReminderOn) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
