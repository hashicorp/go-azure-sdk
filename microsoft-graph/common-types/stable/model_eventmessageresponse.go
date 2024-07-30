package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventMessageResponse struct {
	Attachments                   *[]Attachment                                `json:"attachments,omitempty"`
	BccRecipients                 *[]Recipient                                 `json:"bccRecipients,omitempty"`
	Body                          *ItemBody                                    `json:"body,omitempty"`
	BodyPreview                   *string                                      `json:"bodyPreview,omitempty"`
	Categories                    *[]string                                    `json:"categories,omitempty"`
	CcRecipients                  *[]Recipient                                 `json:"ccRecipients,omitempty"`
	ChangeKey                     *string                                      `json:"changeKey,omitempty"`
	ConversationId                *string                                      `json:"conversationId,omitempty"`
	ConversationIndex             *string                                      `json:"conversationIndex,omitempty"`
	CreatedDateTime               *string                                      `json:"createdDateTime,omitempty"`
	EndDateTime                   *DateTimeTimeZone                            `json:"endDateTime,omitempty"`
	Event                         *Event                                       `json:"event,omitempty"`
	Extensions                    *[]Extension                                 `json:"extensions,omitempty"`
	Flag                          *FollowupFlag                                `json:"flag,omitempty"`
	From                          *Recipient                                   `json:"from,omitempty"`
	HasAttachments                *bool                                        `json:"hasAttachments,omitempty"`
	Id                            *string                                      `json:"id,omitempty"`
	Importance                    *EventMessageResponseImportance              `json:"importance,omitempty"`
	InferenceClassification       *EventMessageResponseInferenceClassification `json:"inferenceClassification,omitempty"`
	InternetMessageHeaders        *[]InternetMessageHeader                     `json:"internetMessageHeaders,omitempty"`
	InternetMessageId             *string                                      `json:"internetMessageId,omitempty"`
	IsAllDay                      *bool                                        `json:"isAllDay,omitempty"`
	IsDelegated                   *bool                                        `json:"isDelegated,omitempty"`
	IsDeliveryReceiptRequested    *bool                                        `json:"isDeliveryReceiptRequested,omitempty"`
	IsDraft                       *bool                                        `json:"isDraft,omitempty"`
	IsOutOfDate                   *bool                                        `json:"isOutOfDate,omitempty"`
	IsRead                        *bool                                        `json:"isRead,omitempty"`
	IsReadReceiptRequested        *bool                                        `json:"isReadReceiptRequested,omitempty"`
	LastModifiedDateTime          *string                                      `json:"lastModifiedDateTime,omitempty"`
	Location                      *Location                                    `json:"location,omitempty"`
	MeetingMessageType            *EventMessageResponseMeetingMessageType      `json:"meetingMessageType,omitempty"`
	MultiValueExtendedProperties  *[]MultiValueLegacyExtendedProperty          `json:"multiValueExtendedProperties,omitempty"`
	ODataType                     *string                                      `json:"@odata.type,omitempty"`
	ParentFolderId                *string                                      `json:"parentFolderId,omitempty"`
	ProposedNewTime               *TimeSlot                                    `json:"proposedNewTime,omitempty"`
	ReceivedDateTime              *string                                      `json:"receivedDateTime,omitempty"`
	Recurrence                    *PatternedRecurrence                         `json:"recurrence,omitempty"`
	ReplyTo                       *[]Recipient                                 `json:"replyTo,omitempty"`
	ResponseType                  *EventMessageResponseResponseType            `json:"responseType,omitempty"`
	Sender                        *Recipient                                   `json:"sender,omitempty"`
	SentDateTime                  *string                                      `json:"sentDateTime,omitempty"`
	SingleValueExtendedProperties *[]SingleValueLegacyExtendedProperty         `json:"singleValueExtendedProperties,omitempty"`
	StartDateTime                 *DateTimeTimeZone                            `json:"startDateTime,omitempty"`
	Subject                       *string                                      `json:"subject,omitempty"`
	ToRecipients                  *[]Recipient                                 `json:"toRecipients,omitempty"`
	Type                          *EventMessageResponseType                    `json:"type,omitempty"`
	UniqueBody                    *ItemBody                                    `json:"uniqueBody,omitempty"`
	WebLink                       *string                                      `json:"webLink,omitempty"`
}
