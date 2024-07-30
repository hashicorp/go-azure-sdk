package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Picture struct {
	Content     *string `json:"content,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Height      *int64  `json:"height,omitempty"`
	Id          *string `json:"id,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Width       *int64  `json:"width,omitempty"`
}
