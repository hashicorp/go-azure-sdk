package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizerMeetingInfoOperationPredicate struct {
	AllowConversationWithoutHost *bool
	ODataType                    *string
}

func (p OrganizerMeetingInfoOperationPredicate) Matches(input OrganizerMeetingInfo) bool {

	if p.AllowConversationWithoutHost != nil && (input.AllowConversationWithoutHost == nil || *p.AllowConversationWithoutHost != *input.AllowConversationWithoutHost) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
