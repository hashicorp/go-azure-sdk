package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookDocumentTask struct {
	Assignees           *[]WorkbookEmailIdentity      `json:"assignees,omitempty"`
	Changes             *[]WorkbookDocumentTaskChange `json:"changes,omitempty"`
	Comment             *WorkbookComment              `json:"comment,omitempty"`
	CompletedBy         *WorkbookEmailIdentity        `json:"completedBy,omitempty"`
	CompletedDateTime   *string                       `json:"completedDateTime,omitempty"`
	CreatedBy           *WorkbookEmailIdentity        `json:"createdBy,omitempty"`
	CreatedDateTime     *string                       `json:"createdDateTime,omitempty"`
	Id                  *string                       `json:"id,omitempty"`
	ODataType           *string                       `json:"@odata.type,omitempty"`
	PercentComplete     *int64                        `json:"percentComplete,omitempty"`
	Priority            *int64                        `json:"priority,omitempty"`
	StartAndDueDateTime *WorkbookDocumentTaskSchedule `json:"startAndDueDateTime,omitempty"`
	Title               *string                       `json:"title,omitempty"`
}
