package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationThreadOperationPredicate struct {
	HasAttachments        *bool
	Id                    *string
	IsLocked              *bool
	LastDeliveredDateTime *string
	ODataType             *string
	Preview               *string
	Topic                 *string
}

func (p ConversationThreadOperationPredicate) Matches(input ConversationThread) bool {

	if p.HasAttachments != nil && (input.HasAttachments == nil || *p.HasAttachments != *input.HasAttachments) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsLocked != nil && (input.IsLocked == nil || *p.IsLocked != *input.IsLocked) {
		return false
	}

	if p.LastDeliveredDateTime != nil && (input.LastDeliveredDateTime == nil || *p.LastDeliveredDateTime != *input.LastDeliveredDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Preview != nil && (input.Preview == nil || *p.Preview != *input.Preview) {
		return false
	}

	if p.Topic != nil && (input.Topic == nil || *p.Topic != *input.Topic) {
		return false
	}

	return true
}
