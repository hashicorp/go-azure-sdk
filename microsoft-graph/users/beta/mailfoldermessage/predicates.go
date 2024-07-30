package mailfoldermessage

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type MessageOperationPredicate struct {
	BodyPreview                *string
	ChangeKey                  *string
	ConversationId             *string
	ConversationIndex          *string
	CreatedDateTime            *string
	HasAttachments             *bool
	Id                         *string
	InternetMessageId          *string
	IsDeliveryReceiptRequested *bool
	IsDraft                    *bool
	IsRead                     *bool
	IsReadReceiptRequested     *bool
	LastModifiedDateTime       *string
	ODataType                  *string
	ParentFolderId             *string
	ReceivedDateTime           *string
	SentDateTime               *string
	Subject                    *string
	UnsubscribeEnabled         *bool
	WebLink                    *string
}

func (p MessageOperationPredicate) Matches(input beta.Message) bool {

	if p.BodyPreview != nil && (input.BodyPreview == nil || *p.BodyPreview != *input.BodyPreview) {
		return false
	}

	if p.ChangeKey != nil && (input.ChangeKey == nil || *p.ChangeKey != *input.ChangeKey) {
		return false
	}

	if p.ConversationId != nil && (input.ConversationId == nil || *p.ConversationId != *input.ConversationId) {
		return false
	}

	if p.ConversationIndex != nil && (input.ConversationIndex == nil || *p.ConversationIndex != *input.ConversationIndex) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.HasAttachments != nil && (input.HasAttachments == nil || *p.HasAttachments != *input.HasAttachments) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InternetMessageId != nil && (input.InternetMessageId == nil || *p.InternetMessageId != *input.InternetMessageId) {
		return false
	}

	if p.IsDeliveryReceiptRequested != nil && (input.IsDeliveryReceiptRequested == nil || *p.IsDeliveryReceiptRequested != *input.IsDeliveryReceiptRequested) {
		return false
	}

	if p.IsDraft != nil && (input.IsDraft == nil || *p.IsDraft != *input.IsDraft) {
		return false
	}

	if p.IsRead != nil && (input.IsRead == nil || *p.IsRead != *input.IsRead) {
		return false
	}

	if p.IsReadReceiptRequested != nil && (input.IsReadReceiptRequested == nil || *p.IsReadReceiptRequested != *input.IsReadReceiptRequested) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ParentFolderId != nil && (input.ParentFolderId == nil || *p.ParentFolderId != *input.ParentFolderId) {
		return false
	}

	if p.ReceivedDateTime != nil && (input.ReceivedDateTime == nil || *p.ReceivedDateTime != *input.ReceivedDateTime) {
		return false
	}

	if p.SentDateTime != nil && (input.SentDateTime == nil || *p.SentDateTime != *input.SentDateTime) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	if p.UnsubscribeEnabled != nil && (input.UnsubscribeEnabled == nil || *p.UnsubscribeEnabled != *input.UnsubscribeEnabled) {
		return false
	}

	if p.WebLink != nil && (input.WebLink == nil || *p.WebLink != *input.WebLink) {
		return false
	}

	return true
}
