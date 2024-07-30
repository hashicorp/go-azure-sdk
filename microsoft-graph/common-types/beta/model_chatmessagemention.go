package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageMention struct {
	Id          *int64                           `json:"id,omitempty"`
	MentionText *string                          `json:"mentionText,omitempty"`
	Mentioned   *ChatMessageMentionedIdentitySet `json:"mentioned,omitempty"`
	ODataType   *string                          `json:"@odata.type,omitempty"`
}
