package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessageOperationPredicate struct {
	ActionRequiredByDateTime *string
	AttachmentsArchive       *string
	EndDateTime              *string
	HasAttachments           *bool
	Id                       *string
	IsMajorChange            *bool
	LastModifiedDateTime     *string
	ODataType                *string
	StartDateTime            *string
	Title                    *string
}

func (p ServiceUpdateMessageOperationPredicate) Matches(input ServiceUpdateMessage) bool {

	if p.ActionRequiredByDateTime != nil && (input.ActionRequiredByDateTime == nil || *p.ActionRequiredByDateTime != *input.ActionRequiredByDateTime) {
		return false
	}

	if p.AttachmentsArchive != nil && (input.AttachmentsArchive == nil || *p.AttachmentsArchive != *input.AttachmentsArchive) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.HasAttachments != nil && (input.HasAttachments == nil || *p.HasAttachments != *input.HasAttachments) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsMajorChange != nil && (input.IsMajorChange == nil || *p.IsMajorChange != *input.IsMajorChange) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	return true
}
