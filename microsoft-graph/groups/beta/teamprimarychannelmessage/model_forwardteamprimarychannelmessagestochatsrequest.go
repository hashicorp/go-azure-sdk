package teamprimarychannelmessage

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForwardTeamPrimaryChannelMessagesToChatsRequest struct {
	AdditionalMessage *beta.ChatMessage `json:"additionalMessage,omitempty"`
	MessageIds        *[]string         `json:"messageIds,omitempty"`
	TargetChatIds     *[]string         `json:"targetChatIds,omitempty"`
}
