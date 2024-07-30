package threadpost

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Post struct {
	Attachments                   *[]Attachment                        `json:"attachments,omitempty"`
	Body                          *ItemBody                            `json:"body,omitempty"`
	Categories                    *[]string                            `json:"categories,omitempty"`
	ChangeKey                     *string                              `json:"changeKey,omitempty"`
	ConversationId                *string                              `json:"conversationId,omitempty"`
	ConversationThreadId          *string                              `json:"conversationThreadId,omitempty"`
	CreatedDateTime               *string                              `json:"createdDateTime,omitempty"`
	Extensions                    *[]Extension                         `json:"extensions,omitempty"`
	From                          *Recipient                           `json:"from,omitempty"`
	HasAttachments                *bool                                `json:"hasAttachments,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	Importance                    *PostImportance                      `json:"importance,omitempty"`
	InReplyTo                     *Post                                `json:"inReplyTo,omitempty"`
	LastModifiedDateTime          *string                              `json:"lastModifiedDateTime,omitempty"`
	Mentions                      *[]Mention                           `json:"mentions,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	NewParticipants               *[]Recipient                         `json:"newParticipants,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	ReceivedDateTime              *string                              `json:"receivedDateTime,omitempty"`
	Sender                        *Recipient                           `json:"sender,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
}
