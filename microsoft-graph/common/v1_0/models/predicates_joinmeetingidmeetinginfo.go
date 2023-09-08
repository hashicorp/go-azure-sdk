package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinMeetingIdMeetingInfoOperationPredicate struct {
	JoinMeetingId *string
	ODataType     *string
	Passcode      *string
}

func (p JoinMeetingIdMeetingInfoOperationPredicate) Matches(input JoinMeetingIdMeetingInfo) bool {

	if p.JoinMeetingId != nil && (input.JoinMeetingId == nil || *p.JoinMeetingId != *input.JoinMeetingId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Passcode != nil && (input.Passcode == nil || *p.Passcode != *input.Passcode) {
		return false
	}

	return true
}
