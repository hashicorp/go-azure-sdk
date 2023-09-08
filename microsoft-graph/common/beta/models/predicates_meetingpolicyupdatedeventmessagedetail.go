package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingPolicyUpdatedEventMessageDetailOperationPredicate struct {
	MeetingChatEnabled *bool
	MeetingChatId      *string
	ODataType          *string
}

func (p MeetingPolicyUpdatedEventMessageDetailOperationPredicate) Matches(input MeetingPolicyUpdatedEventMessageDetail) bool {

	if p.MeetingChatEnabled != nil && (input.MeetingChatEnabled == nil || *p.MeetingChatEnabled != *input.MeetingChatEnabled) {
		return false
	}

	if p.MeetingChatId != nil && (input.MeetingChatId == nil || *p.MeetingChatId != *input.MeetingChatId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
