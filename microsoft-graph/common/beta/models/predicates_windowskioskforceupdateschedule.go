package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskForceUpdateScheduleOperationPredicate struct {
	DayofMonth                         *int64
	ODataType                          *string
	RunImmediatelyIfAfterStartDateTime *bool
	StartDateTime                      *string
}

func (p WindowsKioskForceUpdateScheduleOperationPredicate) Matches(input WindowsKioskForceUpdateSchedule) bool {

	if p.DayofMonth != nil && (input.DayofMonth == nil || *p.DayofMonth != *input.DayofMonth) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RunImmediatelyIfAfterStartDateTime != nil && (input.RunImmediatelyIfAfterStartDateTime == nil || *p.RunImmediatelyIfAfterStartDateTime != *input.RunImmediatelyIfAfterStartDateTime) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	return true
}
