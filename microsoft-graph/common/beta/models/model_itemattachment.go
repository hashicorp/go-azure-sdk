package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemAttachment struct {
	ContentType          *string      `json:"contentType,omitempty"`
	Id                   *string      `json:"id,omitempty"`
	IsInline             *bool        `json:"isInline,omitempty"`
	Item                 *OutlookItem `json:"item,omitempty"`
	LastModifiedDateTime *string      `json:"lastModifiedDateTime,omitempty"`
	Name                 *string      `json:"name,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
	Size                 *int64       `json:"size,omitempty"`
}
