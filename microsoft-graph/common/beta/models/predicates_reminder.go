package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReminderOperationPredicate struct {
	ChangeKey    *string
	EventId      *string
	EventSubject *string
	EventWebLink *string
	ODataType    *string
}

func (p ReminderOperationPredicate) Matches(input Reminder) bool {

	if p.ChangeKey != nil && (input.ChangeKey == nil || *p.ChangeKey != *input.ChangeKey) {
		return false
	}

	if p.EventId != nil && (input.EventId == nil || *p.EventId != *input.EventId) {
		return false
	}

	if p.EventSubject != nil && (input.EventSubject == nil || *p.EventSubject != *input.EventSubject) {
		return false
	}

	if p.EventWebLink != nil && (input.EventWebLink == nil || *p.EventWebLink != *input.EventWebLink) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
