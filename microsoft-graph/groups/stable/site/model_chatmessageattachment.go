package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageAttachment struct {
	Content      *string `json:"content,omitempty"`
	ContentType  *string `json:"contentType,omitempty"`
	ContentUrl   *string `json:"contentUrl,omitempty"`
	Id           *string `json:"id,omitempty"`
	Name         *string `json:"name,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	TeamsAppId   *string `json:"teamsAppId,omitempty"`
	ThumbnailUrl *string `json:"thumbnailUrl,omitempty"`
}
