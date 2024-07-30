package threadpost

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemBody struct {
	Content     *string              `json:"content,omitempty"`
	ContentType *ItemBodyContentType `json:"contentType,omitempty"`
	ODataType   *string              `json:"@odata.type,omitempty"`
}
