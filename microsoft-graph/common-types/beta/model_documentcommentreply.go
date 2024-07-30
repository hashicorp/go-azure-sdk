package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentCommentReply struct {
	Content   *string `json:"content,omitempty"`
	Id        *string `json:"id,omitempty"`
	Location  *string `json:"location,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
