package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenotePage struct {
	Content              *string         `json:"content,omitempty"`
	ContentUrl           *string         `json:"contentUrl,omitempty"`
	CreatedByAppId       *string         `json:"createdByAppId,omitempty"`
	CreatedDateTime      *string         `json:"createdDateTime,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	LastModifiedDateTime *string         `json:"lastModifiedDateTime,omitempty"`
	Level                *int64          `json:"level,omitempty"`
	Links                *PageLinks      `json:"links,omitempty"`
	ODataType            *string         `json:"@odata.type,omitempty"`
	Order                *int64          `json:"order,omitempty"`
	ParentNotebook       *Notebook       `json:"parentNotebook,omitempty"`
	ParentSection        *OnenoteSection `json:"parentSection,omitempty"`
	Self                 *string         `json:"self,omitempty"`
	Title                *string         `json:"title,omitempty"`
	UserTags             *[]string       `json:"userTags,omitempty"`
}
