package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentInfo struct {
	AttachmentType *AttachmentInfoAttachmentType `json:"attachmentType,omitempty"`
	ContentType    *string                       `json:"contentType,omitempty"`
	Name           *string                       `json:"name,omitempty"`
	ODataType      *string                       `json:"@odata.type,omitempty"`
	Size           *int64                        `json:"size,omitempty"`
}
