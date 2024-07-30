package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRulePredicates struct {
	BodyContains           *[]string                               `json:"bodyContains,omitempty"`
	BodyOrSubjectContains  *[]string                               `json:"bodyOrSubjectContains,omitempty"`
	Categories             *[]string                               `json:"categories,omitempty"`
	FromAddresses          *[]Recipient                            `json:"fromAddresses,omitempty"`
	HasAttachments         *bool                                   `json:"hasAttachments,omitempty"`
	HeaderContains         *[]string                               `json:"headerContains,omitempty"`
	Importance             *MessageRulePredicatesImportance        `json:"importance,omitempty"`
	IsApprovalRequest      *bool                                   `json:"isApprovalRequest,omitempty"`
	IsAutomaticForward     *bool                                   `json:"isAutomaticForward,omitempty"`
	IsAutomaticReply       *bool                                   `json:"isAutomaticReply,omitempty"`
	IsEncrypted            *bool                                   `json:"isEncrypted,omitempty"`
	IsMeetingRequest       *bool                                   `json:"isMeetingRequest,omitempty"`
	IsMeetingResponse      *bool                                   `json:"isMeetingResponse,omitempty"`
	IsNonDeliveryReport    *bool                                   `json:"isNonDeliveryReport,omitempty"`
	IsPermissionControlled *bool                                   `json:"isPermissionControlled,omitempty"`
	IsReadReceipt          *bool                                   `json:"isReadReceipt,omitempty"`
	IsSigned               *bool                                   `json:"isSigned,omitempty"`
	IsVoicemail            *bool                                   `json:"isVoicemail,omitempty"`
	MessageActionFlag      *MessageRulePredicatesMessageActionFlag `json:"messageActionFlag,omitempty"`
	NotSentToMe            *bool                                   `json:"notSentToMe,omitempty"`
	ODataType              *string                                 `json:"@odata.type,omitempty"`
	RecipientContains      *[]string                               `json:"recipientContains,omitempty"`
	SenderContains         *[]string                               `json:"senderContains,omitempty"`
	Sensitivity            *MessageRulePredicatesSensitivity       `json:"sensitivity,omitempty"`
	SentCcMe               *bool                                   `json:"sentCcMe,omitempty"`
	SentOnlyToMe           *bool                                   `json:"sentOnlyToMe,omitempty"`
	SentToAddresses        *[]Recipient                            `json:"sentToAddresses,omitempty"`
	SentToMe               *bool                                   `json:"sentToMe,omitempty"`
	SentToOrCcMe           *bool                                   `json:"sentToOrCcMe,omitempty"`
	SubjectContains        *[]string                               `json:"subjectContains,omitempty"`
	WithinSizeRange        *SizeRange                              `json:"withinSizeRange,omitempty"`
}
