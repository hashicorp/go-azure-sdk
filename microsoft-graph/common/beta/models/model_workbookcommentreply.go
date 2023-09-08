package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookCommentReply struct {
	Content     *string               `json:"content,omitempty"`
	ContentType *string               `json:"contentType,omitempty"`
	Id          *string               `json:"id,omitempty"`
	ODataType   *string               `json:"@odata.type,omitempty"`
	Task        *WorkbookDocumentTask `json:"task,omitempty"`
}
