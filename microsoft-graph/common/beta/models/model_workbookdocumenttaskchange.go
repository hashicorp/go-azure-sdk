package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookDocumentTaskChange struct {
	Assignee        *WorkbookEmailIdentity `json:"assignee,omitempty"`
	ChangedBy       *WorkbookEmailIdentity `json:"changedBy,omitempty"`
	CommentId       *string                `json:"commentId,omitempty"`
	CreatedDateTime *string                `json:"createdDateTime,omitempty"`
	DueDateTime     *string                `json:"dueDateTime,omitempty"`
	Id              *string                `json:"id,omitempty"`
	ODataType       *string                `json:"@odata.type,omitempty"`
	PercentComplete *int64                 `json:"percentComplete,omitempty"`
	Priority        *int64                 `json:"priority,omitempty"`
	StartDateTime   *string                `json:"startDateTime,omitempty"`
	Title           *string                `json:"title,omitempty"`
	Type            *string                `json:"type,omitempty"`
	UndoChangeId    *string                `json:"undoChangeId,omitempty"`
}
