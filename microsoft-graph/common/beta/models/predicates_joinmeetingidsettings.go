package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinMeetingIdSettingsOperationPredicate struct {
	IsPasscodeRequired *bool
	JoinMeetingId      *string
	ODataType          *string
	Passcode           *string
}

func (p JoinMeetingIdSettingsOperationPredicate) Matches(input JoinMeetingIdSettings) bool {

	if p.IsPasscodeRequired != nil && (input.IsPasscodeRequired == nil || *p.IsPasscodeRequired != *input.IsPasscodeRequired) {
		return false
	}

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
