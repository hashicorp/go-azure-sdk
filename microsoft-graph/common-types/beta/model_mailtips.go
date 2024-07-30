package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MailTips struct {
	AutomaticReplies     *AutomaticRepliesMailTips `json:"automaticReplies,omitempty"`
	CustomMailTip        *string                   `json:"customMailTip,omitempty"`
	DeliveryRestricted   *bool                     `json:"deliveryRestricted,omitempty"`
	EmailAddress         *EmailAddress             `json:"emailAddress,omitempty"`
	Error                *MailTipsError            `json:"error,omitempty"`
	ExternalMemberCount  *int64                    `json:"externalMemberCount,omitempty"`
	IsModerated          *bool                     `json:"isModerated,omitempty"`
	MailboxFull          *bool                     `json:"mailboxFull,omitempty"`
	MaxMessageSize       *int64                    `json:"maxMessageSize,omitempty"`
	ODataType            *string                   `json:"@odata.type,omitempty"`
	RecipientScope       *MailTipsRecipientScope   `json:"recipientScope,omitempty"`
	RecipientSuggestions *[]Recipient              `json:"recipientSuggestions,omitempty"`
	TotalMemberCount     *int64                    `json:"totalMemberCount,omitempty"`
}
