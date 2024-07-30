package conversationthreadpost

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type PostOperationPredicate struct {
	ChangeKey            *string
	ConversationId       *string
	ConversationThreadId *string
	CreatedDateTime      *string
	HasAttachments       *bool
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	ReceivedDateTime     *string
}

func (p PostOperationPredicate) Matches(input beta.Post) bool {

	if p.ChangeKey != nil && (input.ChangeKey == nil || *p.ChangeKey != *input.ChangeKey) {
		return false
	}

	if p.ConversationId != nil && (input.ConversationId == nil || *p.ConversationId != *input.ConversationId) {
		return false
	}

	if p.ConversationThreadId != nil && (input.ConversationThreadId == nil || *p.ConversationThreadId != *input.ConversationThreadId) {
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

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReceivedDateTime != nil && (input.ReceivedDateTime == nil || *p.ReceivedDateTime != *input.ReceivedDateTime) {
		return false
	}

	return true
}
