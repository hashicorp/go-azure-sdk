package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingReminderOperationPredicate struct {
	Message   *string
	ODataType *string
	Offset    *string
}

func (p BookingReminderOperationPredicate) Matches(input BookingReminder) bool {

	if p.Message != nil && (input.Message == nil || *p.Message != *input.Message) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Offset != nil && (input.Offset == nil || *p.Offset != *input.Offset) {
		return false
	}

	return true
}
