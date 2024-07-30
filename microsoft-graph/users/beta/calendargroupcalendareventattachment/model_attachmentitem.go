package calendargroupcalendareventattachment

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentItem struct {
	AttachmentType *AttachmentItemAttachmentType `json:"attachmentType,omitempty"`
	ContentId      *string                       `json:"contentId,omitempty"`
	ContentType    *string                       `json:"contentType,omitempty"`
	IsInline       *bool                         `json:"isInline,omitempty"`
	Name           *string                       `json:"name,omitempty"`
	ODataType      *string                       `json:"@odata.type,omitempty"`
	Size           *int64                        `json:"size,omitempty"`
}
