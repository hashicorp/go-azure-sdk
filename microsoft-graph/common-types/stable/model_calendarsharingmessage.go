package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalendarSharingMessage struct {
	Attachments                   *[]Attachment                                  `json:"attachments,omitempty"`
	BccRecipients                 *[]Recipient                                   `json:"bccRecipients,omitempty"`
	Body                          *ItemBody                                      `json:"body,omitempty"`
	BodyPreview                   *string                                        `json:"bodyPreview,omitempty"`
	CanAccept                     *bool                                          `json:"canAccept,omitempty"`
	Categories                    *[]string                                      `json:"categories,omitempty"`
	CcRecipients                  *[]Recipient                                   `json:"ccRecipients,omitempty"`
	ChangeKey                     *string                                        `json:"changeKey,omitempty"`
	ConversationId                *string                                        `json:"conversationId,omitempty"`
	ConversationIndex             *string                                        `json:"conversationIndex,omitempty"`
	CreatedDateTime               *string                                        `json:"createdDateTime,omitempty"`
	Extensions                    *[]Extension                                   `json:"extensions,omitempty"`
	Flag                          *FollowupFlag                                  `json:"flag,omitempty"`
	From                          *Recipient                                     `json:"from,omitempty"`
	HasAttachments                *bool                                          `json:"hasAttachments,omitempty"`
	Id                            *string                                        `json:"id,omitempty"`
	Importance                    *CalendarSharingMessageImportance              `json:"importance,omitempty"`
	InferenceClassification       *CalendarSharingMessageInferenceClassification `json:"inferenceClassification,omitempty"`
	InternetMessageHeaders        *[]InternetMessageHeader                       `json:"internetMessageHeaders,omitempty"`
	InternetMessageId             *string                                        `json:"internetMessageId,omitempty"`
	IsDeliveryReceiptRequested    *bool                                          `json:"isDeliveryReceiptRequested,omitempty"`
	IsDraft                       *bool                                          `json:"isDraft,omitempty"`
	IsRead                        *bool                                          `json:"isRead,omitempty"`
	IsReadReceiptRequested        *bool                                          `json:"isReadReceiptRequested,omitempty"`
	LastModifiedDateTime          *string                                        `json:"lastModifiedDateTime,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty            `json:"multiValueExtendedProperties,omitempty"`
	ODataType                     *string                                        `json:"@odata.type,omitempty"`
	ParentFolderId                *string                                        `json:"parentFolderId,omitempty"`
	ReceivedDateTime              *string                                        `json:"receivedDateTime,omitempty"`
	ReplyTo                       *[]Recipient                                   `json:"replyTo,omitempty"`
	Sender                        *Recipient                                     `json:"sender,omitempty"`
	SentDateTime                  *string                                        `json:"sentDateTime,omitempty"`
	SharingMessageAction          *CalendarSharingMessageAction                  `json:"sharingMessageAction,omitempty"`
	SharingMessageActions         *[]CalendarSharingMessageAction                `json:"sharingMessageActions,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty           `json:"singleValueExtendedProperties,omitempty"`
	Subject                       *string                                        `json:"subject,omitempty"`
	SuggestedCalendarName         *string                                        `json:"suggestedCalendarName,omitempty"`
	ToRecipients                  *[]Recipient                                   `json:"toRecipients,omitempty"`
	UniqueBody                    *ItemBody                                      `json:"uniqueBody,omitempty"`
	WebLink                       *string                                        `json:"webLink,omitempty"`
}
