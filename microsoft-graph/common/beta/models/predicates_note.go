package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoteOperationPredicate struct {
	ChangeKey            *string
	CreatedDateTime      *string
	HasAttachments       *bool
	Id                   *string
	IsDeleted            *bool
	LastModifiedDateTime *string
	ODataType            *string
	Subject              *string
}

func (p NoteOperationPredicate) Matches(input Note) bool {

	if p.ChangeKey != nil && (input.ChangeKey == nil || *p.ChangeKey != *input.ChangeKey) {
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

	if p.IsDeleted != nil && (input.IsDeleted == nil || *p.IsDeleted != *input.IsDeleted) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
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
