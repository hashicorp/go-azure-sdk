package chatmessage

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplyChatMessagesWithQuoteRequest struct {
	MessageIds   *[]string         `json:"messageIds,omitempty"`
	ReplyMessage *beta.ChatMessage `json:"replyMessage,omitempty"`
}
