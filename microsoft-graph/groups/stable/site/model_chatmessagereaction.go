package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageReaction struct {
	CreatedDateTime *string                         `json:"createdDateTime,omitempty"`
	ODataType       *string                         `json:"@odata.type,omitempty"`
	ReactionType    *string                         `json:"reactionType,omitempty"`
	User            *ChatMessageReactionIdentitySet `json:"user,omitempty"`
}
