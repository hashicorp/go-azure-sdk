package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMembersNotificationRecipientOperationPredicate struct {
	ChatId    *string
	ODataType *string
}

func (p ChatMembersNotificationRecipientOperationPredicate) Matches(input ChatMembersNotificationRecipient) bool {

	if p.ChatId != nil && (input.ChatId == nil || *p.ChatId != *input.ChatId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
