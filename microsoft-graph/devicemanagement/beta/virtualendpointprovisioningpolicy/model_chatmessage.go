package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessage struct {
	Attachments          *[]ChatMessageAttachment    `json:"attachments,omitempty"`
	Body                 *ItemBody                   `json:"body,omitempty"`
	ChannelIdentity      *ChannelIdentity            `json:"channelIdentity,omitempty"`
	ChatId               *string                     `json:"chatId,omitempty"`
	CreatedDateTime      *string                     `json:"createdDateTime,omitempty"`
	DeletedDateTime      *string                     `json:"deletedDateTime,omitempty"`
	Etag                 *string                     `json:"etag,omitempty"`
	EventDetail          *EventMessageDetail         `json:"eventDetail,omitempty"`
	From                 *ChatMessageFromIdentitySet `json:"from,omitempty"`
	HostedContents       *[]ChatMessageHostedContent `json:"hostedContents,omitempty"`
	Id                   *string                     `json:"id,omitempty"`
	Importance           *ChatMessageImportance      `json:"importance,omitempty"`
	LastEditedDateTime   *string                     `json:"lastEditedDateTime,omitempty"`
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	Locale               *string                     `json:"locale,omitempty"`
	Mentions             *[]ChatMessageMention       `json:"mentions,omitempty"`
	MessageHistory       *[]ChatMessageHistoryItem   `json:"messageHistory,omitempty"`
	MessageType          *ChatMessageMessageType     `json:"messageType,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	OnBehalfOf           *ChatMessageFromIdentitySet `json:"onBehalfOf,omitempty"`
	PolicyViolation      *ChatMessagePolicyViolation `json:"policyViolation,omitempty"`
	Reactions            *[]ChatMessageReaction      `json:"reactions,omitempty"`
	Replies              *[]ChatMessage              `json:"replies,omitempty"`
	ReplyToId            *string                     `json:"replyToId,omitempty"`
	Subject              *string                     `json:"subject,omitempty"`
	Summary              *string                     `json:"summary,omitempty"`
	WebUrl               *string                     `json:"webUrl,omitempty"`
}
