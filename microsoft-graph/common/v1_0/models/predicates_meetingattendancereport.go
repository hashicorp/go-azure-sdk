package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingAttendanceReportOperationPredicate struct {
	Id                    *string
	MeetingEndDateTime    *string
	MeetingStartDateTime  *string
	ODataType             *string
	TotalParticipantCount *int64
}

func (p MeetingAttendanceReportOperationPredicate) Matches(input MeetingAttendanceReport) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MeetingEndDateTime != nil && (input.MeetingEndDateTime == nil || *p.MeetingEndDateTime != *input.MeetingEndDateTime) {
		return false
	}

	if p.MeetingStartDateTime != nil && (input.MeetingStartDateTime == nil || *p.MeetingStartDateTime != *input.MeetingStartDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TotalParticipantCount != nil && (input.TotalParticipantCount == nil || *p.TotalParticipantCount != *input.TotalParticipantCount) {
		return false
	}

	return true
}
