package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceAttachment struct {
	ContentType          *string                          `json:"contentType,omitempty"`
	Id                   *string                          `json:"id,omitempty"`
	IsFolder             *bool                            `json:"isFolder,omitempty"`
	IsInline             *bool                            `json:"isInline,omitempty"`
	LastModifiedDateTime *string                          `json:"lastModifiedDateTime,omitempty"`
	Name                 *string                          `json:"name,omitempty"`
	ODataType            *string                          `json:"@odata.type,omitempty"`
	Permission           *ReferenceAttachmentPermission   `json:"permission,omitempty"`
	PreviewUrl           *string                          `json:"previewUrl,omitempty"`
	ProviderType         *ReferenceAttachmentProviderType `json:"providerType,omitempty"`
	Size                 *int64                           `json:"size,omitempty"`
	SourceUrl            *string                          `json:"sourceUrl,omitempty"`
	ThumbnailUrl         *string                          `json:"thumbnailUrl,omitempty"`
}
