package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatInfo struct {
	MessageId           *string `json:"messageId,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	ReplyChainMessageId *string `json:"replyChainMessageId,omitempty"`
	ThreadId            *string `json:"threadId,omitempty"`
}
