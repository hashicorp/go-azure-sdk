package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttendanceIntervalOperationPredicate struct {
	DurationInSeconds *int64
	JoinDateTime      *string
	LeaveDateTime     *string
	ODataType         *string
}

func (p AttendanceIntervalOperationPredicate) Matches(input AttendanceInterval) bool {

	if p.DurationInSeconds != nil && (input.DurationInSeconds == nil || *p.DurationInSeconds != *input.DurationInSeconds) {
		return false
	}

	if p.JoinDateTime != nil && (input.JoinDateTime == nil || *p.JoinDateTime != *input.JoinDateTime) {
		return false
	}

	if p.LeaveDateTime != nil && (input.LeaveDateTime == nil || *p.LeaveDateTime != *input.LeaveDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
