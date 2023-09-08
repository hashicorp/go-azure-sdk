package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingCapabilityOperationPredicate struct {
	AllowAnonymousUsersToDialOut      *bool
	AllowAnonymousUsersToStartMeeting *bool
	ODataType                         *string
}

func (p MeetingCapabilityOperationPredicate) Matches(input MeetingCapability) bool {

	if p.AllowAnonymousUsersToDialOut != nil && (input.AllowAnonymousUsersToDialOut == nil || *p.AllowAnonymousUsersToDialOut != *input.AllowAnonymousUsersToDialOut) {
		return false
	}

	if p.AllowAnonymousUsersToStartMeeting != nil && (input.AllowAnonymousUsersToStartMeeting == nil || *p.AllowAnonymousUsersToStartMeeting != *input.AllowAnonymousUsersToStartMeeting) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
