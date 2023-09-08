package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookDocumentTaskChangeOperationPredicate struct {
	CommentId       *string
	CreatedDateTime *string
	DueDateTime     *string
	Id              *string
	ODataType       *string
	PercentComplete *int64
	Priority        *int64
	StartDateTime   *string
	Title           *string
	Type            *string
	UndoChangeId    *string
}

func (p WorkbookDocumentTaskChangeOperationPredicate) Matches(input WorkbookDocumentTaskChange) bool {

	if p.CommentId != nil && (input.CommentId == nil || *p.CommentId != *input.CommentId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DueDateTime != nil && (input.DueDateTime == nil || *p.DueDateTime != *input.DueDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PercentComplete != nil && (input.PercentComplete == nil || *p.PercentComplete != *input.PercentComplete) {
		return false
	}

	if p.Priority != nil && (input.Priority == nil || *p.Priority != *input.Priority) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.Title != nil && (input.Title == nil || *p.Title != *input.Title) {
		return false
	}

	if p.Type != nil && (input.Type == nil || *p.Type != *input.Type) {
		return false
	}

	if p.UndoChangeId != nil && (input.UndoChangeId == nil || *p.UndoChangeId != *input.UndoChangeId) {
		return false
	}

	return true
}
