package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BroadcastMeetingSettings struct {
	AllowedAudience            *BroadcastMeetingSettingsAllowedAudience `json:"allowedAudience,omitempty"`
	Captions                   *BroadcastMeetingCaptionSettings         `json:"captions,omitempty"`
	IsAttendeeReportEnabled    *bool                                    `json:"isAttendeeReportEnabled,omitempty"`
	IsQuestionAndAnswerEnabled *bool                                    `json:"isQuestionAndAnswerEnabled,omitempty"`
	IsRecordingEnabled         *bool                                    `json:"isRecordingEnabled,omitempty"`
	IsVideoOnDemandEnabled     *bool                                    `json:"isVideoOnDemandEnabled,omitempty"`
	ODataType                  *string                                  `json:"@odata.type,omitempty"`
}
