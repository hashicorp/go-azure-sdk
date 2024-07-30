package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageHistoryItem struct {
	Actions          *ChatMessageHistoryItemActions `json:"actions,omitempty"`
	ModifiedDateTime *string                        `json:"modifiedDateTime,omitempty"`
	ODataType        *string                        `json:"@odata.type,omitempty"`
	Reaction         *ChatMessageReaction           `json:"reaction,omitempty"`
}
