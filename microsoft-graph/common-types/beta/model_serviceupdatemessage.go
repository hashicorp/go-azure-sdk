package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceUpdateMessage struct {
	ActionRequiredByDateTime *string                          `json:"actionRequiredByDateTime,omitempty"`
	Attachments              *[]ServiceAnnouncementAttachment `json:"attachments,omitempty"`
	AttachmentsArchive       *string                          `json:"attachmentsArchive,omitempty"`
	Body                     *ItemBody                        `json:"body,omitempty"`
	Category                 *ServiceUpdateMessageCategory    `json:"category,omitempty"`
	Details                  *[]KeyValuePair                  `json:"details,omitempty"`
	EndDateTime              *string                          `json:"endDateTime,omitempty"`
	HasAttachments           *bool                            `json:"hasAttachments,omitempty"`
	Id                       *string                          `json:"id,omitempty"`
	IsMajorChange            *bool                            `json:"isMajorChange,omitempty"`
	LastModifiedDateTime     *string                          `json:"lastModifiedDateTime,omitempty"`
	ODataType                *string                          `json:"@odata.type,omitempty"`
	Services                 *[]string                        `json:"services,omitempty"`
	Severity                 *ServiceUpdateMessageSeverity    `json:"severity,omitempty"`
	StartDateTime            *string                          `json:"startDateTime,omitempty"`
	Tags                     *[]string                        `json:"tags,omitempty"`
	Title                    *string                          `json:"title,omitempty"`
	ViewPoint                *ServiceUpdateMessageViewpoint   `json:"viewPoint,omitempty"`
}
