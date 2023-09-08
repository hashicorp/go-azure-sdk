package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BroadcastMeetingSettingsOperationPredicate struct {
	IsAttendeeReportEnabled    *bool
	IsQuestionAndAnswerEnabled *bool
	IsRecordingEnabled         *bool
	IsVideoOnDemandEnabled     *bool
	ODataType                  *string
}

func (p BroadcastMeetingSettingsOperationPredicate) Matches(input BroadcastMeetingSettings) bool {

	if p.IsAttendeeReportEnabled != nil && (input.IsAttendeeReportEnabled == nil || *p.IsAttendeeReportEnabled != *input.IsAttendeeReportEnabled) {
		return false
	}

	if p.IsQuestionAndAnswerEnabled != nil && (input.IsQuestionAndAnswerEnabled == nil || *p.IsQuestionAndAnswerEnabled != *input.IsQuestionAndAnswerEnabled) {
		return false
	}

	if p.IsRecordingEnabled != nil && (input.IsRecordingEnabled == nil || *p.IsRecordingEnabled != *input.IsRecordingEnabled) {
		return false
	}

	if p.IsVideoOnDemandEnabled != nil && (input.IsVideoOnDemandEnabled == nil || *p.IsVideoOnDemandEnabled != *input.IsVideoOnDemandEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
