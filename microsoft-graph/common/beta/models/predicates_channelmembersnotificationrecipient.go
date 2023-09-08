package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelMembersNotificationRecipientOperationPredicate struct {
	ChannelId *string
	ODataType *string
	TeamId    *string
}

func (p ChannelMembersNotificationRecipientOperationPredicate) Matches(input ChannelMembersNotificationRecipient) bool {

	if p.ChannelId != nil && (input.ChannelId == nil || *p.ChannelId != *input.ChannelId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TeamId != nil && (input.TeamId == nil || *p.TeamId != *input.TeamId) {
		return false
	}

	return true
}
