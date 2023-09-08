package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatInfoOperationPredicate struct {
	MessageId           *string
	ODataType           *string
	ReplyChainMessageId *string
	ThreadId            *string
}

func (p ChatInfoOperationPredicate) Matches(input ChatInfo) bool {

	if p.MessageId != nil && (input.MessageId == nil || *p.MessageId != *input.MessageId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReplyChainMessageId != nil && (input.ReplyChainMessageId == nil || *p.ReplyChainMessageId != *input.ReplyChainMessageId) {
		return false
	}

	if p.ThreadId != nil && (input.ThreadId == nil || *p.ThreadId != *input.ThreadId) {
		return false
	}

	return true
}
