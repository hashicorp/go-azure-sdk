package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTask struct {
	AssignedTo                    *string                              `json:"assignedTo,omitempty"`
	Attachments                   *[]Attachment                        `json:"attachments,omitempty"`
	Body                          *ItemBody                            `json:"body,omitempty"`
	Categories                    *[]string                            `json:"categories,omitempty"`
	ChangeKey                     *string                              `json:"changeKey,omitempty"`
	CompletedDateTime             *DateTimeTimeZone                    `json:"completedDateTime,omitempty"`
	CreatedDateTime               *string                              `json:"createdDateTime,omitempty"`
	DueDateTime                   *DateTimeTimeZone                    `json:"dueDateTime,omitempty"`
	HasAttachments                *bool                                `json:"hasAttachments,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	Importance                    *OutlookTaskImportance               `json:"importance,omitempty"`
	IsReminderOn                  *bool                                `json:"isReminderOn,omitempty"`
	LastModifiedDateTime          *string                              `json:"lastModifiedDateTime,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	Owner                         *string                              `json:"owner,omitempty"`
	ParentFolderId                *string                              `json:"parentFolderId,omitempty"`
	Recurrence                    *PatternedRecurrence                 `json:"recurrence,omitempty"`
	ReminderDateTime              *DateTimeTimeZone                    `json:"reminderDateTime,omitempty"`
	Sensitivity                   *OutlookTaskSensitivity              `json:"sensitivity,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	StartDateTime                 *DateTimeTimeZone                    `json:"startDateTime,omitempty"`
	Status                        *OutlookTaskStatus                   `json:"status,omitempty"`
	Subject                       *string                              `json:"subject,omitempty"`
}
