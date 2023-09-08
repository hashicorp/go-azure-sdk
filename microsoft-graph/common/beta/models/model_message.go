package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Message struct {
	Attachments                   *[]Attachment                        `json:"attachments,omitempty"`
	BccRecipients                 *[]Recipient                         `json:"bccRecipients,omitempty"`
	Body                          *ItemBody                            `json:"body,omitempty"`
	BodyPreview                   *string                              `json:"bodyPreview,omitempty"`
	Categories                    *[]string                            `json:"categories,omitempty"`
	CcRecipients                  *[]Recipient                         `json:"ccRecipients,omitempty"`
	ChangeKey                     *string                              `json:"changeKey,omitempty"`
	ConversationId                *string                              `json:"conversationId,omitempty"`
	ConversationIndex             *string                              `json:"conversationIndex,omitempty"`
	CreatedDateTime               *string                              `json:"createdDateTime,omitempty"`
	Extensions                    *[]Extension                         `json:"extensions,omitempty"`
	Flag                          *FollowupFlag                        `json:"flag,omitempty"`
	From                          *Recipient                           `json:"from,omitempty"`
	HasAttachments                *bool                                `json:"hasAttachments,omitempty"`
	Id                            *string                              `json:"id,omitempty"`
	Importance                    *MessageImportance                   `json:"importance,omitempty"`
	InferenceClassification       *MessageInferenceClassification      `json:"inferenceClassification,omitempty"`
	InternetMessageHeaders        *[]InternetMessageHeader             `json:"internetMessageHeaders,omitempty"`
	InternetMessageId             *string                              `json:"internetMessageId,omitempty"`
	IsDeliveryReceiptRequested    *bool                                `json:"isDeliveryReceiptRequested,omitempty"`
	IsDraft                       *bool                                `json:"isDraft,omitempty"`
	IsRead                        *bool                                `json:"isRead,omitempty"`
	IsReadReceiptRequested        *bool                                `json:"isReadReceiptRequested,omitempty"`
	LastModifiedDateTime          *string                              `json:"lastModifiedDateTime,omitempty"`
	Mentions                      *[]Mention                           `json:"mentions,omitempty"`
	MentionsPreview               *MentionsPreview                     `json:"mentionsPreview,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty  `json:"multiValueExtendedProperties,omitempty"`
	ODataType                     *string                              `json:"@odata.type,omitempty"`
	ParentFolderId                *string                              `json:"parentFolderId,omitempty"`
	ReceivedDateTime              *string                              `json:"receivedDateTime,omitempty"`
	ReplyTo                       *[]Recipient                         `json:"replyTo,omitempty"`
	Sender                        *Recipient                           `json:"sender,omitempty"`
	SentDateTime                  *string                              `json:"sentDateTime,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty `json:"singleValueExtendedProperties,omitempty"`
	Subject                       *string                              `json:"subject,omitempty"`
	ToRecipients                  *[]Recipient                         `json:"toRecipients,omitempty"`
	UniqueBody                    *ItemBody                            `json:"uniqueBody,omitempty"`
	UnsubscribeData               *[]string                            `json:"unsubscribeData,omitempty"`
	UnsubscribeEnabled            *bool                                `json:"unsubscribeEnabled,omitempty"`
	WebLink                       *string                              `json:"webLink,omitempty"`
}
