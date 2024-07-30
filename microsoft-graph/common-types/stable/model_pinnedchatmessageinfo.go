package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PinnedChatMessageInfo struct {
	Id        *string      `json:"id,omitempty"`
	Message   *ChatMessage `json:"message,omitempty"`
	ODataType *string      `json:"@odata.type,omitempty"`
}
