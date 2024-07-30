package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoTask struct {
	AttachmentSessions       *[]AttachmentSession `json:"attachmentSessions,omitempty"`
	Attachments              *[]AttachmentBase    `json:"attachments,omitempty"`
	Body                     *ItemBody            `json:"body,omitempty"`
	BodyLastModifiedDateTime *string              `json:"bodyLastModifiedDateTime,omitempty"`
	Categories               *[]string            `json:"categories,omitempty"`
	ChecklistItems           *[]ChecklistItem     `json:"checklistItems,omitempty"`
	CompletedDateTime        *DateTimeTimeZone    `json:"completedDateTime,omitempty"`
	CreatedDateTime          *string              `json:"createdDateTime,omitempty"`
	DueDateTime              *DateTimeTimeZone    `json:"dueDateTime,omitempty"`
	Extensions               *[]Extension         `json:"extensions,omitempty"`
	HasAttachments           *bool                `json:"hasAttachments,omitempty"`
	Id                       *string              `json:"id,omitempty"`
	Importance               *TodoTaskImportance  `json:"importance,omitempty"`
	IsReminderOn             *bool                `json:"isReminderOn,omitempty"`
	LastModifiedDateTime     *string              `json:"lastModifiedDateTime,omitempty"`
	LinkedResources          *[]LinkedResource    `json:"linkedResources,omitempty"`
	ODataType                *string              `json:"@odata.type,omitempty"`
	Recurrence               *PatternedRecurrence `json:"recurrence,omitempty"`
	ReminderDateTime         *DateTimeTimeZone    `json:"reminderDateTime,omitempty"`
	StartDateTime            *DateTimeTimeZone    `json:"startDateTime,omitempty"`
	Status                   *TodoTaskStatus      `json:"status,omitempty"`
	Title                    *string              `json:"title,omitempty"`
}
