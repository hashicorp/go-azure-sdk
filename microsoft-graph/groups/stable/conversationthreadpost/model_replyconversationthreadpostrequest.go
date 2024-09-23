package conversationthreadpost

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplyConversationThreadPostRequest struct {
	Post *stable.Post `json:"Post,omitempty"`
}
