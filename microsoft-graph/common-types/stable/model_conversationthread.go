package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThread struct {
	CcRecipients          *[]Recipient `json:"ccRecipients,omitempty"`
	HasAttachments        *bool        `json:"hasAttachments,omitempty"`
	Id                    *string      `json:"id,omitempty"`
	IsLocked              *bool        `json:"isLocked,omitempty"`
	LastDeliveredDateTime *string      `json:"lastDeliveredDateTime,omitempty"`
	ODataType             *string      `json:"@odata.type,omitempty"`
	Posts                 *[]Post      `json:"posts,omitempty"`
	Preview               *string      `json:"preview,omitempty"`
	ToRecipients          *[]Recipient `json:"toRecipients,omitempty"`
	Topic                 *string      `json:"topic,omitempty"`
	UniqueSenders         *[]string    `json:"uniqueSenders,omitempty"`
}
